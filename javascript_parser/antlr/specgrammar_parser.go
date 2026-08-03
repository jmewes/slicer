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
		"", "'describe'", "'it'", "'beforeEach'", "'('", "')'", "'{'", "'}'",
		"','", "'=>'",
	}
	staticData.SymbolicNames = []string{
		"", "DESCRIBE", "IT", "BEFOREEACH", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"COMMA", "ARROW", "STRING", "LINE_COMMENT", "BLOCK_COMMENT", "WS", "IDENT",
		"OTHER",
	}
	staticData.RuleNames = []string{
		"program", "element", "suite", "suiteKeyword", "suiteArg", "block",
		"parenGroup", "parenContent", "fillerToken", "argFillerToken",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 82, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 5,
		0, 22, 8, 0, 10, 0, 12, 0, 25, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		3, 1, 33, 8, 1, 1, 2, 1, 2, 1, 2, 5, 2, 38, 8, 2, 10, 2, 12, 2, 41, 9,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 52, 8, 4,
		1, 5, 1, 5, 5, 5, 56, 8, 5, 10, 5, 12, 5, 59, 9, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 5, 6, 65, 8, 6, 10, 6, 12, 6, 68, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 3, 7, 76, 8, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 0, 0, 10, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 0, 3, 1, 0, 1, 3, 3, 0, 1, 3, 8, 10, 14, 15,
		3, 0, 1, 3, 8, 9, 14, 15, 85, 0, 23, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4,
		34, 1, 0, 0, 0, 6, 44, 1, 0, 0, 0, 8, 51, 1, 0, 0, 0, 10, 53, 1, 0, 0,
		0, 12, 62, 1, 0, 0, 0, 14, 75, 1, 0, 0, 0, 16, 77, 1, 0, 0, 0, 18, 79,
		1, 0, 0, 0, 20, 22, 3, 2, 1, 0, 21, 20, 1, 0, 0, 0, 22, 25, 1, 0, 0, 0,
		23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 26, 1, 0, 0, 0, 25, 23, 1,
		0, 0, 0, 26, 27, 5, 0, 0, 1, 27, 1, 1, 0, 0, 0, 28, 33, 3, 4, 2, 0, 29,
		33, 3, 10, 5, 0, 30, 33, 3, 12, 6, 0, 31, 33, 3, 16, 8, 0, 32, 28, 1, 0,
		0, 0, 32, 29, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 31, 1, 0, 0, 0, 33, 3,
		1, 0, 0, 0, 34, 35, 3, 6, 3, 0, 35, 39, 5, 4, 0, 0, 36, 38, 3, 8, 4, 0,
		37, 36, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1,
		0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 43, 5, 5, 0, 0, 43,
		5, 1, 0, 0, 0, 44, 45, 7, 0, 0, 0, 45, 7, 1, 0, 0, 0, 46, 52, 5, 10, 0,
		0, 47, 52, 3, 10, 5, 0, 48, 52, 3, 4, 2, 0, 49, 52, 3, 12, 6, 0, 50, 52,
		3, 18, 9, 0, 51, 46, 1, 0, 0, 0, 51, 47, 1, 0, 0, 0, 51, 48, 1, 0, 0, 0,
		51, 49, 1, 0, 0, 0, 51, 50, 1, 0, 0, 0, 52, 9, 1, 0, 0, 0, 53, 57, 5, 6,
		0, 0, 54, 56, 3, 2, 1, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55,
		1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0,
		60, 61, 5, 7, 0, 0, 61, 11, 1, 0, 0, 0, 62, 66, 5, 4, 0, 0, 63, 65, 3,
		14, 7, 0, 64, 63, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66,
		67, 1, 0, 0, 0, 67, 69, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 70, 5, 5, 0,
		0, 70, 13, 1, 0, 0, 0, 71, 76, 3, 10, 5, 0, 72, 76, 3, 12, 6, 0, 73, 76,
		5, 10, 0, 0, 74, 76, 3, 18, 9, 0, 75, 71, 1, 0, 0, 0, 75, 72, 1, 0, 0,
		0, 75, 73, 1, 0, 0, 0, 75, 74, 1, 0, 0, 0, 76, 15, 1, 0, 0, 0, 77, 78,
		7, 1, 0, 0, 78, 17, 1, 0, 0, 0, 79, 80, 7, 2, 0, 0, 80, 19, 1, 0, 0, 0,
		7, 23, 32, 39, 51, 57, 66, 75,
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
	SpecGrammarParserBEFOREEACH    = 3
	SpecGrammarParserLPAREN        = 4
	SpecGrammarParserRPAREN        = 5
	SpecGrammarParserLBRACE        = 6
	SpecGrammarParserRBRACE        = 7
	SpecGrammarParserCOMMA         = 8
	SpecGrammarParserARROW         = 9
	SpecGrammarParserSTRING        = 10
	SpecGrammarParserLINE_COMMENT  = 11
	SpecGrammarParserBLOCK_COMMENT = 12
	SpecGrammarParserWS            = 13
	SpecGrammarParserIDENT         = 14
	SpecGrammarParserOTHER         = 15
)

// SpecGrammarParser rules.
const (
	SpecGrammarParserRULE_program        = 0
	SpecGrammarParserRULE_element        = 1
	SpecGrammarParserRULE_suite          = 2
	SpecGrammarParserRULE_suiteKeyword   = 3
	SpecGrammarParserRULE_suiteArg       = 4
	SpecGrammarParserRULE_block          = 5
	SpecGrammarParserRULE_parenGroup     = 6
	SpecGrammarParserRULE_parenContent   = 7
	SpecGrammarParserRULE_fillerToken    = 8
	SpecGrammarParserRULE_argFillerToken = 9
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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&51038) != 0 {
		{
			p.SetState(20)
			p.Element()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(26)
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
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Suite()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(30)
			p.ParenGroup()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(31)
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
		p.SetState(34)
		p.SuiteKeyword()
	}
	{
		p.SetState(35)
		p.Match(SpecGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&51038) != 0 {
		{
			p.SetState(36)
			p.SuiteArg()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
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

// ISuiteKeywordContext is an interface to support dynamic dispatch.
type ISuiteKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DESCRIBE() antlr.TerminalNode
	IT() antlr.TerminalNode
	BEFOREEACH() antlr.TerminalNode

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

func (s *SuiteKeywordContext) BEFOREEACH() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserBEFOREEACH, 0)
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
	p.EnterRule(localctx, 6, SpecGrammarParserRULE_suiteKeyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
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
	p.EnterRule(localctx, 8, SpecGrammarParserRULE_suiteArg)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(SpecGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.Suite()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.ParenGroup()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(50)
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
	p.EnterRule(localctx, 10, SpecGrammarParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(SpecGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&51038) != 0 {
		{
			p.SetState(54)
			p.Element()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
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
	p.EnterRule(localctx, 12, SpecGrammarParserRULE_parenGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(SpecGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&51038) != 0 {
		{
			p.SetState(63)
			p.ParenContent()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(69)
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
	p.EnterRule(localctx, 14, SpecGrammarParserRULE_parenContent)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SpecGrammarParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Block()
		}

	case SpecGrammarParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.ParenGroup()
		}

	case SpecGrammarParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.Match(SpecGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SpecGrammarParserDESCRIBE, SpecGrammarParserIT, SpecGrammarParserBEFOREEACH, SpecGrammarParserCOMMA, SpecGrammarParserARROW, SpecGrammarParserIDENT, SpecGrammarParserOTHER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)
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
	DESCRIBE() antlr.TerminalNode
	IT() antlr.TerminalNode
	BEFOREEACH() antlr.TerminalNode
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

func (s *FillerTokenContext) DESCRIBE() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserDESCRIBE, 0)
}

func (s *FillerTokenContext) IT() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserIT, 0)
}

func (s *FillerTokenContext) BEFOREEACH() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserBEFOREEACH, 0)
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
	p.EnterRule(localctx, 16, SpecGrammarParserRULE_fillerToken)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&50958) != 0) {
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
	DESCRIBE() antlr.TerminalNode
	IT() antlr.TerminalNode
	BEFOREEACH() antlr.TerminalNode
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

func (s *ArgFillerTokenContext) DESCRIBE() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserDESCRIBE, 0)
}

func (s *ArgFillerTokenContext) IT() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserIT, 0)
}

func (s *ArgFillerTokenContext) BEFOREEACH() antlr.TerminalNode {
	return s.GetToken(SpecGrammarParserBEFOREEACH, 0)
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
	p.EnterRule(localctx, 18, SpecGrammarParserRULE_argFillerToken)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&49934) != 0) {
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
