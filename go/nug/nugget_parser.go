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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 96, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 36, 10, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 56, 10, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 66, 10, 7, 3, 8, 3, 8, 3, 8, 7, 8,
	71, 10, 8, 12, 8, 14, 8, 74, 11, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 80, 10,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 87, 10, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 2, 2, 12, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 2, 4, 4, 2, 14, 14, 17, 19, 4, 2, 15, 16, 20, 20, 2, 94, 2, 35,
	3, 2, 2, 2, 4, 37, 3, 2, 2, 2, 6, 55, 3, 2, 2, 2, 8, 57, 3, 2, 2, 2, 10,
	59, 3, 2, 2, 2, 12, 65, 3, 2, 2, 2, 14, 79, 3, 2, 2, 2, 16, 86, 3, 2, 2,
	2, 18, 88, 3, 2, 2, 2, 20, 90, 3, 2, 2, 2, 22, 23, 5, 20, 11, 2, 23, 24,
	7, 2, 2, 3, 24, 36, 3, 2, 2, 2, 25, 26, 5, 6, 4, 2, 26, 27, 7, 2, 2, 3,
	27, 36, 3, 2, 2, 2, 28, 29, 5, 4, 3, 2, 29, 30, 7, 2, 2, 3, 30, 36, 3,
	2, 2, 2, 31, 32, 5, 18, 10, 2, 32, 33, 7, 2, 2, 3, 33, 36, 3, 2, 2, 2,
	34, 36, 7, 2, 2, 3, 35, 22, 3, 2, 2, 2, 35, 25, 3, 2, 2, 2, 35, 28, 3,
	2, 2, 2, 35, 31, 3, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 3, 3, 2, 2, 2, 37,
	38, 7, 22, 2, 2, 38, 39, 7, 3, 2, 2, 39, 40, 5, 12, 7, 2, 40, 41, 7, 20,
	2, 2, 41, 42, 5, 8, 5, 2, 42, 5, 3, 2, 2, 2, 43, 44, 7, 22, 2, 2, 44, 45,
	7, 3, 2, 2, 45, 46, 3, 2, 2, 2, 46, 47, 5, 10, 6, 2, 47, 48, 5, 16, 9,
	2, 48, 49, 7, 4, 2, 2, 49, 56, 3, 2, 2, 2, 50, 51, 7, 22, 2, 2, 51, 52,
	7, 5, 2, 2, 52, 53, 5, 10, 6, 2, 53, 54, 7, 4, 2, 2, 54, 56, 3, 2, 2, 2,
	55, 43, 3, 2, 2, 2, 55, 50, 3, 2, 2, 2, 56, 7, 3, 2, 2, 2, 57, 58, 9, 2,
	2, 2, 58, 9, 3, 2, 2, 2, 59, 60, 9, 3, 2, 2, 60, 11, 3, 2, 2, 2, 61, 62,
	7, 6, 2, 2, 62, 63, 7, 23, 2, 2, 63, 66, 7, 6, 2, 2, 64, 66, 7, 23, 2,
	2, 65, 61, 3, 2, 2, 2, 65, 64, 3, 2, 2, 2, 66, 13, 3, 2, 2, 2, 67, 72,
	7, 22, 2, 2, 68, 69, 7, 7, 2, 2, 69, 71, 7, 22, 2, 2, 70, 68, 3, 2, 2,
	2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 80,
	3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 76, 7, 8, 2, 2, 76, 77, 5, 14, 8, 2,
	77, 78, 7, 8, 2, 2, 78, 80, 3, 2, 2, 2, 79, 67, 3, 2, 2, 2, 79, 75, 3,
	2, 2, 2, 80, 15, 3, 2, 2, 2, 81, 87, 7, 22, 2, 2, 82, 83, 7, 8, 2, 2, 83,
	84, 5, 16, 9, 2, 84, 85, 7, 8, 2, 2, 85, 87, 3, 2, 2, 2, 86, 81, 3, 2,
	2, 2, 86, 82, 3, 2, 2, 2, 87, 17, 3, 2, 2, 2, 88, 89, 7, 22, 2, 2, 89,
	19, 3, 2, 2, 2, 90, 91, 7, 11, 2, 2, 91, 92, 7, 9, 2, 2, 92, 93, 7, 24,
	2, 2, 93, 94, 7, 10, 2, 2, 94, 21, 3, 2, 2, 2, 8, 35, 55, 65, 72, 79, 86,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "';'", "'|'", "'\"'", "','", "'''", "'('", "')'", "'sin'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "SIN", "LOAD", "FROM", "ALL", "HASH",
	"SELECT", "FILES", "IMAGES", "DOCUMENTS", "EXTRACT", "WS", "Id", "Source",
	"NUMBER", "ROOT", "ATTR_LIST",
}

var ruleNames = []string{
	"nugget", "initextract", "execute", "subtype", "task", "target", "field",
	"sourceidentifier", "printId", "sin",
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
	NuggetParserT__4      = 5
	NuggetParserT__5      = 6
	NuggetParserT__6      = 7
	NuggetParserT__7      = 8
	NuggetParserSIN       = 9
	NuggetParserLOAD      = 10
	NuggetParserFROM      = 11
	NuggetParserALL       = 12
	NuggetParserHASH      = 13
	NuggetParserSELECT    = 14
	NuggetParserFILES     = 15
	NuggetParserIMAGES    = 16
	NuggetParserDOCUMENTS = 17
	NuggetParserEXTRACT   = 18
	NuggetParserWS        = 19
	NuggetParserId        = 20
	NuggetParserSource    = 21
	NuggetParserNUMBER    = 22
	NuggetParserROOT      = 23
	NuggetParserATTR_LIST = 24
)

// NuggetParser rules.
const (
	NuggetParserRULE_nugget           = 0
	NuggetParserRULE_initextract      = 1
	NuggetParserRULE_execute          = 2
	NuggetParserRULE_subtype          = 3
	NuggetParserRULE_task             = 4
	NuggetParserRULE_target           = 5
	NuggetParserRULE_field            = 6
	NuggetParserRULE_sourceidentifier = 7
	NuggetParserRULE_printId          = 8
	NuggetParserRULE_sin              = 9
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

func (s *NuggetContext) Sin() ISinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISinContext)
}

func (s *NuggetContext) EOF() antlr.TerminalNode {
	return s.GetToken(NuggetParserEOF, 0)
}

func (s *NuggetContext) Execute() IExecuteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecuteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecuteContext)
}

func (s *NuggetContext) Initextract() IInitextractContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitextractContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitextractContext)
}

func (s *NuggetContext) PrintId() IPrintIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintIdContext)
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

	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Sin()
		}
		{
			p.SetState(21)
			p.Match(NuggetParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Execute()
		}
		{
			p.SetState(24)
			p.Match(NuggetParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Initextract()
		}
		{
			p.SetState(27)
			p.Match(NuggetParserEOF)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(29)
			p.PrintId()
		}
		{
			p.SetState(30)
			p.Match(NuggetParserEOF)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(32)
			p.Match(NuggetParserEOF)
		}

	}

	return localctx
}

// IInitextractContext is an interface to support dynamic dispatch.
type IInitextractContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitextractContext differentiates from other interfaces.
	IsInitextractContext()
}

type InitextractContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitextractContext() *InitextractContext {
	var p = new(InitextractContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_initextract
	return p
}

func (*InitextractContext) IsInitextractContext() {}

func NewInitextractContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitextractContext {
	var p = new(InitextractContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_initextract

	return p
}

func (s *InitextractContext) GetParser() antlr.Parser { return s.parser }

func (s *InitextractContext) Id() antlr.TerminalNode {
	return s.GetToken(NuggetParserId, 0)
}

func (s *InitextractContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *InitextractContext) EXTRACT() antlr.TerminalNode {
	return s.GetToken(NuggetParserEXTRACT, 0)
}

func (s *InitextractContext) Subtype() ISubtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubtypeContext)
}

func (s *InitextractContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitextractContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitextractContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterInitextract(s)
	}
}

func (s *InitextractContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitInitextract(s)
	}
}

func (p *NuggetParser) Initextract() (localctx IInitextractContext) {
	localctx = NewInitextractContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NuggetParserRULE_initextract)

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
		p.SetState(35)
		p.Match(NuggetParserId)
	}
	{
		p.SetState(36)
		p.Match(NuggetParserT__0)
	}
	{
		p.SetState(37)
		p.Target()
	}
	{
		p.SetState(38)
		p.Match(NuggetParserEXTRACT)
	}
	{
		p.SetState(39)
		p.Subtype()
	}

	return localctx
}

// IExecuteContext is an interface to support dynamic dispatch.
type IExecuteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecuteContext differentiates from other interfaces.
	IsExecuteContext()
}

type ExecuteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecuteContext() *ExecuteContext {
	var p = new(ExecuteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_execute
	return p
}

func (*ExecuteContext) IsExecuteContext() {}

func NewExecuteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecuteContext {
	var p = new(ExecuteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_execute

	return p
}

func (s *ExecuteContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecuteContext) Task() ITaskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITaskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITaskContext)
}

func (s *ExecuteContext) Sourceidentifier() ISourceidentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceidentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceidentifierContext)
}

func (s *ExecuteContext) Id() antlr.TerminalNode {
	return s.GetToken(NuggetParserId, 0)
}

func (s *ExecuteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecuteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecuteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterExecute(s)
	}
}

func (s *ExecuteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitExecute(s)
	}
}

func (p *NuggetParser) Execute() (localctx IExecuteContext) {
	localctx = NewExecuteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NuggetParserRULE_execute)

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

	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(NuggetParserId)
		}
		{
			p.SetState(42)
			p.Match(NuggetParserT__0)
		}

		{
			p.SetState(44)
			p.Task()
		}
		{
			p.SetState(45)
			p.Sourceidentifier()
		}
		{
			p.SetState(46)
			p.Match(NuggetParserT__1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Match(NuggetParserId)
		}
		{
			p.SetState(49)
			p.Match(NuggetParserT__2)
		}
		{
			p.SetState(50)
			p.Task()
		}
		{
			p.SetState(51)
			p.Match(NuggetParserT__1)
		}

	}

	return localctx
}

// ISubtypeContext is an interface to support dynamic dispatch.
type ISubtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubtypeContext differentiates from other interfaces.
	IsSubtypeContext()
}

type SubtypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubtypeContext() *SubtypeContext {
	var p = new(SubtypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_subtype
	return p
}

func (*SubtypeContext) IsSubtypeContext() {}

func NewSubtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubtypeContext {
	var p = new(SubtypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_subtype

	return p
}

func (s *SubtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SubtypeContext) FILES() antlr.TerminalNode {
	return s.GetToken(NuggetParserFILES, 0)
}

func (s *SubtypeContext) IMAGES() antlr.TerminalNode {
	return s.GetToken(NuggetParserIMAGES, 0)
}

func (s *SubtypeContext) DOCUMENTS() antlr.TerminalNode {
	return s.GetToken(NuggetParserDOCUMENTS, 0)
}

func (s *SubtypeContext) ALL() antlr.TerminalNode {
	return s.GetToken(NuggetParserALL, 0)
}

func (s *SubtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterSubtype(s)
	}
}

func (s *SubtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitSubtype(s)
	}
}

func (p *NuggetParser) Subtype() (localctx ISubtypeContext) {
	localctx = NewSubtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NuggetParserRULE_subtype)
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
	p.SetState(55)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NuggetParserALL)|(1<<NuggetParserFILES)|(1<<NuggetParserIMAGES)|(1<<NuggetParserDOCUMENTS))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ITaskContext is an interface to support dynamic dispatch.
type ITaskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTaskContext differentiates from other interfaces.
	IsTaskContext()
}

type TaskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTaskContext() *TaskContext {
	var p = new(TaskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_task
	return p
}

func (*TaskContext) IsTaskContext() {}

func NewTaskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TaskContext {
	var p = new(TaskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_task

	return p
}

func (s *TaskContext) GetParser() antlr.Parser { return s.parser }

func (s *TaskContext) HASH() antlr.TerminalNode {
	return s.GetToken(NuggetParserHASH, 0)
}

func (s *TaskContext) EXTRACT() antlr.TerminalNode {
	return s.GetToken(NuggetParserEXTRACT, 0)
}

func (s *TaskContext) SELECT() antlr.TerminalNode {
	return s.GetToken(NuggetParserSELECT, 0)
}

func (s *TaskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TaskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TaskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterTask(s)
	}
}

func (s *TaskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitTask(s)
	}
}

func (p *NuggetParser) Task() (localctx ITaskContext) {
	localctx = NewTaskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NuggetParserRULE_task)
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
	p.SetState(57)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NuggetParserHASH)|(1<<NuggetParserSELECT)|(1<<NuggetParserEXTRACT))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_target
	return p
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) Source() antlr.TerminalNode {
	return s.GetToken(NuggetParserSource, 0)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitTarget(s)
	}
}

func (p *NuggetParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NuggetParserRULE_target)

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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NuggetParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Match(NuggetParserT__3)
		}
		{
			p.SetState(60)
			p.Match(NuggetParserSource)
		}
		{
			p.SetState(61)
			p.Match(NuggetParserT__3)
		}

	case NuggetParserSource:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Match(NuggetParserSource)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *FieldContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
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
	p.EnterRule(localctx, 12, NuggetParserRULE_field)
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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NuggetParserId:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Match(NuggetParserId)
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NuggetParserT__4 {
			{
				p.SetState(66)
				p.Match(NuggetParserT__4)
			}
			{
				p.SetState(67)
				p.Match(NuggetParserId)
			}

			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case NuggetParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Match(NuggetParserT__5)
		}
		{
			p.SetState(74)
			p.Field()
		}
		{
			p.SetState(75)
			p.Match(NuggetParserT__5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *SourceidentifierContext) Sourceidentifier() ISourceidentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceidentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceidentifierContext)
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
	p.EnterRule(localctx, 14, NuggetParserRULE_sourceidentifier)

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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NuggetParserId:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Match(NuggetParserId)
		}

	case NuggetParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Match(NuggetParserT__5)
		}
		{
			p.SetState(81)
			p.Sourceidentifier()
		}
		{
			p.SetState(82)
			p.Match(NuggetParserT__5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrintIdContext is an interface to support dynamic dispatch.
type IPrintIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintIdContext differentiates from other interfaces.
	IsPrintIdContext()
}

type PrintIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintIdContext() *PrintIdContext {
	var p = new(PrintIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_printId
	return p
}

func (*PrintIdContext) IsPrintIdContext() {}

func NewPrintIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintIdContext {
	var p = new(PrintIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_printId

	return p
}

func (s *PrintIdContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintIdContext) Id() antlr.TerminalNode {
	return s.GetToken(NuggetParserId, 0)
}

func (s *PrintIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterPrintId(s)
	}
}

func (s *PrintIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitPrintId(s)
	}
}

func (p *NuggetParser) PrintId() (localctx IPrintIdContext) {
	localctx = NewPrintIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NuggetParserRULE_printId)

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
		p.SetState(86)
		p.Match(NuggetParserId)
	}

	return localctx
}

// ISinContext is an interface to support dynamic dispatch.
type ISinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value attribute.
	GetValue() float64

	// SetValue sets the value attribute.
	SetValue(float64)

	// IsSinContext differentiates from other interfaces.
	IsSinContext()
}

type SinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  float64
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

func (s *SinContext) GetValue() float64 { return s.value }

func (s *SinContext) SetValue(v float64) { s.value = v }

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
	p.EnterRule(localctx, 18, NuggetParserRULE_sin)

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
		p.SetState(88)
		p.Match(NuggetParserSIN)
	}
	{
		p.SetState(89)
		p.Match(NuggetParserT__6)
	}
	{
		p.SetState(90)
		p.Match(NuggetParserNUMBER)
	}
	{
		p.SetState(91)
		p.Match(NuggetParserT__7)
	}

	return localctx
}
