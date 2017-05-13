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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 57, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 7, 6, 7, 42, 10, 7, 13, 7, 14, 7, 43, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 8, 7, 8, 51, 10, 8, 12, 8, 14, 8, 54, 11, 8, 3, 9, 3, 9, 2, 2, 10,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 2, 3, 2, 4, 5, 2, 11,
	12, 15, 15, 34, 34, 5, 2, 67, 92, 97, 97, 99, 124, 2, 58, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5, 21, 3, 2,
	2, 2, 7, 23, 3, 2, 2, 2, 9, 28, 3, 2, 2, 2, 11, 35, 3, 2, 2, 2, 13, 41,
	3, 2, 2, 2, 15, 47, 3, 2, 2, 2, 17, 55, 3, 2, 2, 2, 19, 20, 7, 61, 2, 2,
	20, 4, 3, 2, 2, 2, 21, 22, 7, 46, 2, 2, 22, 6, 3, 2, 2, 2, 23, 24, 7, 110,
	2, 2, 24, 25, 7, 113, 2, 2, 25, 26, 7, 99, 2, 2, 26, 27, 7, 102, 2, 2,
	27, 8, 3, 2, 2, 2, 28, 29, 7, 117, 2, 2, 29, 30, 7, 103, 2, 2, 30, 31,
	7, 110, 2, 2, 31, 32, 7, 103, 2, 2, 32, 33, 7, 101, 2, 2, 33, 34, 7, 118,
	2, 2, 34, 10, 3, 2, 2, 2, 35, 36, 7, 104, 2, 2, 36, 37, 7, 116, 2, 2, 37,
	38, 7, 113, 2, 2, 38, 39, 7, 111, 2, 2, 39, 12, 3, 2, 2, 2, 40, 42, 9,
	2, 2, 2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43,
	44, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 8, 7, 2, 2, 46, 14, 3, 2, 2,
	2, 47, 52, 9, 3, 2, 2, 48, 51, 9, 3, 2, 2, 49, 51, 5, 17, 9, 2, 50, 48,
	3, 2, 2, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2,
	52, 53, 3, 2, 2, 2, 53, 16, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 56, 4,
	50, 59, 2, 56, 18, 3, 2, 2, 2, 6, 2, 43, 50, 52, 3, 8, 2, 2,
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
	"", "';'", "','", "'load'", "'select'", "'from'",
}

var lexerSymbolicNames = []string{
	"", "", "", "LOAD", "SELECT", "FROM", "WS", "Id",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "LOAD", "SELECT", "FROM", "WS", "Id", "Digit",
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
	NuggetLexerLOAD   = 3
	NuggetLexerSELECT = 4
	NuggetLexerFROM   = 5
	NuggetLexerWS     = 6
	NuggetLexerId     = 7
)
