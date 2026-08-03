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
	"strings"

	"github.com/antlr4-go/antlr/v4"
	messages "github.com/cucumber/messages/go/v28"
	specantlr "github.com/experimental-software/gherkin/javascript_parser/antlr"
)

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
