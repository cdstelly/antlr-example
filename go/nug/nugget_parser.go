// Generated from Nugget.g4 by ANTLR 4.7.

package parser // Nugget

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 30, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 23, 10, 4, 12, 4, 14,
	4, 26, 11, 4, 3, 5, 3, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 2, 2, 26, 2, 10,
	3, 2, 2, 2, 4, 13, 3, 2, 2, 2, 6, 19, 3, 2, 2, 2, 8, 27, 3, 2, 2, 2, 10,
	11, 5, 4, 3, 2, 11, 12, 7, 2, 2, 3, 12, 3, 3, 2, 2, 2, 13, 14, 7, 5, 2,
	2, 14, 15, 5, 6, 4, 2, 15, 16, 7, 7, 2, 2, 16, 17, 5, 8, 5, 2, 17, 18,
	7, 3, 2, 2, 18, 5, 3, 2, 2, 2, 19, 24, 7, 9, 2, 2, 20, 21, 7, 4, 2, 2,
	21, 23, 7, 9, 2, 2, 22, 20, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24, 22, 3,
	2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 7, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 27,
	28, 7, 9, 2, 2, 28, 9, 3, 2, 2, 2, 3, 24,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "','", "'load'", "'select'", "'from'",
}
var symbolicNames = []string{
	"", "", "", "LOAD", "SELECT", "FROM", "WS", "Id", "ROOT", "ATTR_LIST",
}

var ruleNames = []string{
	"nugget", "load_stat", "field", "sourceidentifier",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type NuggetParser struct {
	*antlr.BaseParser
}

func NewNuggetParser(input antlr.TokenStream) *NuggetParser {
	this := new(NuggetParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Nugget.g4"

	return this
}

// NuggetParser tokens.
const (
	NuggetParserEOF       = antlr.TokenEOF
	NuggetParserT__0      = 1
	NuggetParserT__1      = 2
	NuggetParserLOAD      = 3
	NuggetParserSELECT    = 4
	NuggetParserFROM      = 5
	NuggetParserWS        = 6
	NuggetParserId        = 7
	NuggetParserROOT      = 8
	NuggetParserATTR_LIST = 9
)

// NuggetParser rules.
const (
	NuggetParserRULE_nugget           = 0
	NuggetParserRULE_load_stat        = 1
	NuggetParserRULE_field            = 2
	NuggetParserRULE_sourceidentifier = 3
)

// INuggetContext is an interface to support dynamic dispatch.
type INuggetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNuggetContext differentiates from other interfaces.
	IsNuggetContext()
}

type NuggetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNuggetContext() *NuggetContext {
	var p = new(NuggetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_nugget
	return p
}

func (*NuggetContext) IsNuggetContext() {}

func NewNuggetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NuggetContext {
	var p = new(NuggetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_nugget

	return p
}

func (s *NuggetContext) GetParser() antlr.Parser { return s.parser }

func (s *NuggetContext) Load_stat() ILoad_statContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoad_statContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoad_statContext)
}

func (s *NuggetContext) EOF() antlr.TerminalNode {
	return s.GetToken(NuggetParserEOF, 0)
}

func (s *NuggetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NuggetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NuggetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterNugget(s)
	}
}

func (s *NuggetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitNugget(s)
	}
}

func (p *NuggetParser) Nugget() (localctx INuggetContext) {
	localctx = NewNuggetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NuggetParserRULE_nugget)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.Load_stat()
	}
	{
		p.SetState(9)
		p.Match(NuggetParserEOF)
	}

	return localctx
}

// ILoad_statContext is an interface to support dynamic dispatch.
type ILoad_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoad_statContext differentiates from other interfaces.
	IsLoad_statContext()
}

type Load_statContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoad_statContext() *Load_statContext {
	var p = new(Load_statContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_load_stat
	return p
}

func (*Load_statContext) IsLoad_statContext() {}

func NewLoad_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Load_statContext {
	var p = new(Load_statContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_load_stat

	return p
}

func (s *Load_statContext) GetParser() antlr.Parser { return s.parser }

func (s *Load_statContext) LOAD() antlr.TerminalNode {
	return s.GetToken(NuggetParserLOAD, 0)
}

func (s *Load_statContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Load_statContext) FROM() antlr.TerminalNode {
	return s.GetToken(NuggetParserFROM, 0)
}

func (s *Load_statContext) Sourceidentifier() ISourceidentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceidentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceidentifierContext)
}

func (s *Load_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Load_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Load_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterLoad_stat(s)
	}
}

func (s *Load_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitLoad_stat(s)
	}
}

func (p *NuggetParser) Load_stat() (localctx ILoad_statContext) {
	localctx = NewLoad_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NuggetParserRULE_load_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(11)
		p.Match(NuggetParserLOAD)
	}
	{
		p.SetState(12)
		p.Field()
	}
	{
		p.SetState(13)
		p.Match(NuggetParserFROM)
	}
	{
		p.SetState(14)
		p.Sourceidentifier()
	}
	{
		p.SetState(15)
		p.Match(NuggetParserT__0)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) AllId() []antlr.TerminalNode {
	return s.GetTokens(NuggetParserId)
}

func (s *FieldContext) Id(i int) antlr.TerminalNode {
	return s.GetToken(NuggetParserId, i)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *NuggetParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NuggetParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(17)
		p.Match(NuggetParserId)
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NuggetParserT__1 {
		{
			p.SetState(18)
			p.Match(NuggetParserT__1)
		}
		{
			p.SetState(19)
			p.Match(NuggetParserId)
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISourceidentifierContext is an interface to support dynamic dispatch.
type ISourceidentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceidentifierContext differentiates from other interfaces.
	IsSourceidentifierContext()
}

type SourceidentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceidentifierContext() *SourceidentifierContext {
	var p = new(SourceidentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_sourceidentifier
	return p
}

func (*SourceidentifierContext) IsSourceidentifierContext() {}

func NewSourceidentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceidentifierContext {
	var p = new(SourceidentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_sourceidentifier

	return p
}

func (s *SourceidentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceidentifierContext) Id() antlr.TerminalNode {
	return s.GetToken(NuggetParserId, 0)
}

func (s *SourceidentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceidentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceidentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterSourceidentifier(s)
	}
}

func (s *SourceidentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitSourceidentifier(s)
	}
}

func (p *NuggetParser) Sourceidentifier() (localctx ISourceidentifierContext) {
	localctx = NewSourceidentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NuggetParserRULE_sourceidentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(NuggetParserId)
	}

	return localctx
}
