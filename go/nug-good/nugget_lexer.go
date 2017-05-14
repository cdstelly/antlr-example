// Generated from Nugget.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 78, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 6, 10, 58, 10, 10, 13,
	10, 14, 10, 59, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 67, 10, 11, 12,
	11, 14, 11, 70, 11, 11, 3, 12, 6, 12, 73, 10, 12, 13, 12, 14, 12, 74, 3,
	13, 3, 13, 2, 2, 14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 21, 12, 23, 13, 25, 2, 3, 2, 4, 5, 2, 11, 12, 15, 15, 34, 34, 5,
	2, 67, 92, 97, 97, 99, 124, 2, 80, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2,
	2, 2, 2, 23, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 29, 3, 2, 2, 2, 7, 31, 3,
	2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 35, 3, 2, 2, 2, 13, 39, 3, 2, 2, 2, 15,
	44, 3, 2, 2, 2, 17, 51, 3, 2, 2, 2, 19, 57, 3, 2, 2, 2, 21, 63, 3, 2, 2,
	2, 23, 72, 3, 2, 2, 2, 25, 76, 3, 2, 2, 2, 27, 28, 7, 61, 2, 2, 28, 4,
	3, 2, 2, 2, 29, 30, 7, 46, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 42, 2, 2,
	32, 8, 3, 2, 2, 2, 33, 34, 7, 43, 2, 2, 34, 10, 3, 2, 2, 2, 35, 36, 7,
	117, 2, 2, 36, 37, 7, 107, 2, 2, 37, 38, 7, 112, 2, 2, 38, 12, 3, 2, 2,
	2, 39, 40, 7, 110, 2, 2, 40, 41, 7, 113, 2, 2, 41, 42, 7, 99, 2, 2, 42,
	43, 7, 102, 2, 2, 43, 14, 3, 2, 2, 2, 44, 45, 7, 117, 2, 2, 45, 46, 7,
	103, 2, 2, 46, 47, 7, 110, 2, 2, 47, 48, 7, 103, 2, 2, 48, 49, 7, 101,
	2, 2, 49, 50, 7, 118, 2, 2, 50, 16, 3, 2, 2, 2, 51, 52, 7, 104, 2, 2, 52,
	53, 7, 116, 2, 2, 53, 54, 7, 113, 2, 2, 54, 55, 7, 111, 2, 2, 55, 18, 3,
	2, 2, 2, 56, 58, 9, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59,
	57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 62, 8, 10,
	2, 2, 62, 20, 3, 2, 2, 2, 63, 68, 9, 3, 2, 2, 64, 67, 9, 3, 2, 2, 65, 67,
	5, 25, 13, 2, 66, 64, 3, 2, 2, 2, 66, 65, 3, 2, 2, 2, 67, 70, 3, 2, 2,
	2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 22, 3, 2, 2, 2, 70, 68,
	3, 2, 2, 2, 71, 73, 5, 25, 13, 2, 72, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2,
	2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 24, 3, 2, 2, 2, 76, 77,
	4, 50, 59, 2, 77, 26, 3, 2, 2, 2, 7, 2, 59, 66, 68, 74, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "','", "'('", "')'", "'sin'", "'load'", "'select'", "'from'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "SIN", "LOAD", "SELECT", "FROM", "WS", "Id", "NUMBER",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "SIN", "LOAD", "SELECT", "FROM", "WS",
	"Id", "NUMBER", "DIGIT",
}

type NuggetLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewNuggetLexer(input antlr.CharStream) *NuggetLexer {

	l := new(NuggetLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Nugget.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NuggetLexer tokens.
const (
	NuggetLexerT__0   = 1
	NuggetLexerT__1   = 2
	NuggetLexerT__2   = 3
	NuggetLexerT__3   = 4
	NuggetLexerSIN    = 5
	NuggetLexerLOAD   = 6
	NuggetLexerSELECT = 7
	NuggetLexerFROM   = 8
	NuggetLexerWS     = 9
	NuggetLexerId     = 10
	NuggetLexerNUMBER = 11
)
