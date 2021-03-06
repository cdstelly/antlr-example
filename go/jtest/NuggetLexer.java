// Generated from Nugget.g4 by ANTLR 4.7
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class NuggetLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		CTIME=10, MTIME=11, SIN=12, LOAD=13, FROM=14, ALL=15, HASH=16, SELECT=17, 
		FILES=18, IMAGES=19, DOCUMENTS=20, EXTRACT=21, JOIN=22, WS=23, COMMENT=24, 
		NUMBER=25, StrLit=26, OPERATION=27, DATE=28;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
		"CTIME", "MTIME", "SIN", "LOAD", "FROM", "ALL", "HASH", "SELECT", "FILES", 
		"IMAGES", "DOCUMENTS", "EXTRACT", "JOIN", "WS", "COMMENT", "NUMBER", "StrLit", 
		"OPERATION", "DATE", "DIGIT", "DLMT", "WILDCARD"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'='", "'|'", "','", "'save'", "'to'", "'\"'", "'''", "'('", "')'", 
		null, null, "'sin'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, null, null, "CTIME", "MTIME", 
		"SIN", "LOAD", "FROM", "ALL", "HASH", "SELECT", "FILES", "IMAGES", "DOCUMENTS", 
		"EXTRACT", "JOIN", "WS", "COMMENT", "NUMBER", "StrLit", "OPERATION", "DATE"
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


	public NuggetLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Nugget.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\36\u0125\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\b"+
		"\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\5\13b\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\fn\n\f\3\r\3\r\3"+
		"\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16|\n\16\3\17\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\5\17\u0086\n\17\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\5\20\u008e\n\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\5\21\u0098"+
		"\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22"+
		"\u00a6\n\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u00b2"+
		"\n\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24"+
		"\u00c0\n\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u00d4\n\25\3\26\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u00e4\n\26\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\5\27\u00ee\n\27\3\30\6\30\u00f1\n\30\r"+
		"\30\16\30\u00f2\3\30\3\30\3\31\3\31\3\31\3\31\7\31\u00fb\n\31\f\31\16"+
		"\31\u00fe\13\31\3\31\3\31\3\32\6\32\u0103\n\32\r\32\16\32\u0104\3\33\3"+
		"\33\6\33\u0109\n\33\r\33\16\33\u010a\3\34\3\34\3\34\3\34\3\34\3\34\5\34"+
		"\u0113\n\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36"+
		"\3\36\3\37\3\37\3 \3 \3\u00fc\2!\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23"+
		"\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31"+
		"\61\32\63\33\65\34\67\359\36;\2=\2?\2\3\2\5\5\2\13\f\17\17\"\"\7\2,,\60"+
		"\60<<C\\c|\4\2>>@@\2\u0135\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2"+
		"\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2"+
		"\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3"+
		"\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2"+
		"\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67"+
		"\3\2\2\2\29\3\2\2\2\3A\3\2\2\2\5C\3\2\2\2\7E\3\2\2\2\tG\3\2\2\2\13L\3"+
		"\2\2\2\rO\3\2\2\2\17Q\3\2\2\2\21S\3\2\2\2\23U\3\2\2\2\25a\3\2\2\2\27m"+
		"\3\2\2\2\31o\3\2\2\2\33{\3\2\2\2\35\u0085\3\2\2\2\37\u008d\3\2\2\2!\u0097"+
		"\3\2\2\2#\u00a5\3\2\2\2%\u00b1\3\2\2\2\'\u00bf\3\2\2\2)\u00d3\3\2\2\2"+
		"+\u00e3\3\2\2\2-\u00ed\3\2\2\2/\u00f0\3\2\2\2\61\u00f6\3\2\2\2\63\u0102"+
		"\3\2\2\2\65\u0108\3\2\2\2\67\u0112\3\2\2\29\u0114\3\2\2\2;\u011f\3\2\2"+
		"\2=\u0121\3\2\2\2?\u0123\3\2\2\2AB\7?\2\2B\4\3\2\2\2CD\7~\2\2D\6\3\2\2"+
		"\2EF\7.\2\2F\b\3\2\2\2GH\7u\2\2HI\7c\2\2IJ\7x\2\2JK\7g\2\2K\n\3\2\2\2"+
		"LM\7v\2\2MN\7q\2\2N\f\3\2\2\2OP\7$\2\2P\16\3\2\2\2QR\7)\2\2R\20\3\2\2"+
		"\2ST\7*\2\2T\22\3\2\2\2UV\7+\2\2V\24\3\2\2\2WX\7e\2\2XY\7v\2\2YZ\7k\2"+
		"\2Z[\7o\2\2[b\7g\2\2\\]\7E\2\2]^\7V\2\2^_\7K\2\2_`\7O\2\2`b\7G\2\2aW\3"+
		"\2\2\2a\\\3\2\2\2b\26\3\2\2\2cd\7o\2\2de\7v\2\2ef\7k\2\2fg\7o\2\2gn\7"+
		"g\2\2hi\7O\2\2ij\7V\2\2jk\7K\2\2kl\7O\2\2ln\7G\2\2mc\3\2\2\2mh\3\2\2\2"+
		"n\30\3\2\2\2op\7u\2\2pq\7k\2\2qr\7p\2\2r\32\3\2\2\2st\7n\2\2tu\7q\2\2"+
		"uv\7c\2\2v|\7f\2\2wx\7N\2\2xy\7Q\2\2yz\7C\2\2z|\7F\2\2{s\3\2\2\2{w\3\2"+
		"\2\2|\34\3\2\2\2}~\7h\2\2~\177\7t\2\2\177\u0080\7q\2\2\u0080\u0086\7o"+
		"\2\2\u0081\u0082\7H\2\2\u0082\u0083\7T\2\2\u0083\u0084\7Q\2\2\u0084\u0086"+
		"\7O\2\2\u0085}\3\2\2\2\u0085\u0081\3\2\2\2\u0086\36\3\2\2\2\u0087\u0088"+
		"\7c\2\2\u0088\u0089\7n\2\2\u0089\u008e\7n\2\2\u008a\u008b\7C\2\2\u008b"+
		"\u008c\7N\2\2\u008c\u008e\7N\2\2\u008d\u0087\3\2\2\2\u008d\u008a\3\2\2"+
		"\2\u008e \3\2\2\2\u008f\u0090\7j\2\2\u0090\u0091\7c\2\2\u0091\u0092\7"+
		"u\2\2\u0092\u0098\7j\2\2\u0093\u0094\7J\2\2\u0094\u0095\7C\2\2\u0095\u0096"+
		"\7U\2\2\u0096\u0098\7J\2\2\u0097\u008f\3\2\2\2\u0097\u0093\3\2\2\2\u0098"+
		"\"\3\2\2\2\u0099\u009a\7u\2\2\u009a\u009b\7g\2\2\u009b\u009c\7n\2\2\u009c"+
		"\u009d\7g\2\2\u009d\u009e\7e\2\2\u009e\u00a6\7v\2\2\u009f\u00a0\7U\2\2"+
		"\u00a0\u00a1\7G\2\2\u00a1\u00a2\7N\2\2\u00a2\u00a3\7G\2\2\u00a3\u00a4"+
		"\7E\2\2\u00a4\u00a6\7V\2\2\u00a5\u0099\3\2\2\2\u00a5\u009f\3\2\2\2\u00a6"+
		"$\3\2\2\2\u00a7\u00a8\7h\2\2\u00a8\u00a9\7k\2\2\u00a9\u00aa\7n\2\2\u00aa"+
		"\u00ab\7g\2\2\u00ab\u00b2\7u\2\2\u00ac\u00ad\7H\2\2\u00ad\u00ae\7K\2\2"+
		"\u00ae\u00af\7N\2\2\u00af\u00b0\7G\2\2\u00b0\u00b2\7U\2\2\u00b1\u00a7"+
		"\3\2\2\2\u00b1\u00ac\3\2\2\2\u00b2&\3\2\2\2\u00b3\u00b4\7k\2\2\u00b4\u00b5"+
		"\7o\2\2\u00b5\u00b6\7c\2\2\u00b6\u00b7\7i\2\2\u00b7\u00b8\7g\2\2\u00b8"+
		"\u00c0\7u\2\2\u00b9\u00ba\7K\2\2\u00ba\u00bb\7O\2\2\u00bb\u00bc\7C\2\2"+
		"\u00bc\u00bd\7I\2\2\u00bd\u00be\7G\2\2\u00be\u00c0\7U\2\2\u00bf\u00b3"+
		"\3\2\2\2\u00bf\u00b9\3\2\2\2\u00c0(\3\2\2\2\u00c1\u00c2\7f\2\2\u00c2\u00c3"+
		"\7q\2\2\u00c3\u00c4\7e\2\2\u00c4\u00c5\7w\2\2\u00c5\u00c6\7o\2\2\u00c6"+
		"\u00c7\7g\2\2\u00c7\u00c8\7p\2\2\u00c8\u00c9\7v\2\2\u00c9\u00d4\7u\2\2"+
		"\u00ca\u00cb\7F\2\2\u00cb\u00cc\7Q\2\2\u00cc\u00cd\7E\2\2\u00cd\u00ce"+
		"\7W\2\2\u00ce\u00cf\7O\2\2\u00cf\u00d0\7G\2\2\u00d0\u00d1\7P\2\2\u00d1"+
		"\u00d2\7V\2\2\u00d2\u00d4\7U\2\2\u00d3\u00c1\3\2\2\2\u00d3\u00ca\3\2\2"+
		"\2\u00d4*\3\2\2\2\u00d5\u00d6\7g\2\2\u00d6\u00d7\7z\2\2\u00d7\u00d8\7"+
		"v\2\2\u00d8\u00d9\7t\2\2\u00d9\u00da\7c\2\2\u00da\u00db\7e\2\2\u00db\u00e4"+
		"\7v\2\2\u00dc\u00dd\7G\2\2\u00dd\u00de\7Z\2\2\u00de\u00df\7V\2\2\u00df"+
		"\u00e0\7T\2\2\u00e0\u00e1\7C\2\2\u00e1\u00e2\7E\2\2\u00e2\u00e4\7V\2\2"+
		"\u00e3\u00d5\3\2\2\2\u00e3\u00dc\3\2\2\2\u00e4,\3\2\2\2\u00e5\u00e6\7"+
		"l\2\2\u00e6\u00e7\7q\2\2\u00e7\u00e8\7k\2\2\u00e8\u00ee\7p\2\2\u00e9\u00ea"+
		"\7L\2\2\u00ea\u00eb\7Q\2\2\u00eb\u00ec\7K\2\2\u00ec\u00ee\7P\2\2\u00ed"+
		"\u00e5\3\2\2\2\u00ed\u00e9\3\2\2\2\u00ee.\3\2\2\2\u00ef\u00f1\t\2\2\2"+
		"\u00f0\u00ef\3\2\2\2\u00f1\u00f2\3\2\2\2\u00f2\u00f0\3\2\2\2\u00f2\u00f3"+
		"\3\2\2\2\u00f3\u00f4\3\2\2\2\u00f4\u00f5\b\30\2\2\u00f5\60\3\2\2\2\u00f6"+
		"\u00f7\7\61\2\2\u00f7\u00f8\7\61\2\2\u00f8\u00fc\3\2\2\2\u00f9\u00fb\13"+
		"\2\2\2\u00fa\u00f9\3\2\2\2\u00fb\u00fe\3\2\2\2\u00fc\u00fd\3\2\2\2\u00fc"+
		"\u00fa\3\2\2\2\u00fd\u00ff\3\2\2\2\u00fe\u00fc\3\2\2\2\u00ff\u0100\b\31"+
		"\2\2\u0100\62\3\2\2\2\u0101\u0103\5;\36\2\u0102\u0101\3\2\2\2\u0103\u0104"+
		"\3\2\2\2\u0104\u0102\3\2\2\2\u0104\u0105\3\2\2\2\u0105\64\3\2\2\2\u0106"+
		"\u0109\t\3\2\2\u0107\u0109\5;\36\2\u0108\u0106\3\2\2\2\u0108\u0107\3\2"+
		"\2\2\u0109\u010a\3\2\2\2\u010a\u0108\3\2\2\2\u010a\u010b\3\2\2\2\u010b"+
		"\66\3\2\2\2\u010c\u0113\t\4\2\2\u010d\u010e\7>\2\2\u010e\u0113\7?\2\2"+
		"\u010f\u0110\7@\2\2\u0110\u0113\7?\2\2\u0111\u0113\7?\2\2\u0112\u010c"+
		"\3\2\2\2\u0112\u010d\3\2\2\2\u0112\u010f\3\2\2\2\u0112\u0111\3\2\2\2\u0113"+
		"8\3\2\2\2\u0114\u0115\5;\36\2\u0115\u0116\5;\36\2\u0116\u0117\7\61\2\2"+
		"\u0117\u0118\5;\36\2\u0118\u0119\5;\36\2\u0119\u011a\7\61\2\2\u011a\u011b"+
		"\5;\36\2\u011b\u011c\5;\36\2\u011c\u011d\5;\36\2\u011d\u011e\5;\36\2\u011e"+
		":\3\2\2\2\u011f\u0120\4\62;\2\u0120<\3\2\2\2\u0121\u0122\7.\2\2\u0122"+
		">\3\2\2\2\u0123\u0124\7,\2\2\u0124@\3\2\2\2\25\2am{\u0085\u008d\u0097"+
		"\u00a5\u00b1\u00bf\u00d3\u00e3\u00ed\u00f2\u00fc\u0104\u0108\u010a\u0112"+
		"\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}