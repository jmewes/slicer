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
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	messages "github.com/cucumber/messages/go/v28"
	specantlr "github.com/experimental-software/gherkin/parsers/javascript/antlr"
)

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
// ANTLR-generated parser (see parsers/javascript/antlr/SpecGrammar.g4) driving a small
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
