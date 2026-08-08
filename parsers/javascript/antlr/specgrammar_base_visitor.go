// Code generated from SpecGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package specantlr // SpecGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseSpecGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSpecGrammarVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpecGrammarVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpecGrammarVisitor) VisitSuite(ctx *SuiteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpecGrammarVisitor) VisitSuiteKeyword(ctx *SuiteKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpecGrammarVisitor) VisitSuiteArg(ctx *SuiteArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpecGrammarVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpecGrammarVisitor) VisitParenGroup(ctx *ParenGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpecGrammarVisitor) VisitParenContent(ctx *ParenContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpecGrammarVisitor) VisitFillerToken(ctx *FillerTokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpecGrammarVisitor) VisitArgFillerToken(ctx *ArgFillerTokenContext) interface{} {
	return v.VisitChildren(ctx)
}
