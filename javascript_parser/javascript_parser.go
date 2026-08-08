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

package javascript_parser

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	messages "github.com/cucumber/messages/go/v28"
	specantlr "github.com/experimental-software/gherkin/javascript_parser/antlr"
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

// specVisitor walks the ANTLR parse tree of a Jasmine-style spec file and
// builds a tree of *frame values equivalent to the one produced by the old
// line scanner. It replaces the manual brace-depth loop in the previous
// implementation: nesting is derived from the tree, and Gherkin steps are
// pulled from hidden-channel comment tokens that fall inside each callback's
// token span.
type specVisitor struct {
	tokens            *antlr.CommonTokenStream
	relaxedMode       bool
	topLevelDescribes []*frame
	describeStack     []*frame
}

// visitProgram is the entry point: it walks every top-level element of the
// program rule.
func (v *specVisitor) visitProgram(ctx specantlr.IProgramContext) {
	for _, el := range ctx.AllElement() {
		v.visitElement(el)
	}
}

// visitElement dispatches on the kind of element (suite, block, or filler)
// and recurses into structural nodes so we discover nested suites regardless
// of the exact surface syntax.
func (v *specVisitor) visitElement(el specantlr.IElementContext) {
	if el == nil {
		return
	}
	if s := el.Suite(); s != nil {
		v.visitSuite(s)
		return
	}
	if b := el.Block(); b != nil {
		v.visitBlock(b)
	}
}

// visitBlock walks the children of a `{ ... }` block, allowing us to reach
// suites nested inside stray braces (e.g. IIFE wrappers, arrow-function
// bodies used as filler).
func (v *specVisitor) visitBlock(b specantlr.IBlockContext) {
	for _, el := range b.AllElement() {
		v.visitElement(el)
	}
}

// visitSuite handles a describe/it/beforeEach call: it opens the appropriate
// frame, walks its callback block for nested suites, and then closes the
// frame — mirroring the semantics of the old brace-depth stack.
func (v *specVisitor) visitSuite(s specantlr.ISuiteContext) {
	kw := suiteKind(s.SuiteKeyword())
	if kw == "" {
		return
	}

	name, body := extractNameAndBlock(s.AllSuiteArg())

	switch kw {
	case frameDescribe:
		if name == "" {
			// Without a name we cannot build a Feature — still walk the body
			// so nested suites are discovered.
			if body != nil {
				v.visitBlock(body)
			}
			return
		}
		var ancestors []string
		for _, f := range v.describeStack {
			ancestors = append(ancestors, f.name)
		}
		fr := &frame{kind: frameDescribe, name: name, ancestors: ancestors}
		v.describeStack = append(v.describeStack, fr)
		if body != nil {
			v.visitBlock(body)
		}
		v.describeStack = v.describeStack[:len(v.describeStack)-1]

		if len(ancestors) == 0 {
			v.topLevelDescribes = append(v.topLevelDescribes, fr)
		} else {
			parent := v.describeStack[len(v.describeStack)-1]
			parent.children = append(parent.children, fr)
		}

	case frameIt:
		if name == "" {
			return
		}
		steps := v.stepsFromBlock(body)
		if enclosing := v.enclosingDescribe(); enclosing != nil {
			if len(steps) > 0 || v.relaxed() {
				enclosing.scenarios = append(enclosing.scenarios, &messages.Scenario{
					Name:  normalizeScenarioTitle(name),
					Steps: steps,
				})
			}
		}

	case frameBeforeEach:
		steps := v.stepsFromBlock(body)
		var bgSteps []*messages.Step
		for _, s := range steps {
			kw := strings.TrimRight(s.Keyword, " ")
			if strings.EqualFold(kw, "given") || strings.EqualFold(kw, "and") {
				bgSteps = append(bgSteps, s)
			}
		}
		if len(bgSteps) > 0 {
			if enclosing := v.enclosingDescribe(); enclosing != nil {
				enclosing.background = bgSteps
			}
		}
	}
}

// relaxed reports whether the caller enabled relaxed mode (scenarios without
// Gherkin step comments are still emitted).
func (v *specVisitor) relaxed() bool { return v.relaxedMode }

// enclosingDescribe returns the nearest open describe frame, or nil.
func (v *specVisitor) enclosingDescribe() *frame {
	if len(v.describeStack) == 0 {
		return nil
	}
	return v.describeStack[len(v.describeStack)-1]
}

// stepsFromBlock scans hidden-channel comment tokens that fall inside the
// given block's token range and turns each matching Gherkin comment
// (`// Given ...`, `// When ...`, ...) into a *messages.Step.
func (v *specVisitor) stepsFromBlock(b specantlr.IBlockContext) []*messages.Step {
	if b == nil || v.tokens == nil {
		return nil
	}
	start := b.GetStart()
	stop := b.GetStop()
	if start == nil || stop == nil {
		return nil
	}
	from := start.GetTokenIndex()
	to := stop.GetTokenIndex()
	if from < 0 || to < from {
		return nil
	}

	var steps []*messages.Step
	for i := from; i <= to; i++ {
		tok := v.tokens.Get(i)
		if tok == nil {
			continue
		}
		if tok.GetChannel() != antlr.TokenHiddenChannel {
			continue
		}
		if tok.GetTokenType() != specantlr.SpecGrammarLexerLINE_COMMENT {
			continue
		}
		if kw, txt, ok := matchStepComment(tok.GetText()); ok {
			steps = append(steps, &messages.Step{Keyword: kw, Text: txt})
		}
	}
	return steps
}

// suiteKind maps a suiteKeyword parse node to the internal frame kind used
// by the emitter.
func suiteKind(k specantlr.ISuiteKeywordContext) string {
	if k == nil {
		return ""
	}
	if k.DESCRIBE() != nil {
		return frameDescribe
	}
	if k.IT() != nil {
		return frameIt
	}
	if k.BEFOREEACH() != nil {
		return frameBeforeEach
	}
	return ""
}

// extractNameAndBlock returns the first string-literal argument (with quotes
// stripped) and the first callback block found in the suite's arguments.
func extractNameAndBlock(args []specantlr.ISuiteArgContext) (string, specantlr.IBlockContext) {
	var name string
	var body specantlr.IBlockContext
	for _, a := range args {
		if a == nil {
			continue
		}
		if name == "" {
			if s := a.STRING(); s != nil {
				name = stripStringQuotes(s.GetText())
			}
		}
		if body == nil {
			if b := a.Block(); b != nil {
				body = b
			}
		}
	}
	return name, body
}

// stripStringQuotes removes the surrounding quotes (single, double, or
// backtick) from a raw STRING token's text and unescapes the common
// backslash-escape sequences we care about.
func stripStringQuotes(raw string) string {
	if len(raw) < 2 {
		return raw
	}
	q := raw[0]
	if q != '\'' && q != '"' && q != '`' {
		return raw
	}
	if raw[len(raw)-1] != q {
		return raw
	}
	inner := raw[1 : len(raw)-1]
	// Minimal unescape: only the escaped-quote / backslash pair, which is all
	// the current test data exercises.
	inner = strings.ReplaceAll(inner, "\\"+string(q), string(q))
	inner = strings.ReplaceAll(inner, "\\\\", "\\")
	return inner
}
