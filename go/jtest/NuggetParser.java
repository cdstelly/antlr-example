// Generated from Nugget.g4 by ANTLR 4.7
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class NuggetParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		CTIME=10, MTIME=11, SIN=12, LOAD=13, FROM=14, ALL=15, HASH=16, SELECT=17, 
		FILES=18, IMAGES=19, DOCUMENTS=20, EXTRACT=21, JOIN=22, WS=23, COMMENT=24, 
		NUMBER=25, StrLit=26, OPERATION=27, DATE=28, ROOT=29, ATTR_LIST=30;
	public static final int
		RULE_nugget = 0, RULE_initextract = 1, RULE_assign = 2, RULE_execute = 3, 
		RULE_streamTask = 4, RULE_save = 5, RULE_filter = 6, RULE_filename = 7, 
		RULE_timefilter = 8, RULE_reference = 9, RULE_date = 10, RULE_subtype = 11, 
		RULE_task = 12, RULE_target = 13, RULE_field = 14, RULE_sourceidentifier = 15, 
		RULE_printId = 16, RULE_sin = 17;
	public static final String[] ruleNames = {
		"nugget", "initextract", "assign", "execute", "streamTask", "save", "filter", 
		"filename", "timefilter", "reference", "date", "subtype", "task", "target", 
		"field", "sourceidentifier", "printId", "sin"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'='", "'|'", "','", "'save'", "'to'", "'\"'", "'''", "'('", "')'", 
		null, null, "'sin'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, null, null, "CTIME", "MTIME", 
		"SIN", "LOAD", "FROM", "ALL", "HASH", "SELECT", "FILES", "IMAGES", "DOCUMENTS", 
		"EXTRACT", "JOIN", "WS", "COMMENT", "NUMBER", "StrLit", "OPERATION", "DATE", 
		"ROOT", "ATTR_LIST"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Nugget.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public NuggetParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class NuggetContext extends ParserRuleContext {
		public SinContext sin() {
			return getRuleContext(SinContext.class,0);
		}
		public TerminalNode EOF() { return getToken(NuggetParser.EOF, 0); }
		public ExecuteContext execute() {
			return getRuleContext(ExecuteContext.class,0);
		}
		public InitextractContext initextract() {
			return getRuleContext(InitextractContext.class,0);
		}
		public PrintIdContext printId() {
			return getRuleContext(PrintIdContext.class,0);
		}
		public AssignContext assign() {
			return getRuleContext(AssignContext.class,0);
		}
		public SaveContext save() {
			return getRuleContext(SaveContext.class,0);
		}
		public NuggetContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nugget; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterNugget(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitNugget(this);
		}
	}

	public final NuggetContext nugget() throws RecognitionException {
		NuggetContext _localctx = new NuggetContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_nugget);
		try {
			setState(55);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(36);
				sin();
				setState(37);
				match(EOF);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(39);
				execute();
				setState(40);
				match(EOF);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(42);
				initextract(0);
				setState(43);
				match(EOF);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(45);
				printId();
				setState(46);
				match(EOF);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(48);
				assign();
				setState(49);
				match(EOF);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(51);
				save();
				setState(52);
				match(EOF);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(54);
				match(EOF);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InitextractContext extends ParserRuleContext {
		public TerminalNode StrLit() { return getToken(NuggetParser.StrLit, 0); }
		public TargetContext target() {
			return getRuleContext(TargetContext.class,0);
		}
		public TerminalNode EXTRACT() { return getToken(NuggetParser.EXTRACT, 0); }
		public SubtypeContext subtype() {
			return getRuleContext(SubtypeContext.class,0);
		}
		public InitextractContext initextract() {
			return getRuleContext(InitextractContext.class,0);
		}
		public StreamTaskContext streamTask() {
			return getRuleContext(StreamTaskContext.class,0);
		}
		public InitextractContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_initextract; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterInitextract(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitInitextract(this);
		}
	}

	public final InitextractContext initextract() throws RecognitionException {
		return initextract(0);
	}

	private InitextractContext initextract(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		InitextractContext _localctx = new InitextractContext(_ctx, _parentState);
		InitextractContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_initextract, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(58);
			match(StrLit);
			setState(59);
			match(T__0);
			setState(60);
			target();
			setState(61);
			match(EXTRACT);
			setState(62);
			subtype();
			}
			_ctx.stop = _input.LT(-1);
			setState(68);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new InitextractContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_initextract);
					setState(64);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(65);
					streamTask();
					}
					} 
				}
				setState(70);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class AssignContext extends ParserRuleContext {
		public List<TerminalNode> StrLit() { return getTokens(NuggetParser.StrLit); }
		public TerminalNode StrLit(int i) {
			return getToken(NuggetParser.StrLit, i);
		}
		public AssignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterAssign(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitAssign(this);
		}
	}

	public final AssignContext assign() throws RecognitionException {
		AssignContext _localctx = new AssignContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_assign);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			match(StrLit);
			setState(72);
			match(T__0);
			setState(73);
			match(StrLit);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExecuteContext extends ParserRuleContext {
		public List<TerminalNode> StrLit() { return getTokens(NuggetParser.StrLit); }
		public TerminalNode StrLit(int i) {
			return getToken(NuggetParser.StrLit, i);
		}
		public TaskContext task() {
			return getRuleContext(TaskContext.class,0);
		}
		public List<StreamTaskContext> streamTask() {
			return getRuleContexts(StreamTaskContext.class);
		}
		public StreamTaskContext streamTask(int i) {
			return getRuleContext(StreamTaskContext.class,i);
		}
		public List<FilterContext> filter() {
			return getRuleContexts(FilterContext.class);
		}
		public FilterContext filter(int i) {
			return getRuleContext(FilterContext.class,i);
		}
		public ExecuteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_execute; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterExecute(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitExecute(this);
		}
	}

	public final ExecuteContext execute() throws RecognitionException {
		ExecuteContext _localctx = new ExecuteContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_execute);
		int _la;
		try {
			setState(117);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(75);
				match(StrLit);
				setState(76);
				match(T__0);
				setState(77);
				match(StrLit);
				setState(78);
				match(T__1);
				setState(79);
				task();
				setState(83);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__1) {
					{
					{
					setState(80);
					streamTask();
					}
					}
					setState(85);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(86);
				match(StrLit);
				setState(87);
				match(T__0);
				setState(88);
				match(StrLit);
				setState(89);
				match(T__1);
				setState(90);
				task();
				setState(91);
				filter();
				setState(96);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__2) {
					{
					{
					setState(92);
					match(T__2);
					setState(93);
					filter();
					}
					}
					setState(98);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(102);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__1) {
					{
					{
					setState(99);
					streamTask();
					}
					}
					setState(104);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(105);
				match(StrLit);
				setState(106);
				match(T__0);
				setState(107);
				match(StrLit);
				setState(108);
				match(T__1);
				setState(109);
				task();
				setState(110);
				match(StrLit);
				setState(114);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__1) {
					{
					{
					setState(111);
					streamTask();
					}
					}
					setState(116);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StreamTaskContext extends ParserRuleContext {
		public TaskContext task() {
			return getRuleContext(TaskContext.class,0);
		}
		public List<FilterContext> filter() {
			return getRuleContexts(FilterContext.class);
		}
		public FilterContext filter(int i) {
			return getRuleContext(FilterContext.class,i);
		}
		public TerminalNode StrLit() { return getToken(NuggetParser.StrLit, 0); }
		public StreamTaskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_streamTask; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterStreamTask(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitStreamTask(this);
		}
	}

	public final StreamTaskContext streamTask() throws RecognitionException {
		StreamTaskContext _localctx = new StreamTaskContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_streamTask);
		try {
			int _alt;
			setState(135);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(119);
				match(T__1);
				setState(120);
				task();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(121);
				match(T__1);
				setState(122);
				task();
				setState(123);
				filter();
				setState(128);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(124);
						match(T__2);
						setState(125);
						filter();
						}
						} 
					}
					setState(130);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
				}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(131);
				match(T__1);
				setState(132);
				task();
				setState(133);
				match(StrLit);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SaveContext extends ParserRuleContext {
		public List<TerminalNode> StrLit() { return getTokens(NuggetParser.StrLit); }
		public TerminalNode StrLit(int i) {
			return getToken(NuggetParser.StrLit, i);
		}
		public SaveContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_save; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterSave(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitSave(this);
		}
	}

	public final SaveContext save() throws RecognitionException {
		SaveContext _localctx = new SaveContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_save);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(137);
			match(T__3);
			setState(138);
			match(StrLit);
			setState(139);
			match(T__4);
			setState(140);
			match(StrLit);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FilterContext extends ParserRuleContext {
		public FilenameContext filename() {
			return getRuleContext(FilenameContext.class,0);
		}
		public TimefilterContext timefilter() {
			return getRuleContext(TimefilterContext.class,0);
		}
		public FilterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_filter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterFilter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitFilter(this);
		}
	}

	public final FilterContext filter() throws RecognitionException {
		FilterContext _localctx = new FilterContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_filter);
		try {
			setState(144);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__5:
				enterOuterAlt(_localctx, 1);
				{
				setState(142);
				filename();
				}
				break;
			case CTIME:
			case MTIME:
				enterOuterAlt(_localctx, 2);
				{
				setState(143);
				timefilter();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FilenameContext extends ParserRuleContext {
		public TerminalNode StrLit() { return getToken(NuggetParser.StrLit, 0); }
		public FilenameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_filename; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterFilename(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitFilename(this);
		}
	}

	public final FilenameContext filename() throws RecognitionException {
		FilenameContext _localctx = new FilenameContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_filename);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(146);
			match(T__5);
			setState(147);
			match(StrLit);
			setState(148);
			match(T__5);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TimefilterContext extends ParserRuleContext {
		public ReferenceContext reference() {
			return getRuleContext(ReferenceContext.class,0);
		}
		public TerminalNode OPERATION() { return getToken(NuggetParser.OPERATION, 0); }
		public DateContext date() {
			return getRuleContext(DateContext.class,0);
		}
		public TimefilterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_timefilter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterTimefilter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitTimefilter(this);
		}
	}

	public final TimefilterContext timefilter() throws RecognitionException {
		TimefilterContext _localctx = new TimefilterContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_timefilter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(150);
			reference();
			setState(151);
			match(OPERATION);
			setState(152);
			date();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ReferenceContext extends ParserRuleContext {
		public TerminalNode CTIME() { return getToken(NuggetParser.CTIME, 0); }
		public TerminalNode MTIME() { return getToken(NuggetParser.MTIME, 0); }
		public ReferenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_reference; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterReference(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitReference(this);
		}
	}

	public final ReferenceContext reference() throws RecognitionException {
		ReferenceContext _localctx = new ReferenceContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_reference);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(154);
			_la = _input.LA(1);
			if ( !(_la==CTIME || _la==MTIME) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DateContext extends ParserRuleContext {
		public TerminalNode DATE() { return getToken(NuggetParser.DATE, 0); }
		public DateContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_date; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterDate(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitDate(this);
		}
	}

	public final DateContext date() throws RecognitionException {
		DateContext _localctx = new DateContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_date);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(156);
			match(T__5);
			setState(157);
			match(DATE);
			setState(158);
			match(T__5);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SubtypeContext extends ParserRuleContext {
		public TerminalNode FILES() { return getToken(NuggetParser.FILES, 0); }
		public TerminalNode IMAGES() { return getToken(NuggetParser.IMAGES, 0); }
		public TerminalNode DOCUMENTS() { return getToken(NuggetParser.DOCUMENTS, 0); }
		public TerminalNode ALL() { return getToken(NuggetParser.ALL, 0); }
		public SubtypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subtype; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterSubtype(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitSubtype(this);
		}
	}

	public final SubtypeContext subtype() throws RecognitionException {
		SubtypeContext _localctx = new SubtypeContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_subtype);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(160);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ALL) | (1L << FILES) | (1L << IMAGES) | (1L << DOCUMENTS))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TaskContext extends ParserRuleContext {
		public TerminalNode HASH() { return getToken(NuggetParser.HASH, 0); }
		public TerminalNode EXTRACT() { return getToken(NuggetParser.EXTRACT, 0); }
		public TerminalNode SELECT() { return getToken(NuggetParser.SELECT, 0); }
		public TerminalNode JOIN() { return getToken(NuggetParser.JOIN, 0); }
		public TaskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_task; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterTask(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitTask(this);
		}
	}

	public final TaskContext task() throws RecognitionException {
		TaskContext _localctx = new TaskContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_task);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << HASH) | (1L << SELECT) | (1L << EXTRACT) | (1L << JOIN))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TargetContext extends ParserRuleContext {
		public TerminalNode StrLit() { return getToken(NuggetParser.StrLit, 0); }
		public TargetContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_target; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterTarget(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitTarget(this);
		}
	}

	public final TargetContext target() throws RecognitionException {
		TargetContext _localctx = new TargetContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_target);
		try {
			setState(168);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__5:
				enterOuterAlt(_localctx, 1);
				{
				setState(164);
				match(T__5);
				setState(165);
				match(StrLit);
				setState(166);
				match(T__5);
				}
				break;
			case StrLit:
				enterOuterAlt(_localctx, 2);
				{
				setState(167);
				match(StrLit);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FieldContext extends ParserRuleContext {
		public List<TerminalNode> StrLit() { return getTokens(NuggetParser.StrLit); }
		public TerminalNode StrLit(int i) {
			return getToken(NuggetParser.StrLit, i);
		}
		public FieldContext field() {
			return getRuleContext(FieldContext.class,0);
		}
		public FieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_field; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterField(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitField(this);
		}
	}

	public final FieldContext field() throws RecognitionException {
		FieldContext _localctx = new FieldContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_field);
		int _la;
		try {
			setState(182);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case StrLit:
				enterOuterAlt(_localctx, 1);
				{
				setState(170);
				match(StrLit);
				setState(175);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__2) {
					{
					{
					setState(171);
					match(T__2);
					setState(172);
					match(StrLit);
					}
					}
					setState(177);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 2);
				{
				setState(178);
				match(T__6);
				setState(179);
				field();
				setState(180);
				match(T__6);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SourceidentifierContext extends ParserRuleContext {
		public TerminalNode StrLit() { return getToken(NuggetParser.StrLit, 0); }
		public SourceidentifierContext sourceidentifier() {
			return getRuleContext(SourceidentifierContext.class,0);
		}
		public SourceidentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sourceidentifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterSourceidentifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitSourceidentifier(this);
		}
	}

	public final SourceidentifierContext sourceidentifier() throws RecognitionException {
		SourceidentifierContext _localctx = new SourceidentifierContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_sourceidentifier);
		try {
			setState(189);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case StrLit:
				enterOuterAlt(_localctx, 1);
				{
				setState(184);
				match(StrLit);
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 2);
				{
				setState(185);
				match(T__6);
				setState(186);
				sourceidentifier();
				setState(187);
				match(T__6);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrintIdContext extends ParserRuleContext {
		public TerminalNode StrLit() { return getToken(NuggetParser.StrLit, 0); }
		public PrintIdContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printId; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterPrintId(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitPrintId(this);
		}
	}

	public final PrintIdContext printId() throws RecognitionException {
		PrintIdContext _localctx = new PrintIdContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_printId);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(191);
			match(StrLit);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SinContext extends ParserRuleContext {
		public TerminalNode SIN() { return getToken(NuggetParser.SIN, 0); }
		public TerminalNode NUMBER() { return getToken(NuggetParser.NUMBER, 0); }
		public SinContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sin; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).enterSin(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof NuggetListener ) ((NuggetListener)listener).exitSin(this);
		}
	}

	public final SinContext sin() throws RecognitionException {
		SinContext _localctx = new SinContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_sin);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(193);
			match(SIN);
			setState(194);
			match(T__7);
			setState(195);
			match(NUMBER);
			setState(196);
			match(T__8);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return initextract_sempred((InitextractContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean initextract_sempred(InitextractContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3 \u00c9\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\3\2\5\2:\n\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\7\3E\n\3"+
		"\f\3\16\3H\13\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\7\5T\n\5\f\5\16"+
		"\5W\13\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\7\5a\n\5\f\5\16\5d\13\5\3\5\7"+
		"\5g\n\5\f\5\16\5j\13\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\7\5s\n\5\f\5\16\5v"+
		"\13\5\5\5x\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\7\6\u0081\n\6\f\6\16\6\u0084"+
		"\13\6\3\6\3\6\3\6\3\6\5\6\u008a\n\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\5\b\u0093"+
		"\n\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\13\3\13\3\f\3\f\3\f\3\f\3\r\3\r"+
		"\3\16\3\16\3\17\3\17\3\17\3\17\5\17\u00ab\n\17\3\20\3\20\3\20\7\20\u00b0"+
		"\n\20\f\20\16\20\u00b3\13\20\3\20\3\20\3\20\3\20\5\20\u00b9\n\20\3\21"+
		"\3\21\3\21\3\21\3\21\5\21\u00c0\n\21\3\22\3\22\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\2\3\4\24\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$\2\5\3\2\f\r"+
		"\4\2\21\21\24\26\4\2\22\23\27\30\2\u00cb\29\3\2\2\2\4;\3\2\2\2\6I\3\2"+
		"\2\2\bw\3\2\2\2\n\u0089\3\2\2\2\f\u008b\3\2\2\2\16\u0092\3\2\2\2\20\u0094"+
		"\3\2\2\2\22\u0098\3\2\2\2\24\u009c\3\2\2\2\26\u009e\3\2\2\2\30\u00a2\3"+
		"\2\2\2\32\u00a4\3\2\2\2\34\u00aa\3\2\2\2\36\u00b8\3\2\2\2 \u00bf\3\2\2"+
		"\2\"\u00c1\3\2\2\2$\u00c3\3\2\2\2&\'\5$\23\2\'(\7\2\2\3(:\3\2\2\2)*\5"+
		"\b\5\2*+\7\2\2\3+:\3\2\2\2,-\5\4\3\2-.\7\2\2\3.:\3\2\2\2/\60\5\"\22\2"+
		"\60\61\7\2\2\3\61:\3\2\2\2\62\63\5\6\4\2\63\64\7\2\2\3\64:\3\2\2\2\65"+
		"\66\5\f\7\2\66\67\7\2\2\3\67:\3\2\2\28:\7\2\2\39&\3\2\2\29)\3\2\2\29,"+
		"\3\2\2\29/\3\2\2\29\62\3\2\2\29\65\3\2\2\298\3\2\2\2:\3\3\2\2\2;<\b\3"+
		"\1\2<=\7\34\2\2=>\7\3\2\2>?\5\34\17\2?@\7\27\2\2@A\5\30\r\2AF\3\2\2\2"+
		"BC\f\3\2\2CE\5\n\6\2DB\3\2\2\2EH\3\2\2\2FD\3\2\2\2FG\3\2\2\2G\5\3\2\2"+
		"\2HF\3\2\2\2IJ\7\34\2\2JK\7\3\2\2KL\7\34\2\2L\7\3\2\2\2MN\7\34\2\2NO\7"+
		"\3\2\2OP\7\34\2\2PQ\7\4\2\2QU\5\32\16\2RT\5\n\6\2SR\3\2\2\2TW\3\2\2\2"+
		"US\3\2\2\2UV\3\2\2\2Vx\3\2\2\2WU\3\2\2\2XY\7\34\2\2YZ\7\3\2\2Z[\7\34\2"+
		"\2[\\\7\4\2\2\\]\5\32\16\2]b\5\16\b\2^_\7\5\2\2_a\5\16\b\2`^\3\2\2\2a"+
		"d\3\2\2\2b`\3\2\2\2bc\3\2\2\2ch\3\2\2\2db\3\2\2\2eg\5\n\6\2fe\3\2\2\2"+
		"gj\3\2\2\2hf\3\2\2\2hi\3\2\2\2ix\3\2\2\2jh\3\2\2\2kl\7\34\2\2lm\7\3\2"+
		"\2mn\7\34\2\2no\7\4\2\2op\5\32\16\2pt\7\34\2\2qs\5\n\6\2rq\3\2\2\2sv\3"+
		"\2\2\2tr\3\2\2\2tu\3\2\2\2ux\3\2\2\2vt\3\2\2\2wM\3\2\2\2wX\3\2\2\2wk\3"+
		"\2\2\2x\t\3\2\2\2yz\7\4\2\2z\u008a\5\32\16\2{|\7\4\2\2|}\5\32\16\2}\u0082"+
		"\5\16\b\2~\177\7\5\2\2\177\u0081\5\16\b\2\u0080~\3\2\2\2\u0081\u0084\3"+
		"\2\2\2\u0082\u0080\3\2\2\2\u0082\u0083\3\2\2\2\u0083\u008a\3\2\2\2\u0084"+
		"\u0082\3\2\2\2\u0085\u0086\7\4\2\2\u0086\u0087\5\32\16\2\u0087\u0088\7"+
		"\34\2\2\u0088\u008a\3\2\2\2\u0089y\3\2\2\2\u0089{\3\2\2\2\u0089\u0085"+
		"\3\2\2\2\u008a\13\3\2\2\2\u008b\u008c\7\6\2\2\u008c\u008d\7\34\2\2\u008d"+
		"\u008e\7\7\2\2\u008e\u008f\7\34\2\2\u008f\r\3\2\2\2\u0090\u0093\5\20\t"+
		"\2\u0091\u0093\5\22\n\2\u0092\u0090\3\2\2\2\u0092\u0091\3\2\2\2\u0093"+
		"\17\3\2\2\2\u0094\u0095\7\b\2\2\u0095\u0096\7\34\2\2\u0096\u0097\7\b\2"+
		"\2\u0097\21\3\2\2\2\u0098\u0099\5\24\13\2\u0099\u009a\7\35\2\2\u009a\u009b"+
		"\5\26\f\2\u009b\23\3\2\2\2\u009c\u009d\t\2\2\2\u009d\25\3\2\2\2\u009e"+
		"\u009f\7\b\2\2\u009f\u00a0\7\36\2\2\u00a0\u00a1\7\b\2\2\u00a1\27\3\2\2"+
		"\2\u00a2\u00a3\t\3\2\2\u00a3\31\3\2\2\2\u00a4\u00a5\t\4\2\2\u00a5\33\3"+
		"\2\2\2\u00a6\u00a7\7\b\2\2\u00a7\u00a8\7\34\2\2\u00a8\u00ab\7\b\2\2\u00a9"+
		"\u00ab\7\34\2\2\u00aa\u00a6\3\2\2\2\u00aa\u00a9\3\2\2\2\u00ab\35\3\2\2"+
		"\2\u00ac\u00b1\7\34\2\2\u00ad\u00ae\7\5\2\2\u00ae\u00b0\7\34\2\2\u00af"+
		"\u00ad\3\2\2\2\u00b0\u00b3\3\2\2\2\u00b1\u00af\3\2\2\2\u00b1\u00b2\3\2"+
		"\2\2\u00b2\u00b9\3\2\2\2\u00b3\u00b1\3\2\2\2\u00b4\u00b5\7\t\2\2\u00b5"+
		"\u00b6\5\36\20\2\u00b6\u00b7\7\t\2\2\u00b7\u00b9\3\2\2\2\u00b8\u00ac\3"+
		"\2\2\2\u00b8\u00b4\3\2\2\2\u00b9\37\3\2\2\2\u00ba\u00c0\7\34\2\2\u00bb"+
		"\u00bc\7\t\2\2\u00bc\u00bd\5 \21\2\u00bd\u00be\7\t\2\2\u00be\u00c0\3\2"+
		"\2\2\u00bf\u00ba\3\2\2\2\u00bf\u00bb\3\2\2\2\u00c0!\3\2\2\2\u00c1\u00c2"+
		"\7\34\2\2\u00c2#\3\2\2\2\u00c3\u00c4\7\16\2\2\u00c4\u00c5\7\n\2\2\u00c5"+
		"\u00c6\7\33\2\2\u00c6\u00c7\7\13\2\2\u00c7%\3\2\2\2\209FUbhtw\u0082\u0089"+
		"\u0092\u00aa\u00b1\u00b8\u00bf";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}