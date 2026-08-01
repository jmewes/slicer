// Code generated from SpecGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package specantlr // SpecGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SpecGrammarParser.
type SpecGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SpecGrammarParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by SpecGrammarParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by SpecGrammarParser#suite.
	VisitSuite(ctx *SuiteContext) interface{}

	// Visit a parse tree produced by SpecGrammarParser#suiteKeyword.
	VisitSuiteKeyword(ctx *SuiteKeywordContext) interface{}

	// Visit a parse tree produced by SpecGrammarParser#suiteArg.
	VisitSuiteArg(ctx *SuiteArgContext) interface{}

	// Visit a parse tree produced by SpecGrammarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SpecGrammarParser#parenGroup.
	VisitParenGroup(ctx *ParenGroupContext) interface{}

	// Visit a parse tree produced by SpecGrammarParser#parenContent.
	VisitParenContent(ctx *ParenContentContext) interface{}

	// Visit a parse tree produced by SpecGrammarParser#fillerToken.
	VisitFillerToken(ctx *FillerTokenContext) interface{}

	// Visit a parse tree produced by SpecGrammarParser#argFillerToken.
	VisitArgFillerToken(ctx *ArgFillerTokenContext) interface{}
}
