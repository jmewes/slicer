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
		"", "'describe'", "'it'", "'fit'", "'xit'", "'beforeEach'", "'('", "')'",
		"'{'", "'}'", "','", "'=>'",
	}
	staticData.SymbolicNames = []string{
		"", "DESCRIBE", "IT", "FIT", "XIT", "BEFOREEACH", "LPAREN", "RPAREN",
		"LBRACE", "RBRACE", "COMMA", "ARROW", "STRING", "LINE_COMMENT", "BLOCK_COMMENT",
		"WS", "IDENT", "OTHER",
	}
	staticData.RuleNames = []string{
		"DESCRIBE", "IT", "FIT", "XIT", "BEFOREEACH", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "COMMA", "ARROW", "STRING", "LINE_COMMENT", "BLOCK_COMMENT",
		"WS", "IDENT", "OTHER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 152, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 84, 8, 11, 10, 11, 12, 11, 87, 9, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 94, 8, 11, 10, 11, 12, 11, 97,
		9, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 104, 8, 11, 10, 11, 12,
		11, 107, 9, 11, 1, 11, 3, 11, 110, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5,
		12, 116, 8, 12, 10, 12, 12, 12, 119, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 5, 13, 127, 8, 13, 10, 13, 12, 13, 130, 9, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 4, 14, 138, 8, 14, 11, 14, 12, 14, 139,
		1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 146, 8, 15, 10, 15, 12, 15, 149, 9,
		15, 1, 16, 1, 16, 1, 128, 0, 17, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 1, 0, 7, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 4, 0, 10, 10, 13,
		13, 34, 34, 92, 92, 2, 0, 92, 92, 96, 96, 2, 0, 10, 10, 13, 13, 3, 0, 9,
		10, 13, 13, 32, 32, 4, 0, 36, 36, 65, 90, 95, 95, 97, 122, 5, 0, 36, 36,
		48, 57, 65, 90, 95, 95, 97, 122, 163, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 1, 35,
		1, 0, 0, 0, 3, 44, 1, 0, 0, 0, 5, 47, 1, 0, 0, 0, 7, 51, 1, 0, 0, 0, 9,
		55, 1, 0, 0, 0, 11, 66, 1, 0, 0, 0, 13, 68, 1, 0, 0, 0, 15, 70, 1, 0, 0,
		0, 17, 72, 1, 0, 0, 0, 19, 74, 1, 0, 0, 0, 21, 76, 1, 0, 0, 0, 23, 109,
		1, 0, 0, 0, 25, 111, 1, 0, 0, 0, 27, 122, 1, 0, 0, 0, 29, 137, 1, 0, 0,
		0, 31, 143, 1, 0, 0, 0, 33, 150, 1, 0, 0, 0, 35, 36, 5, 100, 0, 0, 36,
		37, 5, 101, 0, 0, 37, 38, 5, 115, 0, 0, 38, 39, 5, 99, 0, 0, 39, 40, 5,
		114, 0, 0, 40, 41, 5, 105, 0, 0, 41, 42, 5, 98, 0, 0, 42, 43, 5, 101, 0,
		0, 43, 2, 1, 0, 0, 0, 44, 45, 5, 105, 0, 0, 45, 46, 5, 116, 0, 0, 46, 4,
		1, 0, 0, 0, 47, 48, 5, 102, 0, 0, 48, 49, 5, 105, 0, 0, 49, 50, 5, 116,
		0, 0, 50, 6, 1, 0, 0, 0, 51, 52, 5, 120, 0, 0, 52, 53, 5, 105, 0, 0, 53,
		54, 5, 116, 0, 0, 54, 8, 1, 0, 0, 0, 55, 56, 5, 98, 0, 0, 56, 57, 5, 101,
		0, 0, 57, 58, 5, 102, 0, 0, 58, 59, 5, 111, 0, 0, 59, 60, 5, 114, 0, 0,
		60, 61, 5, 101, 0, 0, 61, 62, 5, 69, 0, 0, 62, 63, 5, 97, 0, 0, 63, 64,
		5, 99, 0, 0, 64, 65, 5, 104, 0, 0, 65, 10, 1, 0, 0, 0, 66, 67, 5, 40, 0,
		0, 67, 12, 1, 0, 0, 0, 68, 69, 5, 41, 0, 0, 69, 14, 1, 0, 0, 0, 70, 71,
		5, 123, 0, 0, 71, 16, 1, 0, 0, 0, 72, 73, 5, 125, 0, 0, 73, 18, 1, 0, 0,
		0, 74, 75, 5, 44, 0, 0, 75, 20, 1, 0, 0, 0, 76, 77, 5, 61, 0, 0, 77, 78,
		5, 62, 0, 0, 78, 22, 1, 0, 0, 0, 79, 85, 5, 39, 0, 0, 80, 81, 5, 92, 0,
		0, 81, 84, 9, 0, 0, 0, 82, 84, 8, 0, 0, 0, 83, 80, 1, 0, 0, 0, 83, 82,
		1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0,
		86, 88, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 110, 5, 39, 0, 0, 89, 95, 5,
		34, 0, 0, 90, 91, 5, 92, 0, 0, 91, 94, 9, 0, 0, 0, 92, 94, 8, 1, 0, 0,
		93, 90, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1,
		0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98,
		110, 5, 34, 0, 0, 99, 105, 5, 96, 0, 0, 100, 101, 5, 92, 0, 0, 101, 104,
		9, 0, 0, 0, 102, 104, 8, 2, 0, 0, 103, 100, 1, 0, 0, 0, 103, 102, 1, 0,
		0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0,
		106, 108, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 110, 5, 96, 0, 0, 109,
		79, 1, 0, 0, 0, 109, 89, 1, 0, 0, 0, 109, 99, 1, 0, 0, 0, 110, 24, 1, 0,
		0, 0, 111, 112, 5, 47, 0, 0, 112, 113, 5, 47, 0, 0, 113, 117, 1, 0, 0,
		0, 114, 116, 8, 3, 0, 0, 115, 114, 1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117,
		115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 120, 1, 0, 0, 0, 119, 117,
		1, 0, 0, 0, 120, 121, 6, 12, 0, 0, 121, 26, 1, 0, 0, 0, 122, 123, 5, 47,
		0, 0, 123, 124, 5, 42, 0, 0, 124, 128, 1, 0, 0, 0, 125, 127, 9, 0, 0, 0,
		126, 125, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 128,
		126, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 132,
		5, 42, 0, 0, 132, 133, 5, 47, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 6,
		13, 0, 0, 135, 28, 1, 0, 0, 0, 136, 138, 7, 4, 0, 0, 137, 136, 1, 0, 0,
		0, 138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140,
		141, 1, 0, 0, 0, 141, 142, 6, 14, 1, 0, 142, 30, 1, 0, 0, 0, 143, 147,
		7, 5, 0, 0, 144, 146, 7, 6, 0, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0,
		0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 32, 1, 0, 0, 0,
		149, 147, 1, 0, 0, 0, 150, 151, 9, 0, 0, 0, 151, 34, 1, 0, 0, 0, 12, 0,
		83, 85, 93, 95, 103, 105, 109, 117, 128, 139, 147, 2, 0, 1, 0, 6, 0, 0,
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
	SpecGrammarLexerXIT           = 4
	SpecGrammarLexerBEFOREEACH    = 5
	SpecGrammarLexerLPAREN        = 6
	SpecGrammarLexerRPAREN        = 7
	SpecGrammarLexerLBRACE        = 8
	SpecGrammarLexerRBRACE        = 9
	SpecGrammarLexerCOMMA         = 10
	SpecGrammarLexerARROW         = 11
	SpecGrammarLexerSTRING        = 12
	SpecGrammarLexerLINE_COMMENT  = 13
	SpecGrammarLexerBLOCK_COMMENT = 14
	SpecGrammarLexerWS            = 15
	SpecGrammarLexerIDENT         = 16
	SpecGrammarLexerOTHER         = 17
)
