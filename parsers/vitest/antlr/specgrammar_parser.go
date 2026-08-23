// Code generated from SpecGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package specantlr // SpecGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SpecGrammarParser struct {
	*antlr.BaseParser
}

var SpecGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func specgrammarParserInit() {
	staticData := &SpecGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'describe'", "'it'", "'test'", "'beforeEach'", "'afterEach'", "'beforeAll'",
		"'afterAll'", "'('", "')'", "'{'", "'}'", "','", "'=>'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "DESCRIBE", "IT", "TEST", "BEFOREEACH", "AFTEREACH", "BEFOREALL",
		"AFTERALL", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "COMMA", "ARROW",
		"DOT", "STRING", "LINE_COMMENT", "BLOCK_COMMENT", "WS", "IDENT", "OTHER",
	}
	staticData.RuleNames = []string{
		"program", "element", "suite", "modifier", "dataGroup", "suiteKeyword",
		"suiteArg", "block", "parenGroup", "parenContent", "fillerToken", "argFillerToken",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 100, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 37, 8, 1, 1, 2, 1, 2, 5, 2, 41, 8, 2,
		10, 2, 12, 2, 44, 9, 2, 1, 2, 3, 2, 47, 8, 2, 1, 2, 1, 2, 5, 2, 51, 8,
		2, 10, 2, 12, 2, 54, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 70, 8, 6, 1, 7, 1, 7, 5, 7,
		74, 8, 7, 10, 7, 12, 7, 77, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 5, 8, 83, 8,
		8, 10, 8, 12, 8, 86, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 94,
		8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 0, 0, 12, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 0, 3, 1, 0, 1, 7, 3, 0, 1, 7, 12, 15, 19, 20, 3, 0,
		1, 7, 12, 14, 19, 20, 103, 0, 27, 1, 0, 0, 0, 2, 36, 1, 0, 0, 0, 4, 38,
		1, 0, 0, 0, 6, 57, 1, 0, 0, 0, 8, 60, 1, 0, 0, 0, 10, 62, 1, 0, 0, 0, 12,
		69, 1, 0, 0, 0, 14, 71, 1, 0, 0, 0, 16, 80, 1, 0, 0, 0, 18, 93, 1, 0, 0,
		0, 20, 95, 1, 0, 0, 0, 22, 97, 1, 0, 0, 0, 24, 26, 3, 2, 1, 0, 25, 24,
		1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0,
		28, 30, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 31, 5, 0, 0, 1, 31, 1, 1, 0,
		0, 0, 32, 37, 3, 4, 2, 0, 33, 37, 3, 14, 7, 0, 34, 37, 3, 16, 8, 0, 35,
		37, 3, 20, 10, 0, 36, 32, 1, 0, 0, 0, 36, 33, 1, 0, 0, 0, 36, 34, 1, 0,
		0, 0, 36, 35, 1, 0, 0, 0, 37, 3, 1, 0, 0, 0, 38, 42, 3, 10, 5, 0, 39, 41,
		3, 6, 3, 0, 40, 39, 1, 0, 0, 0, 41, 44, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0,
		42, 43, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 45, 47, 3,
		8, 4, 0, 46, 45, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48,
		52, 5, 8, 0, 0, 49, 51, 3, 12, 6, 0, 50, 49, 1, 0, 0, 0, 51, 54, 1, 0,
		0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 55, 1, 0, 0, 0, 54, 52,
		1, 0, 0, 0, 55, 56, 5, 9, 0, 0, 56, 5, 1, 0, 0, 0, 57, 58, 5, 14, 0, 0,
		58, 59, 5, 19, 0, 0, 59, 7, 1, 0, 0, 0, 60, 61, 3, 16, 8, 0, 61, 9, 1,
		0, 0, 0, 62, 63, 7, 0, 0, 0, 63, 11, 1, 0, 0, 0, 64, 70, 5, 15, 0, 0, 65,
		70, 3, 14, 7, 0, 66, 70, 3, 4, 2, 0, 67, 70, 3, 16, 8, 0, 68, 70, 3, 22,
		11, 0, 69, 64, 1, 0, 0, 0, 69, 65, 1, 0, 0, 0, 69, 66, 1, 0, 0, 0, 69,
		67, 1, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 13, 1, 0, 0, 0, 71, 75, 5, 10,
		0, 0, 72, 74, 3, 2, 1, 0, 73, 72, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73,
		1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 78, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0,
		78, 79, 5, 11, 0, 0, 79, 15, 1, 0, 0, 0, 80, 84, 5, 8, 0, 0, 81, 83, 3,
		18, 9, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84,
		85, 1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 88, 5, 9, 0,
		0, 88, 17, 1, 0, 0, 0, 89, 94, 3, 14, 7, 0, 90, 94, 3, 16, 8, 0, 91, 94,
		5, 15, 0, 0, 92, 94, 3, 22, 11, 0, 93, 89, 1, 0, 0, 0, 93, 90, 1, 0, 0,
		0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 19, 1, 0, 0, 0, 95, 96,
		7, 1, 0, 0, 96, 21, 1, 0, 0, 0, 97, 98, 7, 2, 0, 0, 98, 23, 1, 0, 0, 0,
		9, 27, 36, 42, 46, 52, 69, 75, 84, 93,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SpecGrammarParserInit initializes any static state used to implement SpecGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSpecGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SpecGrammarParserInit() {
	staticData := &SpecGrammarParserStaticData
	staticData.once.Do(specgrammarParserInit)
}

// NewSpecGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSpecGrammarParser(input antlr.TokenStream) *SpecGrammarParser {
	SpecGrammarParserInit()
	this := new(SpecGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SpecGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SpecGrammar.g4"

	return this
}

// SpecGrammarParser tokens.
const (
	SpecGrammarParserEOF           = antlr.TokenEOF
	SpecGrammarParserDESCRIBE      = 1
	SpecGrammarParserIT            = 2
	SpecGrammarParserTEST          = 3
	SpecGrammarParserBEFOREEACH    = 4
	SpecGrammarParserAFTEREACH     = 5
	SpecGrammarParserBEFOREALL     = 6
	SpecGrammarParserAFTERALL      = 7
	SpecGrammarParserLPAREN        = 8
	SpecGrammarParserRPAREN        = 9
	SpecGrammarParserLBRACE        = 10
	SpecGrammarParserRBRACE        = 11
	SpecGrammarParserCOMMA         = 12
	SpecGrammarParserARROW         = 13
	SpecGrammarParserDOT           = 14
	SpecGrammarParserSTRING        = 15
	SpecGrammarParserLINE_COMMENT  = 16
	SpecGrammarParserBLOCK_COMMENT = 17
	SpecGrammarParserWS            = 18
	SpecGrammarParserIDENT         = 19
	SpecGrammarParserOTHER         = 20
)

// SpecGrammarParser rules.
const (
	SpecGrammarParserRULE_program        = 0
	SpecGrammarParserRULE_element        = 1
	SpecGrammarParserRULE_suite          = 2
	SpecGrammarParserRULE_modifier       = 3
	SpecGrammarParserRULE_dataGroup      = 4
	SpecGrammarParserRULE_suiteKeyword   = 5
	SpecGrammarParserRULE_suiteArg       = 6
	SpecGrammarParserRULE_block          = 7
	SpecGrammarParserRULE_parenGroup     = 8
	SpecGrammarParserRULE_parenContent   = 9
	SpecGrammarParserRULE_fillerToken    = 10
	SpecGrammarParserRULE_argFillerToken = 11
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllElement() []IElementContext
	Element(i int) IElementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserEOF, 0)
}

func (s *ProgramContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SpecGrammarParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1635838) != 0 {
		{
			p.SetState(24)
			p.Element()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
		p.Match(SpecGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Suite() ISuiteContext
	Block() IBlockContext
	ParenGroup() IParenGroupContext
	FillerToken() IFillerTokenContext

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_element
	return p
}

func InitEmptyElementContext(p *ElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_element
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) Suite() ISuiteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISuiteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISuiteContext)
}

func (s *ElementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElementContext) ParenGroup() IParenGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenGroupContext)
}

func (s *ElementContext) FillerToken() IFillerTokenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFillerTokenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFillerTokenContext)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SpecGrammarParserRULE_element)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Suite()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.ParenGroup()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(35)
			p.FillerToken()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISuiteContext is an interface to support dynamic dispatch.
type ISuiteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SuiteKeyword() ISuiteKeywordContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllModifier() []IModifierContext
	Modifier(i int) IModifierContext
	DataGroup() IDataGroupContext
	AllSuiteArg() []ISuiteArgContext
	SuiteArg(i int) ISuiteArgContext

	// IsSuiteContext differentiates from other interfaces.
	IsSuiteContext()
}

type SuiteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuiteContext() *SuiteContext {
	var p = new(SuiteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_suite
	return p
}

func InitEmptySuiteContext(p *SuiteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_suite
}

func (*SuiteContext) IsSuiteContext() {}

func NewSuiteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuiteContext {
	var p = new(SuiteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_suite

	return p
}

func (s *SuiteContext) GetParser() antlr.Parser { return s.parser }

func (s *SuiteContext) SuiteKeyword() ISuiteKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISuiteKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISuiteKeywordContext)
}

func (s *SuiteContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserLPAREN, 0)
}

func (s *SuiteContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserRPAREN, 0)
}

func (s *SuiteContext) AllModifier() []IModifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IModifierContext); ok {
			len++
		}
	}

	tst := make([]IModifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IModifierContext); ok {
			tst[i] = t.(IModifierContext)
			i++
		}
	}

	return tst
}

func (s *SuiteContext) Modifier(i int) IModifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *SuiteContext) DataGroup() IDataGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataGroupContext)
}

func (s *SuiteContext) AllSuiteArg() []ISuiteArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISuiteArgContext); ok {
			len++
		}
	}

	tst := make([]ISuiteArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISuiteArgContext); ok {
			tst[i] = t.(ISuiteArgContext)
			i++
		}
	}

	return tst
}

func (s *SuiteContext) SuiteArg(i int) ISuiteArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISuiteArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISuiteArgContext)
}

func (s *SuiteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuiteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuiteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitSuite(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) Suite() (localctx ISuiteContext) {
	localctx = NewSuiteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SpecGrammarParserRULE_suite)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.SuiteKeyword()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SpecGrammarParserDOT {
		{
			p.SetState(39)
			p.Modifier()
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(45)
			p.DataGroup()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(48)
		p.Match(SpecGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1635838) != 0 {
		{
			p.SetState(49)
			p.SuiteArg()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
		p.Match(SpecGrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModifierContext is an interface to support dynamic dispatch.
type IModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsModifierContext differentiates from other interfaces.
	IsModifierContext()
}

type ModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifierContext() *ModifierContext {
	var p = new(ModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_modifier
	return p
}

func InitEmptyModifierContext(p *ModifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_modifier
}

func (*ModifierContext) IsModifierContext() {}

func NewModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifierContext {
	var p = new(ModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_modifier

	return p
}

func (s *ModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifierContext) DOT() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserDOT, 0)
}

func (s *ModifierContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserIDENT, 0)
}

func (s *ModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) Modifier() (localctx IModifierContext) {
	localctx = NewModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SpecGrammarParserRULE_modifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(SpecGrammarParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(SpecGrammarParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataGroupContext is an interface to support dynamic dispatch.
type IDataGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParenGroup() IParenGroupContext

	// IsDataGroupContext differentiates from other interfaces.
	IsDataGroupContext()
}

type DataGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataGroupContext() *DataGroupContext {
	var p = new(DataGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_dataGroup
	return p
}

func InitEmptyDataGroupContext(p *DataGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_dataGroup
}

func (*DataGroupContext) IsDataGroupContext() {}

func NewDataGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataGroupContext {
	var p = new(DataGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_dataGroup

	return p
}

func (s *DataGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *DataGroupContext) ParenGroup() IParenGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenGroupContext)
}

func (s *DataGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitDataGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) DataGroup() (localctx IDataGroupContext) {
	localctx = NewDataGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SpecGrammarParserRULE_dataGroup)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.ParenGroup()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISuiteKeywordContext is an interface to support dynamic dispatch.
type ISuiteKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DESCRIBE() antlr.TerminalNode
	IT() antlr.TerminalNode
	TEST() antlr.TerminalNode
	BEFOREEACH() antlr.TerminalNode
	AFTEREACH() antlr.TerminalNode
	BEFOREALL() antlr.TerminalNode
	AFTERALL() antlr.TerminalNode

	// IsSuiteKeywordContext differentiates from other interfaces.
	IsSuiteKeywordContext()
}

type SuiteKeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuiteKeywordContext() *SuiteKeywordContext {
	var p = new(SuiteKeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_suiteKeyword
	return p
}

func InitEmptySuiteKeywordContext(p *SuiteKeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_suiteKeyword
}

func (*SuiteKeywordContext) IsSuiteKeywordContext() {}

func NewSuiteKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuiteKeywordContext {
	var p = new(SuiteKeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_suiteKeyword

	return p
}

func (s *SuiteKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *SuiteKeywordContext) DESCRIBE() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserDESCRIBE, 0)
}

func (s *SuiteKeywordContext) IT() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserIT, 0)
}

func (s *SuiteKeywordContext) TEST() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserTEST, 0)
}

func (s *SuiteKeywordContext) BEFOREEACH() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserBEFOREEACH, 0)
}

func (s *SuiteKeywordContext) AFTEREACH() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserAFTEREACH, 0)
}

func (s *SuiteKeywordContext) BEFOREALL() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserBEFOREALL, 0)
}

func (s *SuiteKeywordContext) AFTERALL() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserAFTERALL, 0)
}

func (s *SuiteKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuiteKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuiteKeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitSuiteKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) SuiteKeyword() (localctx ISuiteKeywordContext) {
	localctx = NewSuiteKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SpecGrammarParserRULE_suiteKeyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISuiteArgContext is an interface to support dynamic dispatch.
type ISuiteArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	Block() IBlockContext
	Suite() ISuiteContext
	ParenGroup() IParenGroupContext
	ArgFillerToken() IArgFillerTokenContext

	// IsSuiteArgContext differentiates from other interfaces.
	IsSuiteArgContext()
}

type SuiteArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuiteArgContext() *SuiteArgContext {
	var p = new(SuiteArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_suiteArg
	return p
}

func InitEmptySuiteArgContext(p *SuiteArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_suiteArg
}

func (*SuiteArgContext) IsSuiteArgContext() {}

func NewSuiteArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuiteArgContext {
	var p = new(SuiteArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_suiteArg

	return p
}

func (s *SuiteArgContext) GetParser() antlr.Parser { return s.parser }

func (s *SuiteArgContext) STRING() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserSTRING, 0)
}

func (s *SuiteArgContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SuiteArgContext) Suite() ISuiteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISuiteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISuiteContext)
}

func (s *SuiteArgContext) ParenGroup() IParenGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenGroupContext)
}

func (s *SuiteArgContext) ArgFillerToken() IArgFillerTokenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgFillerTokenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgFillerTokenContext)
}

func (s *SuiteArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuiteArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuiteArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitSuiteArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) SuiteArg() (localctx ISuiteArgContext) {
	localctx = NewSuiteArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SpecGrammarParserRULE_suiteArg)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Match(SpecGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Suite()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)
			p.ParenGroup()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.ArgFillerToken()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllElement() []IElementContext
	Element(i int) IElementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserRBRACE, 0)
}

func (s *BlockContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SpecGrammarParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(SpecGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1635838) != 0 {
		{
			p.SetState(72)
			p.Element()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Match(SpecGrammarParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParenGroupContext is an interface to support dynamic dispatch.
type IParenGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllParenContent() []IParenContentContext
	ParenContent(i int) IParenContentContext

	// IsParenGroupContext differentiates from other interfaces.
	IsParenGroupContext()
}

type ParenGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenGroupContext() *ParenGroupContext {
	var p = new(ParenGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_parenGroup
	return p
}

func InitEmptyParenGroupContext(p *ParenGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_parenGroup
}

func (*ParenGroupContext) IsParenGroupContext() {}

func NewParenGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParenGroupContext {
	var p = new(ParenGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_parenGroup

	return p
}

func (s *ParenGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *ParenGroupContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserLPAREN, 0)
}

func (s *ParenGroupContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserRPAREN, 0)
}

func (s *ParenGroupContext) AllParenContent() []IParenContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParenContentContext); ok {
			len++
		}
	}

	tst := make([]IParenContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParenContentContext); ok {
			tst[i] = t.(IParenContentContext)
			i++
		}
	}

	return tst
}

func (s *ParenGroupContext) ParenContent(i int) IParenContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenContentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenContentContext)
}

func (s *ParenGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParenGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitParenGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) ParenGroup() (localctx IParenGroupContext) {
	localctx = NewParenGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SpecGrammarParserRULE_parenGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(SpecGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1635838) != 0 {
		{
			p.SetState(81)
			p.ParenContent()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(SpecGrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParenContentContext is an interface to support dynamic dispatch.
type IParenContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	ParenGroup() IParenGroupContext
	STRING() antlr.TerminalNode
	ArgFillerToken() IArgFillerTokenContext

	// IsParenContentContext differentiates from other interfaces.
	IsParenContentContext()
}

type ParenContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenContentContext() *ParenContentContext {
	var p = new(ParenContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_parenContent
	return p
}

func InitEmptyParenContentContext(p *ParenContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_parenContent
}

func (*ParenContentContext) IsParenContentContext() {}

func NewParenContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParenContentContext {
	var p = new(ParenContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_parenContent

	return p
}

func (s *ParenContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ParenContentContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ParenContentContext) ParenGroup() IParenGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenGroupContext)
}

func (s *ParenContentContext) STRING() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserSTRING, 0)
}

func (s *ParenContentContext) ArgFillerToken() IArgFillerTokenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgFillerTokenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgFillerTokenContext)
}

func (s *ParenContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParenContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitParenContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) ParenContent() (localctx IParenContentContext) {
	localctx = NewParenContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SpecGrammarParserRULE_parenContent)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SpecGrammarParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Block()
		}

	case SpecGrammarParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.ParenGroup()
		}

	case SpecGrammarParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.Match(SpecGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SpecGrammarParserDESCRIBE, SpecGrammarParserIT, SpecGrammarParserTEST, SpecGrammarParserBEFOREEACH, SpecGrammarParserAFTEREACH, SpecGrammarParserBEFOREALL, SpecGrammarParserAFTERALL, SpecGrammarParserCOMMA, SpecGrammarParserARROW, SpecGrammarParserDOT, SpecGrammarParserIDENT, SpecGrammarParserOTHER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(92)
			p.ArgFillerToken()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFillerTokenContext is an interface to support dynamic dispatch.
type IFillerTokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	ARROW() antlr.TerminalNode
	DOT() antlr.TerminalNode
	DESCRIBE() antlr.TerminalNode
	IT() antlr.TerminalNode
	TEST() antlr.TerminalNode
	BEFOREEACH() antlr.TerminalNode
	AFTEREACH() antlr.TerminalNode
	BEFOREALL() antlr.TerminalNode
	AFTERALL() antlr.TerminalNode
	OTHER() antlr.TerminalNode

	// IsFillerTokenContext differentiates from other interfaces.
	IsFillerTokenContext()
}

type FillerTokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFillerTokenContext() *FillerTokenContext {
	var p = new(FillerTokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_fillerToken
	return p
}

func InitEmptyFillerTokenContext(p *FillerTokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_fillerToken
}

func (*FillerTokenContext) IsFillerTokenContext() {}

func NewFillerTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FillerTokenContext {
	var p = new(FillerTokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_fillerToken

	return p
}

func (s *FillerTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *FillerTokenContext) STRING() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserSTRING, 0)
}

func (s *FillerTokenContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserIDENT, 0)
}

func (s *FillerTokenContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserCOMMA, 0)
}

func (s *FillerTokenContext) ARROW() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserARROW, 0)
}

func (s *FillerTokenContext) DOT() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserDOT, 0)
}

func (s *FillerTokenContext) DESCRIBE() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserDESCRIBE, 0)
}

func (s *FillerTokenContext) IT() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserIT, 0)
}

func (s *FillerTokenContext) TEST() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserTEST, 0)
}

func (s *FillerTokenContext) BEFOREEACH() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserBEFOREEACH, 0)
}

func (s *FillerTokenContext) AFTEREACH() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserAFTEREACH, 0)
}

func (s *FillerTokenContext) BEFOREALL() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserBEFOREALL, 0)
}

func (s *FillerTokenContext) AFTERALL() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserAFTERALL, 0)
}

func (s *FillerTokenContext) OTHER() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserOTHER, 0)
}

func (s *FillerTokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FillerTokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FillerTokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitFillerToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) FillerToken() (localctx IFillerTokenContext) {
	localctx = NewFillerTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SpecGrammarParserRULE_fillerToken)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1634558) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgFillerTokenContext is an interface to support dynamic dispatch.
type IArgFillerTokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	ARROW() antlr.TerminalNode
	DOT() antlr.TerminalNode
	DESCRIBE() antlr.TerminalNode
	IT() antlr.TerminalNode
	TEST() antlr.TerminalNode
	BEFOREEACH() antlr.TerminalNode
	AFTEREACH() antlr.TerminalNode
	BEFOREALL() antlr.TerminalNode
	AFTERALL() antlr.TerminalNode
	OTHER() antlr.TerminalNode

	// IsArgFillerTokenContext differentiates from other interfaces.
	IsArgFillerTokenContext()
}

type ArgFillerTokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgFillerTokenContext() *ArgFillerTokenContext {
	var p = new(ArgFillerTokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_argFillerToken
	return p
}

func InitEmptyArgFillerTokenContext(p *ArgFillerTokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SpecGrammarParserRULE_argFillerToken
}

func (*ArgFillerTokenContext) IsArgFillerTokenContext() {}

func NewArgFillerTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgFillerTokenContext {
	var p = new(ArgFillerTokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpecGrammarParserRULE_argFillerToken

	return p
}

func (s *ArgFillerTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgFillerTokenContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserIDENT, 0)
}

func (s *ArgFillerTokenContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserCOMMA, 0)
}

func (s *ArgFillerTokenContext) ARROW() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserARROW, 0)
}

func (s *ArgFillerTokenContext) DOT() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserDOT, 0)
}

func (s *ArgFillerTokenContext) DESCRIBE() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserDESCRIBE, 0)
}

func (s *ArgFillerTokenContext) IT() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserIT, 0)
}

func (s *ArgFillerTokenContext) TEST() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserTEST, 0)
}

func (s *ArgFillerTokenContext) BEFOREEACH() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserBEFOREEACH, 0)
}

func (s *ArgFillerTokenContext) AFTEREACH() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserAFTEREACH, 0)
}

func (s *ArgFillerTokenContext) BEFOREALL() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserBEFOREALL, 0)
}

func (s *ArgFillerTokenContext) AFTERALL() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserAFTERALL, 0)
}

func (s *ArgFillerTokenContext) OTHER() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserOTHER, 0)
}

func (s *ArgFillerTokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgFillerTokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgFillerTokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpecGrammarVisitor:
		return t.VisitArgFillerToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpecGrammarParser) ArgFillerToken() (localctx IArgFillerTokenContext) {
	localctx = NewArgFillerTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SpecGrammarParserRULE_argFillerToken)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1601790) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
