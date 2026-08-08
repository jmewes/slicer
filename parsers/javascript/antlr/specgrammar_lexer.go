// Code generated from SpecGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package specantlr

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SpecGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SpecGrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func specgrammarlexerLexerInit() {
	staticData := &SpecGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"DESCRIBE", "IT", "BEFOREEACH", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"COMMA", "ARROW", "STRING", "LINE_COMMENT", "BLOCK_COMMENT", "WS", "IDENT",
		"OTHER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 140, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 5, 9, 72, 8, 9, 10, 9, 12, 9, 75, 9, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 5, 9, 82, 8, 9, 10, 9, 12, 9, 85, 9, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 5, 9, 92, 8, 9, 10, 9, 12, 9, 95, 9, 9, 1, 9, 3, 9, 98, 8,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 104, 8, 10, 10, 10, 12, 10, 107,
		9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 115, 8, 11, 10,
		11, 12, 11, 118, 9, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 4, 12,
		126, 8, 12, 11, 12, 12, 12, 127, 1, 12, 1, 12, 1, 13, 1, 13, 5, 13, 134,
		8, 13, 10, 13, 12, 13, 137, 9, 13, 1, 14, 1, 14, 1, 116, 0, 15, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 1, 0, 7, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92,
		4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 2, 0, 92, 92, 96, 96, 2, 0, 10, 10,
		13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 4, 0, 36, 36, 65, 90, 95, 95, 97,
		122, 5, 0, 36, 36, 48, 57, 65, 90, 95, 95, 97, 122, 151, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 1, 31, 1, 0, 0, 0, 3, 40,
		1, 0, 0, 0, 5, 43, 1, 0, 0, 0, 7, 54, 1, 0, 0, 0, 9, 56, 1, 0, 0, 0, 11,
		58, 1, 0, 0, 0, 13, 60, 1, 0, 0, 0, 15, 62, 1, 0, 0, 0, 17, 64, 1, 0, 0,
		0, 19, 97, 1, 0, 0, 0, 21, 99, 1, 0, 0, 0, 23, 110, 1, 0, 0, 0, 25, 125,
		1, 0, 0, 0, 27, 131, 1, 0, 0, 0, 29, 138, 1, 0, 0, 0, 31, 32, 5, 100, 0,
		0, 32, 33, 5, 101, 0, 0, 33, 34, 5, 115, 0, 0, 34, 35, 5, 99, 0, 0, 35,
		36, 5, 114, 0, 0, 36, 37, 5, 105, 0, 0, 37, 38, 5, 98, 0, 0, 38, 39, 5,
		101, 0, 0, 39, 2, 1, 0, 0, 0, 40, 41, 5, 105, 0, 0, 41, 42, 5, 116, 0,
		0, 42, 4, 1, 0, 0, 0, 43, 44, 5, 98, 0, 0, 44, 45, 5, 101, 0, 0, 45, 46,
		5, 102, 0, 0, 46, 47, 5, 111, 0, 0, 47, 48, 5, 114, 0, 0, 48, 49, 5, 101,
		0, 0, 49, 50, 5, 69, 0, 0, 50, 51, 5, 97, 0, 0, 51, 52, 5, 99, 0, 0, 52,
		53, 5, 104, 0, 0, 53, 6, 1, 0, 0, 0, 54, 55, 5, 40, 0, 0, 55, 8, 1, 0,
		0, 0, 56, 57, 5, 41, 0, 0, 57, 10, 1, 0, 0, 0, 58, 59, 5, 123, 0, 0, 59,
		12, 1, 0, 0, 0, 60, 61, 5, 125, 0, 0, 61, 14, 1, 0, 0, 0, 62, 63, 5, 44,
		0, 0, 63, 16, 1, 0, 0, 0, 64, 65, 5, 61, 0, 0, 65, 66, 5, 62, 0, 0, 66,
		18, 1, 0, 0, 0, 67, 73, 5, 39, 0, 0, 68, 69, 5, 92, 0, 0, 69, 72, 9, 0,
		0, 0, 70, 72, 8, 0, 0, 0, 71, 68, 1, 0, 0, 0, 71, 70, 1, 0, 0, 0, 72, 75,
		1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0,
		75, 73, 1, 0, 0, 0, 76, 98, 5, 39, 0, 0, 77, 83, 5, 34, 0, 0, 78, 79, 5,
		92, 0, 0, 79, 82, 9, 0, 0, 0, 80, 82, 8, 1, 0, 0, 81, 78, 1, 0, 0, 0, 81,
		80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0,
		0, 84, 86, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 98, 5, 34, 0, 0, 87, 93,
		5, 96, 0, 0, 88, 89, 5, 92, 0, 0, 89, 92, 9, 0, 0, 0, 90, 92, 8, 2, 0,
		0, 91, 88, 1, 0, 0, 0, 91, 90, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91,
		1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0,
		96, 98, 5, 96, 0, 0, 97, 67, 1, 0, 0, 0, 97, 77, 1, 0, 0, 0, 97, 87, 1,
		0, 0, 0, 98, 20, 1, 0, 0, 0, 99, 100, 5, 47, 0, 0, 100, 101, 5, 47, 0,
		0, 101, 105, 1, 0, 0, 0, 102, 104, 8, 3, 0, 0, 103, 102, 1, 0, 0, 0, 104,
		107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 108,
		1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 109, 6, 10, 0, 0, 109, 22, 1, 0,
		0, 0, 110, 111, 5, 47, 0, 0, 111, 112, 5, 42, 0, 0, 112, 116, 1, 0, 0,
		0, 113, 115, 9, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116,
		117, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 119, 1, 0, 0, 0, 118, 116,
		1, 0, 0, 0, 119, 120, 5, 42, 0, 0, 120, 121, 5, 47, 0, 0, 121, 122, 1,
		0, 0, 0, 122, 123, 6, 11, 0, 0, 123, 24, 1, 0, 0, 0, 124, 126, 7, 4, 0,
		0, 125, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 6, 12, 1, 0, 130, 26,
		1, 0, 0, 0, 131, 135, 7, 5, 0, 0, 132, 134, 7, 6, 0, 0, 133, 132, 1, 0,
		0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0,
		136, 28, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 139, 9, 0, 0, 0, 139, 30,
		1, 0, 0, 0, 12, 0, 71, 73, 81, 83, 91, 93, 97, 105, 116, 127, 135, 2, 0,
		1, 0, 6, 0, 0,
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

// SpecGrammarLexerInit initializes any static state used to implement SpecGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSpecGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SpecGrammarLexerInit() {
	staticData := &SpecGrammarLexerLexerStaticData
	staticData.once.Do(specgrammarlexerLexerInit)
}

// NewSpecGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSpecGrammarLexer(input antlr.CharStream) *SpecGrammarLexer {
	SpecGrammarLexerInit()
	l := new(SpecGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SpecGrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SpecGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SpecGrammarLexer tokens.
const (
	SpecGrammarLexerDESCRIBE      = 1
	SpecGrammarLexerIT            = 2
	SpecGrammarLexerBEFOREEACH    = 3
	SpecGrammarLexerLPAREN        = 4
	SpecGrammarLexerRPAREN        = 5
	SpecGrammarLexerLBRACE        = 6
	SpecGrammarLexerRBRACE        = 7
	SpecGrammarLexerCOMMA         = 8
	SpecGrammarLexerARROW         = 9
	SpecGrammarLexerSTRING        = 10
	SpecGrammarLexerLINE_COMMENT  = 11
	SpecGrammarLexerBLOCK_COMMENT = 12
	SpecGrammarLexerWS            = 13
	SpecGrammarLexerIDENT         = 14
	SpecGrammarLexerOTHER         = 15
)
