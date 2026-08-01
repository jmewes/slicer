// Copyright (C) 2026 John Doe
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

package core

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	messages "github.com/cucumber/messages/go/v28"
	specantlr "github.com/experimental-software/gherkin/core/antlr"
	"github.com/experimental-software/gherkin/utils"
)

// normalizeScenarioTitle strips a leading "should " prefix (case-insensitive).
func normalizeScenarioTitle(title string) string {
	if strings.HasPrefix(strings.ToLower(title), "should ") {
		return title[len("should "):]
	}
	return title
}

var stepCommentRe = regexp.MustCompile(`(?i)^\s*//\s*(given|when|then|and|but)\s+(.+)$`)

// matchStepComment returns the keyword (e.g. "Given ") and the remainder text
// if line matches a Gherkin step comment, otherwise ok is false.
func matchStepComment(line string) (keyword, text string, ok bool) {
	m := stepCommentRe.FindStringSubmatch(line)
	if m == nil {
		return "", "", false
	}
	kw := strings.ToUpper(m[1][:1]) + strings.ToLower(m[1][1:])
	return kw + " ", strings.TrimSpace(m[2]), true
}

const (
	frameDescribe   = "describe"
	frameIt         = "it"
	frameBeforeEach = "beforeEach"
)

// frame is the intermediate representation shared between the ANTLR-driven
// visitor (see spec_visitor.go) and the feature-document emitter below. It
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

// collectingErrorListener records ANTLR syntax errors so that ParseSpecFile
// can preserve the "return whatever we managed to parse, even on error"
// contract that `cmd/rev.go` relies on.
type collectingErrorListener struct {
	*antlr.DefaultErrorListener
	errs []string
}

func (l *collectingErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, col int, msg string, _ antlr.RecognitionException) {
	l.errs = append(l.errs, fmt.Sprintf("line %d:%d %s", line, col, msg))
}

// ParseSpecFile reads a Jasmine *.spec.ts file and converts its
// describe/it/beforeEach structure into GherkinDocuments. It uses an
// ANTLR-generated parser (see core/antlr/SpecGrammar.g4) driving a small
// tree visitor that emits the same *frame tree the previous line-scanner
// produced, so downstream emission logic is unchanged.
func ParseSpecFile(path string, relaxed bool) ([]*messages.GherkinDocument, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	input := antlr.NewInputStream(string(data))
	lexer := specantlr.NewSpecGrammarLexer(input)
	lexer.RemoveErrorListeners()
	lexErrs := &collectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	lexer.AddErrorListener(lexErrs)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := specantlr.NewSpecGrammarParser(tokens)
	parser.RemoveErrorListeners()
	parseErrs := &collectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	parser.AddErrorListener(parseErrs)

	tree := parser.Program()

	visitor := &specVisitor{
		tokens:      tokens,
		relaxedMode: relaxed,
	}
	visitor.visitProgram(tree)

	var docs []*messages.GherkinDocument
	for _, d := range visitor.topLevelDescribes {
		emitDescribe(d, path, nil, &docs)
	}

	if errs := append(append([]string{}, lexErrs.errs...), parseErrs.errs...); len(errs) > 0 {
		return docs, fmt.Errorf("parse errors in %s: %s", path, strings.Join(errs, "; "))
	}
	return docs, nil
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
