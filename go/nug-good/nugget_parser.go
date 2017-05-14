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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 42, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 19, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 7, 4, 30, 10, 4, 12, 4, 14, 4, 33, 11, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2, 2,
	38, 2, 18, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 26, 3, 2, 2, 2, 8, 34, 3,
	2, 2, 2, 10, 36, 3, 2, 2, 2, 12, 13, 5, 4, 3, 2, 13, 14, 7, 2, 2, 3, 14,
	19, 3, 2, 2, 2, 15, 16, 5, 10, 6, 2, 16, 17, 7, 2, 2, 3, 17, 19, 3, 2,
	2, 2, 18, 12, 3, 2, 2, 2, 18, 15, 3, 2, 2, 2, 19, 3, 3, 2, 2, 2, 20, 21,
	7, 8, 2, 2, 21, 22, 5, 6, 4, 2, 22, 23, 7, 10, 2, 2, 23, 24, 5, 8, 5, 2,
	24, 25, 7, 3, 2, 2, 25, 5, 3, 2, 2, 2, 26, 31, 7, 12, 2, 2, 27, 28, 7,
	4, 2, 2, 28, 30, 7, 12, 2, 2, 29, 27, 3, 2, 2, 2, 30, 33, 3, 2, 2, 2, 31,
	29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 7, 3, 2, 2, 2, 33, 31, 3, 2, 2,
	2, 34, 35, 7, 12, 2, 2, 35, 9, 3, 2, 2, 2, 36, 37, 7, 7, 2, 2, 37, 38,
	7, 5, 2, 2, 38, 39, 7, 13, 2, 2, 39, 40, 7, 6, 2, 2, 40, 11, 3, 2, 2, 2,
	4, 18, 31,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "','", "'('", "')'", "'sin'", "'load'", "'select'", "'from'",
}
var symbolicNames = []string{
	"", "", "", "", "", "SIN", "LOAD", "SELECT", "FROM", "WS", "Id", "NUMBER",
	"ROOT", "ATTR_LIST",
}

var ruleNames = []string{
	"nugget", "load_stat", "field", "sourceidentifier", "sin",
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
	NuggetParserT__2      = 3
	NuggetParserT__3      = 4
	NuggetParserSIN       = 5
	NuggetParserLOAD      = 6
	NuggetParserSELECT    = 7
	NuggetParserFROM      = 8
	NuggetParserWS        = 9
	NuggetParserId        = 10
	NuggetParserNUMBER    = 11
	NuggetParserROOT      = 12
	NuggetParserATTR_LIST = 13
)

// NuggetParser rules.
const (
	NuggetParserRULE_nugget           = 0
	NuggetParserRULE_load_stat        = 1
	NuggetParserRULE_field            = 2
	NuggetParserRULE_sourceidentifier = 3
	NuggetParserRULE_sin              = 4
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

func (s *NuggetContext) Sin() ISinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISinContext)
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

	p.SetState(16)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NuggetParserLOAD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.Load_stat()
		}
		{
			p.SetState(11)
			p.Match(NuggetParserEOF)
		}

	case NuggetParserSIN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.Sin()
		}
		{
			p.SetState(14)
			p.Match(NuggetParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
		p.SetState(18)
		p.Match(NuggetParserLOAD)
	}
	{
		p.SetState(19)
		p.Field()
	}
	{
		p.SetState(20)
		p.Match(NuggetParserFROM)
	}
	{
		p.SetState(21)
		p.Sourceidentifier()
	}
	{
		p.SetState(22)
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
		p.SetState(24)
		p.Match(NuggetParserId)
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NuggetParserT__1 {
		{
			p.SetState(25)
			p.Match(NuggetParserT__1)
		}
		{
			p.SetState(26)
			p.Match(NuggetParserId)
		}

		p.SetState(31)
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
		p.SetState(32)
		p.Match(NuggetParserId)
	}

	return localctx
}

// ISinContext is an interface to support dynamic dispatch.
type ISinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSinContext differentiates from other interfaces.
	IsSinContext()
}

type SinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySinContext() *SinContext {
	var p = new(SinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_sin
	return p
}

func (*SinContext) IsSinContext() {}

func NewSinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SinContext {
	var p = new(SinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_sin

	return p
}

func (s *SinContext) GetParser() antlr.Parser { return s.parser }

func (s *SinContext) SIN() antlr.TerminalNode {
	return s.GetToken(NuggetParserSIN, 0)
}

func (s *SinContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NuggetParserNUMBER, 0)
}

func (s *SinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterSin(s)
	}
}

func (s *SinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitSin(s)
	}
}

func (p *NuggetParser) Sin() (localctx ISinContext) {
	localctx = NewSinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NuggetParserRULE_sin)

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
		p.SetState(34)
		p.Match(NuggetParserSIN)
	}
	{
		p.SetState(35)
		p.Match(NuggetParserT__2)
	}
	{
		p.SetState(36)
		p.Match(NuggetParserNUMBER)
	}
	{
		p.SetState(37)
		p.Match(NuggetParserT__3)
	}

	return localctx
}
