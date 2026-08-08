// Copyright (C) 2026 Jan Mewes
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package javascript

import (
	"strings"

	messages "github.com/cucumber/messages/go/v28"
	"github.com/experimental-software/gherkin/utils"
)

const (
	frameDescribe   = "describe"
	frameIt         = "it"
	frameBeforeEach = "beforeEach"
)

// frame is the intermediate representation shared between the ANTLR-driven
// visitor (see visitor.go) and the feature-document emitter below. It
// captures a single describe / it / beforeEach node together with anything
// that has been discovered inside it while walking the parse tree.
type frame struct {
	kind       string
	name       string   // describe or it title
	ancestors  []string // ancestor describe names (for describe frames)
	steps      []*messages.Step
	scenarios  []*messages.Scenario
	background []*messages.Step
	children   []*frame // nested describe frames
}

// hasTests reports whether the describe frame contains scenarios directly or
// in any transitively nested describe.
func hasTests(d *frame) bool {
	if len(d.scenarios) > 0 {
		return true
	}
	for _, c := range d.children {
		if hasTests(c) {
			return true
		}
	}
	return false
}

// emitDescribe converts a describe frame (and its nested describes) into one
// or more GherkinDocuments and appends them to docs. The uriAncestors argument
// contains the describe names that lead to the current frame but were skipped
// during emission (i.e. their parent describes had no scenarios and were not
// themselves emitted as Features).
func emitDescribe(d *frame, path string, uriAncestors []string, docs *[]*messages.GherkinDocument) {
	var testChildren []*frame
	for _, c := range d.children {
		if hasTests(c) {
			testChildren = append(testChildren, c)
		}
	}

	// Case: describe with no test-bearing children.
	if len(testChildren) == 0 {
		if len(d.scenarios) == 0 {
			return
		}
		*docs = append(*docs, buildFeatureDoc(d, path, uriAncestors))
		return
	}

	// Case: describe has scenarios and nested test-bearing describes.
	// Emit this describe as its own Feature (with just its direct
	// scenarios), and each test-bearing child as a separate top-level
	// Feature.
	if len(d.scenarios) > 0 {
		*docs = append(*docs, buildFeatureDoc(d, path, uriAncestors))
		childAncestors := append(append([]string{}, uriAncestors...), d.name)
		for _, c := range testChildren {
			emitDescribe(c, path, childAncestors, docs)
		}
		return
	}

	childAncestors := append(append([]string{}, uriAncestors...), d.name)
	for _, c := range testChildren {
		emitDescribe(c, path, childAncestors, docs)
	}
}

func buildFeatureDoc(d *frame, path string, uriAncestors []string) *messages.GherkinDocument {
	var children []*messages.FeatureChild
	if len(d.background) > 0 {
		children = append(children, &messages.FeatureChild{
			Background: &messages.Background{Steps: d.background},
		})
	}
	for _, sc := range d.scenarios {
		children = append(children, &messages.FeatureChild{Scenario: sc})
	}

	var uriParts []string
	for _, p := range append(append([]string{}, uriAncestors...), d.name) {
		uriParts = append(uriParts, utils.ToKebabCase(p))
	}
	uri := path + "/" + strings.Join(uriParts, "/")
	if idx := strings.Index(uri, "src/app/"); idx >= 0 {
		uri = uri[idx+len("src/app/"):]
	}

	return &messages.GherkinDocument{
		Uri: normalizeUri(uri),
		Feature: &messages.Feature{
			Name:     d.name,
			Children: children,
		},
	}
}

func normalizeUri(path string) string {
	result := ""
	pathParts := strings.Split(path, "/")
	for i := 0; i < len(pathParts); i++ {
		pathPart := pathParts[i]
		pathPart = strings.Replace(pathPart, ".spec.ts", "", 1)
		result += utils.ToKebabCase(pathPart)
		if i+1 < len(pathParts) {
			result += "/"
		}
	}
	return result
}
