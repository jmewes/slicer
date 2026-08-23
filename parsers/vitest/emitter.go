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

package vitest

import (
	"path/filepath"
	"strings"
	"unicode"

	messages "github.com/cucumber/messages/go/v28"
	"github.com/jmewes/slicer/utils"
)

const (
	frameDescribe = "describe"
	frameIt       = "it"
	frameBefore   = "before"
	frameAfter    = "after"
)

// frame is the intermediate representation shared between the ANTLR-driven
// visitor (see visitor.go) and the feature-document emitter below. It
// captures a single describe / it node together with anything that has been
// discovered inside it while walking the parse tree.
type frame struct {
	kind       string
	name       string   // describe or it/test title
	ancestors  []string // ancestor describe names (for describe frames)
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

// emitAll converts the visitor's collected frames into GherkinDocuments. It
// emits the top-level describes exactly like the Jasmine emitter, and, when
// the visitor collected top-level (describe-less) scenarios, prepends a
// synthetic Feature named after the file itself so those scenarios are not
// dropped.
func emitAll(v *specVisitor, path string, docs *[]*messages.GherkinDocument) {
	for _, d := range v.topLevelDescribes {
		emitDescribe(d, path, nil, docs)
	}

	if len(v.rootScenarios) > 0 {
		rootDoc := buildRootFeatureDoc(v, path)
		*docs = append([]*messages.GherkinDocument{rootDoc}, *docs...)
	}
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

	uri = filepath.ToSlash(uri)

	return &messages.GherkinDocument{
		Uri: normalizeUri(uri),
		Feature: &messages.Feature{
			Name:     d.name,
			Children: children,
		},
	}
}

// buildRootFeatureDoc builds the synthetic Feature for describe-less files:
// its name is derived from the file stem (see featureNameFromPath) and its
// URI is the kebab-cased file path itself, with no extra segment appended.
func buildRootFeatureDoc(v *specVisitor, path string) *messages.GherkinDocument {
	var children []*messages.FeatureChild
	if len(v.rootBackground) > 0 {
		children = append(children, &messages.FeatureChild{
			Background: &messages.Background{Steps: v.rootBackground},
		})
	}
	for _, sc := range v.rootScenarios {
		children = append(children, &messages.FeatureChild{Scenario: sc})
	}

	uri := normalizeUri(filepath.ToSlash(path))

	return &messages.GherkinDocument{
		Uri: uri,
		Feature: &messages.Feature{
			Name:     featureNameFromPath(path),
			Children: children,
		},
	}
}

// featureNameFromPath derives an implicit Feature name from a spec file's
// path: the file stem with its recognized test/spec suffix stripped,
// `-`/`_` replaced by spaces, and the first rune capitalized.
func featureNameFromPath(path string) string {
	base := filepath.Base(path)
	stem := stripSupportedSuffix(base)
	stem = strings.ReplaceAll(stem, "-", " ")
	stem = strings.ReplaceAll(stem, "_", " ")
	if stem == "" {
		return stem
	}
	r := []rune(stem)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// stripSupportedSuffix removes one of the eight supported *.test.* /
// *.spec.* suffixes from s (case-insensitively), if present.
func stripSupportedSuffix(s string) string {
	lower := strings.ToLower(s)
	for _, suf := range SupportedExtensions() {
		if strings.HasSuffix(lower, suf) {
			return s[:len(s)-len(suf)]
		}
	}
	return s
}

func normalizeUri(path string) string {
	result := ""
	pathParts := strings.Split(path, "/")
	for i := 0; i < len(pathParts); i++ {
		pathPart := stripSupportedSuffix(pathParts[i])
		result += utils.ToKebabCase(pathPart)
		if i+1 < len(pathParts) {
			result += "/"
		}
	}
	return result
}
