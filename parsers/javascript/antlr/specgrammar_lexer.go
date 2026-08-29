// Code generated from SpecGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"", "'describe'", "'it'", "'fit'", "'beforeEach'", "'('", "')'", "'{'",
		"'}'", "','", "'=>'",
	}
	staticData.SymbolicNames = []string{
		"", "DESCRIBE", "IT", "FIT", "BEFOREEACH", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "COMMA", "ARROW", "STRING", "LINE_COMMENT", "BLOCK_COMMENT",
		"WS", "IDENT", "OTHER",
	}
	staticData.RuleNames = []string{
		"DESCRIBE", "IT", "FIT", "BEFOREEACH", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "COMMA", "ARROW", "STRING", "LINE_COMMENT", "BLOCK_COMMENT",
		"WS", "IDENT", "OTHER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 146, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 78, 8,
		10, 10, 10, 12, 10, 81, 9, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		88, 8, 10, 10, 10, 12, 10, 91, 9, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		5, 10, 98, 8, 10, 10, 10, 12, 10, 101, 9, 10, 1, 10, 3, 10, 104, 8, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 110, 8, 11, 10, 11, 12, 11, 113, 9,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 121, 8, 12, 10, 12,
		12, 12, 124, 9, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 4, 13, 132,
		8, 13, 11, 13, 12, 13, 133, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 140, 8,
		14, 10, 14, 12, 14, 143, 9, 14, 1, 15, 1, 15, 1, 122, 0, 16, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 1, 0, 7, 4, 0, 10, 10, 13, 13, 39, 39, 92,
		92, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 2, 0, 92, 92, 96, 96, 2, 0, 10,
		10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 4, 0, 36, 36, 65, 90, 95, 95,
		97, 122, 5, 0, 36, 36, 48, 57, 65, 90, 95, 95, 97, 122, 157, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 1,
		33, 1, 0, 0, 0, 3, 42, 1, 0, 0, 0, 5, 45, 1, 0, 0, 0, 7, 49, 1, 0, 0, 0,
		9, 60, 1, 0, 0, 0, 11, 62, 1, 0, 0, 0, 13, 64, 1, 0, 0, 0, 15, 66, 1, 0,
		0, 0, 17, 68, 1, 0, 0, 0, 19, 70, 1, 0, 0, 0, 21, 103, 1, 0, 0, 0, 23,
		105, 1, 0, 0, 0, 25, 116, 1, 0, 0, 0, 27, 131, 1, 0, 0, 0, 29, 137, 1,
		0, 0, 0, 31, 144, 1, 0, 0, 0, 33, 34, 5, 100, 0, 0, 34, 35, 5, 101, 0,
		0, 35, 36, 5, 115, 0, 0, 36, 37, 5, 99, 0, 0, 37, 38, 5, 114, 0, 0, 38,
		39, 5, 105, 0, 0, 39, 40, 5, 98, 0, 0, 40, 41, 5, 101, 0, 0, 41, 2, 1,
		0, 0, 0, 42, 43, 5, 105, 0, 0, 43, 44, 5, 116, 0, 0, 44, 4, 1, 0, 0, 0,
		45, 46, 5, 102, 0, 0, 46, 47, 5, 105, 0, 0, 47, 48, 5, 116, 0, 0, 48, 6,
		1, 0, 0, 0, 49, 50, 5, 98, 0, 0, 50, 51, 5, 101, 0, 0, 51, 52, 5, 102,
		0, 0, 52, 53, 5, 111, 0, 0, 53, 54, 5, 114, 0, 0, 54, 55, 5, 101, 0, 0,
		55, 56, 5, 69, 0, 0, 56, 57, 5, 97, 0, 0, 57, 58, 5, 99, 0, 0, 58, 59,
		5, 104, 0, 0, 59, 8, 1, 0, 0, 0, 60, 61, 5, 40, 0, 0, 61, 10, 1, 0, 0,
		0, 62, 63, 5, 41, 0, 0, 63, 12, 1, 0, 0, 0, 64, 65, 5, 123, 0, 0, 65, 14,
		1, 0, 0, 0, 66, 67, 5, 125, 0, 0, 67, 16, 1, 0, 0, 0, 68, 69, 5, 44, 0,
		0, 69, 18, 1, 0, 0, 0, 70, 71, 5, 61, 0, 0, 71, 72, 5, 62, 0, 0, 72, 20,
		1, 0, 0, 0, 73, 79, 5, 39, 0, 0, 74, 75, 5, 92, 0, 0, 75, 78, 9, 0, 0,
		0, 76, 78, 8, 0, 0, 0, 77, 74, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 81,
		1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0,
		81, 79, 1, 0, 0, 0, 82, 104, 5, 39, 0, 0, 83, 89, 5, 34, 0, 0, 84, 85,
		5, 92, 0, 0, 85, 88, 9, 0, 0, 0, 86, 88, 8, 1, 0, 0, 87, 84, 1, 0, 0, 0,
		87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1,
		0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 104, 5, 34, 0, 0,
		93, 99, 5, 96, 0, 0, 94, 95, 5, 92, 0, 0, 95, 98, 9, 0, 0, 0, 96, 98, 8,
		2, 0, 0, 97, 94, 1, 0, 0, 0, 97, 96, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99,
		97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 102, 1, 0, 0, 0, 101, 99, 1,
		0, 0, 0, 102, 104, 5, 96, 0, 0, 103, 73, 1, 0, 0, 0, 103, 83, 1, 0, 0,
		0, 103, 93, 1, 0, 0, 0, 104, 22, 1, 0, 0, 0, 105, 106, 5, 47, 0, 0, 106,
		107, 5, 47, 0, 0, 107, 111, 1, 0, 0, 0, 108, 110, 8, 3, 0, 0, 109, 108,
		1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0,
		0, 0, 112, 114, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 115, 6, 11, 0, 0,
		115, 24, 1, 0, 0, 0, 116, 117, 5, 47, 0, 0, 117, 118, 5, 42, 0, 0, 118,
		122, 1, 0, 0, 0, 119, 121, 9, 0, 0, 0, 120, 119, 1, 0, 0, 0, 121, 124,
		1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 125, 1, 0,
		0, 0, 124, 122, 1, 0, 0, 0, 125, 126, 5, 42, 0, 0, 126, 127, 5, 47, 0,
		0, 127, 128, 1, 0, 0, 0, 128, 129, 6, 12, 0, 0, 129, 26, 1, 0, 0, 0, 130,
		132, 7, 4, 0, 0, 131, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 131,
		1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 6, 13,
		1, 0, 136, 28, 1, 0, 0, 0, 137, 141, 7, 5, 0, 0, 138, 140, 7, 6, 0, 0,
		139, 138, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 30, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 145, 9,
		0, 0, 0, 145, 32, 1, 0, 0, 0, 12, 0, 77, 79, 87, 89, 97, 99, 103, 111,
		122, 133, 141, 2, 0, 1, 0, 6, 0, 0,
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
	SpecGrammarLexerFIT           = 3
	SpecGrammarLexerBEFOREEACH    = 4
	SpecGrammarLexerLPAREN        = 5
	SpecGrammarLexerRPAREN        = 6
	SpecGrammarLexerLBRACE        = 7
	SpecGrammarLexerRBRACE        = 8
	SpecGrammarLexerCOMMA         = 9
	SpecGrammarLexerARROW         = 10
	SpecGrammarLexerSTRING        = 11
	SpecGrammarLexerLINE_COMMENT  = 12
	SpecGrammarLexerBLOCK_COMMENT = 13
	SpecGrammarLexerWS            = 14
	SpecGrammarLexerIDENT         = 15
	SpecGrammarLexerOTHER         = 16
)
