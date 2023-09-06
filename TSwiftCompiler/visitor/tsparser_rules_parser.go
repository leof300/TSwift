// Code generated from TSParser_rules.g4 by ANTLR 4.13.1. DO NOT EDIT.

package TSVisitor // TSParser_rules
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TSParser_rulesParser struct {
	*antlr.BaseParser
}

var TSParser_rulesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tsparser_rulesParserInit() {
	staticData := &TSParser_rulesParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'-'", "'!'", "'%'", "'*'", "'/'", "'+'", "'=='", "'!='",
		"'>'", "'<'", "'>='", "'<='", "'('", "')'", "'&&'", "'||'", "'+='",
		"'-='", "':'", "'{'", "'}'", "','", "", "", "'var'", "'let'", "'nil'",
		"'String'", "'Int'", "'Float'", "'Bool'", "'Character'", "'if'", "'else'",
		"'switch'", "'default'", "'case'", "'while'", "'in'", "'for'", "'guard'",
		"'continue'", "'break'", "'return'", "'print'", "'struct'", "'self'",
		"'mutating'", "'func'", "'inout'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "NL", "WS", "VAR", "LET", "NIL", "STRING",
		"INT", "FLOAT", "BOOL", "CHARACTER", "IF", "ELSE", "SWITCH", "DEFAULT",
		"CASE", "WHILE", "IN", "FOR", "GUARD", "CONTINUE", "BREAK", "RETURN",
		"PRINT", "STRUCT", "SELF", "MUTATING", "FUNC", "INOUT", "VBOOL", "VSTRING",
		"VFLOAT", "VINTEGER", "ID", "SL_COMMENT", "ML_COMMENT",
	}
	staticData.RuleNames = []string{
		"start", "lsents", "sents", "expr", "declar", "block", "if", "switch",
		"default", "print",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 166, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 1, 5, 1, 25, 8, 1, 10, 1, 12, 1, 28, 9, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 39, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 55,
		8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 84, 8, 3, 10, 3, 12, 3, 87, 9, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 109, 8, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 131, 8, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7, 141, 8, 7, 11, 7, 12, 7, 142, 1, 7, 3,
		7, 146, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 5, 9, 159, 8, 9, 10, 9, 12, 9, 162, 9, 9, 1, 9, 1, 9, 1, 9, 0,
		1, 6, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 3, 1, 0, 5, 6, 2, 0, 2,
		2, 7, 7, 1, 0, 8, 13, 186, 0, 20, 1, 0, 0, 0, 2, 26, 1, 0, 0, 0, 4, 38,
		1, 0, 0, 0, 6, 54, 1, 0, 0, 0, 8, 108, 1, 0, 0, 0, 10, 110, 1, 0, 0, 0,
		12, 130, 1, 0, 0, 0, 14, 132, 1, 0, 0, 0, 16, 149, 1, 0, 0, 0, 18, 153,
		1, 0, 0, 0, 20, 21, 3, 2, 1, 0, 21, 22, 5, 0, 0, 1, 22, 1, 1, 0, 0, 0,
		23, 25, 3, 4, 2, 0, 24, 23, 1, 0, 0, 0, 25, 28, 1, 0, 0, 0, 26, 24, 1,
		0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 3, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 29,
		39, 3, 12, 6, 0, 30, 39, 3, 14, 7, 0, 31, 39, 3, 8, 4, 0, 32, 33, 3, 8,
		4, 0, 33, 34, 5, 1, 0, 0, 34, 35, 3, 6, 3, 0, 35, 39, 1, 0, 0, 0, 36, 39,
		3, 18, 9, 0, 37, 39, 3, 6, 3, 0, 38, 29, 1, 0, 0, 0, 38, 30, 1, 0, 0, 0,
		38, 31, 1, 0, 0, 0, 38, 32, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 37, 1,
		0, 0, 0, 39, 5, 1, 0, 0, 0, 40, 41, 6, 3, -1, 0, 41, 42, 5, 2, 0, 0, 42,
		55, 3, 6, 3, 17, 43, 44, 5, 3, 0, 0, 44, 55, 3, 6, 3, 16, 45, 46, 5, 14,
		0, 0, 46, 47, 3, 6, 3, 0, 47, 48, 5, 15, 0, 0, 48, 55, 1, 0, 0, 0, 49,
		55, 5, 53, 0, 0, 50, 55, 5, 55, 0, 0, 51, 55, 5, 54, 0, 0, 52, 55, 5, 52,
		0, 0, 53, 55, 5, 56, 0, 0, 54, 40, 1, 0, 0, 0, 54, 43, 1, 0, 0, 0, 54,
		45, 1, 0, 0, 0, 54, 49, 1, 0, 0, 0, 54, 50, 1, 0, 0, 0, 54, 51, 1, 0, 0,
		0, 54, 52, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 85, 1, 0, 0, 0, 56, 57,
		10, 15, 0, 0, 57, 58, 5, 4, 0, 0, 58, 84, 3, 6, 3, 16, 59, 60, 10, 14,
		0, 0, 60, 61, 7, 0, 0, 0, 61, 84, 3, 6, 3, 15, 62, 63, 10, 13, 0, 0, 63,
		64, 7, 1, 0, 0, 64, 84, 3, 6, 3, 14, 65, 66, 10, 12, 0, 0, 66, 67, 7, 2,
		0, 0, 67, 84, 3, 6, 3, 13, 68, 69, 10, 10, 0, 0, 69, 70, 5, 16, 0, 0, 70,
		84, 3, 6, 3, 11, 71, 72, 10, 9, 0, 0, 72, 73, 5, 17, 0, 0, 73, 84, 3, 6,
		3, 10, 74, 75, 10, 8, 0, 0, 75, 76, 5, 1, 0, 0, 76, 84, 3, 6, 3, 8, 77,
		78, 10, 7, 0, 0, 78, 79, 5, 18, 0, 0, 79, 84, 3, 6, 3, 7, 80, 81, 10, 6,
		0, 0, 81, 82, 5, 19, 0, 0, 82, 84, 3, 6, 3, 6, 83, 56, 1, 0, 0, 0, 83,
		59, 1, 0, 0, 0, 83, 62, 1, 0, 0, 0, 83, 65, 1, 0, 0, 0, 83, 68, 1, 0, 0,
		0, 83, 71, 1, 0, 0, 0, 83, 74, 1, 0, 0, 0, 83, 77, 1, 0, 0, 0, 83, 80,
		1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0,
		86, 7, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 89, 5, 26, 0, 0, 89, 90, 5,
		56, 0, 0, 90, 91, 5, 20, 0, 0, 91, 109, 5, 29, 0, 0, 92, 93, 5, 26, 0,
		0, 93, 94, 5, 56, 0, 0, 94, 95, 5, 20, 0, 0, 95, 109, 5, 30, 0, 0, 96,
		97, 5, 26, 0, 0, 97, 98, 5, 56, 0, 0, 98, 99, 5, 20, 0, 0, 99, 109, 5,
		31, 0, 0, 100, 101, 5, 26, 0, 0, 101, 102, 5, 56, 0, 0, 102, 103, 5, 20,
		0, 0, 103, 109, 5, 32, 0, 0, 104, 105, 5, 26, 0, 0, 105, 106, 5, 56, 0,
		0, 106, 107, 5, 20, 0, 0, 107, 109, 5, 33, 0, 0, 108, 88, 1, 0, 0, 0, 108,
		92, 1, 0, 0, 0, 108, 96, 1, 0, 0, 0, 108, 100, 1, 0, 0, 0, 108, 104, 1,
		0, 0, 0, 109, 9, 1, 0, 0, 0, 110, 111, 5, 21, 0, 0, 111, 112, 3, 2, 1,
		0, 112, 113, 5, 22, 0, 0, 113, 11, 1, 0, 0, 0, 114, 115, 5, 34, 0, 0, 115,
		116, 3, 6, 3, 0, 116, 117, 3, 10, 5, 0, 117, 131, 1, 0, 0, 0, 118, 119,
		5, 34, 0, 0, 119, 120, 3, 6, 3, 0, 120, 121, 3, 10, 5, 0, 121, 122, 5,
		35, 0, 0, 122, 123, 3, 10, 5, 0, 123, 131, 1, 0, 0, 0, 124, 125, 5, 34,
		0, 0, 125, 126, 3, 6, 3, 0, 126, 127, 3, 10, 5, 0, 127, 128, 5, 35, 0,
		0, 128, 129, 3, 12, 6, 0, 129, 131, 1, 0, 0, 0, 130, 114, 1, 0, 0, 0, 130,
		118, 1, 0, 0, 0, 130, 124, 1, 0, 0, 0, 131, 13, 1, 0, 0, 0, 132, 133, 5,
		36, 0, 0, 133, 134, 3, 6, 3, 0, 134, 140, 5, 21, 0, 0, 135, 136, 5, 38,
		0, 0, 136, 137, 3, 6, 3, 0, 137, 138, 5, 20, 0, 0, 138, 139, 3, 2, 1, 0,
		139, 141, 1, 0, 0, 0, 140, 135, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142,
		140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 1, 0, 0, 0, 144, 146,
		3, 16, 8, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 1, 0,
		0, 0, 147, 148, 5, 22, 0, 0, 148, 15, 1, 0, 0, 0, 149, 150, 5, 37, 0, 0,
		150, 151, 5, 20, 0, 0, 151, 152, 3, 2, 1, 0, 152, 17, 1, 0, 0, 0, 153,
		154, 5, 46, 0, 0, 154, 155, 5, 14, 0, 0, 155, 160, 3, 6, 3, 0, 156, 157,
		5, 23, 0, 0, 157, 159, 3, 6, 3, 0, 158, 156, 1, 0, 0, 0, 159, 162, 1, 0,
		0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 163, 1, 0, 0, 0,
		162, 160, 1, 0, 0, 0, 163, 164, 5, 15, 0, 0, 164, 19, 1, 0, 0, 0, 10, 26,
		38, 54, 83, 85, 108, 130, 142, 145, 160,
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

// TSParser_rulesParserInit initializes any static state used to implement TSParser_rulesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTSParser_rulesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TSParser_rulesParserInit() {
	staticData := &TSParser_rulesParserStaticData
	staticData.once.Do(tsparser_rulesParserInit)
}

// NewTSParser_rulesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTSParser_rulesParser(input antlr.TokenStream) *TSParser_rulesParser {
	TSParser_rulesParserInit()
	this := new(TSParser_rulesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TSParser_rulesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TSParser_rules.g4"

	return this
}

// TSParser_rulesParser tokens.
const (
	TSParser_rulesParserEOF        = antlr.TokenEOF
	TSParser_rulesParserT__0       = 1
	TSParser_rulesParserT__1       = 2
	TSParser_rulesParserT__2       = 3
	TSParser_rulesParserT__3       = 4
	TSParser_rulesParserT__4       = 5
	TSParser_rulesParserT__5       = 6
	TSParser_rulesParserT__6       = 7
	TSParser_rulesParserT__7       = 8
	TSParser_rulesParserT__8       = 9
	TSParser_rulesParserT__9       = 10
	TSParser_rulesParserT__10      = 11
	TSParser_rulesParserT__11      = 12
	TSParser_rulesParserT__12      = 13
	TSParser_rulesParserT__13      = 14
	TSParser_rulesParserT__14      = 15
	TSParser_rulesParserT__15      = 16
	TSParser_rulesParserT__16      = 17
	TSParser_rulesParserT__17      = 18
	TSParser_rulesParserT__18      = 19
	TSParser_rulesParserT__19      = 20
	TSParser_rulesParserT__20      = 21
	TSParser_rulesParserT__21      = 22
	TSParser_rulesParserT__22      = 23
	TSParser_rulesParserNL         = 24
	TSParser_rulesParserWS         = 25
	TSParser_rulesParserVAR        = 26
	TSParser_rulesParserLET        = 27
	TSParser_rulesParserNIL        = 28
	TSParser_rulesParserSTRING     = 29
	TSParser_rulesParserINT        = 30
	TSParser_rulesParserFLOAT      = 31
	TSParser_rulesParserBOOL       = 32
	TSParser_rulesParserCHARACTER  = 33
	TSParser_rulesParserIF         = 34
	TSParser_rulesParserELSE       = 35
	TSParser_rulesParserSWITCH     = 36
	TSParser_rulesParserDEFAULT    = 37
	TSParser_rulesParserCASE       = 38
	TSParser_rulesParserWHILE      = 39
	TSParser_rulesParserIN         = 40
	TSParser_rulesParserFOR        = 41
	TSParser_rulesParserGUARD      = 42
	TSParser_rulesParserCONTINUE   = 43
	TSParser_rulesParserBREAK      = 44
	TSParser_rulesParserRETURN     = 45
	TSParser_rulesParserPRINT      = 46
	TSParser_rulesParserSTRUCT     = 47
	TSParser_rulesParserSELF       = 48
	TSParser_rulesParserMUTATING   = 49
	TSParser_rulesParserFUNC       = 50
	TSParser_rulesParserINOUT      = 51
	TSParser_rulesParserVBOOL      = 52
	TSParser_rulesParserVSTRING    = 53
	TSParser_rulesParserVFLOAT     = 54
	TSParser_rulesParserVINTEGER   = 55
	TSParser_rulesParserID         = 56
	TSParser_rulesParserSL_COMMENT = 57
	TSParser_rulesParserML_COMMENT = 58
)

// TSParser_rulesParser rules.
const (
	TSParser_rulesParserRULE_start   = 0
	TSParser_rulesParserRULE_lsents  = 1
	TSParser_rulesParserRULE_sents   = 2
	TSParser_rulesParserRULE_expr    = 3
	TSParser_rulesParserRULE_declar  = 4
	TSParser_rulesParserRULE_block   = 5
	TSParser_rulesParserRULE_if      = 6
	TSParser_rulesParserRULE_switch  = 7
	TSParser_rulesParserRULE_default = 8
	TSParser_rulesParserRULE_print   = 9
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lsents() ILsentsContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Lsents() ILsentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILsentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILsentsContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TSParser_rulesParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Lsents()
	}
	{
		p.SetState(21)
		p.Match(TSParser_rulesParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILsentsContext is an interface to support dynamic dispatch.
type ILsentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSents() []ISentsContext
	Sents(i int) ISentsContext

	// IsLsentsContext differentiates from other interfaces.
	IsLsentsContext()
}

type LsentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLsentsContext() *LsentsContext {
	var p = new(LsentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_lsents
	return p
}

func InitEmptyLsentsContext(p *LsentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_lsents
}

func (*LsentsContext) IsLsentsContext() {}

func NewLsentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LsentsContext {
	var p = new(LsentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_lsents

	return p
}

func (s *LsentsContext) GetParser() antlr.Parser { return s.parser }

func (s *LsentsContext) AllSents() []ISentsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentsContext); ok {
			len++
		}
	}

	tst := make([]ISentsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentsContext); ok {
			tst[i] = t.(ISentsContext)
			i++
		}
	}

	return tst
}

func (s *LsentsContext) Sents(i int) ISentsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentsContext)
}

func (s *LsentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LsentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LsentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitLsents(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Lsents() (localctx ILsentsContext) {
	localctx = NewLsentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TSParser_rulesParserRULE_lsents)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&139682043159134220) != 0 {
		{
			p.SetState(23)
			p.Sents()
		}

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentsContext is an interface to support dynamic dispatch.
type ISentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentsContext differentiates from other interfaces.
	IsSentsContext()
}

type SentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentsContext() *SentsContext {
	var p = new(SentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_sents
	return p
}

func InitEmptySentsContext(p *SentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_sents
}

func (*SentsContext) IsSentsContext() {}

func NewSentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentsContext {
	var p = new(SentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_sents

	return p
}

func (s *SentsContext) GetParser() antlr.Parser { return s.parser }

func (s *SentsContext) CopyAll(ctx *SentsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SIfContext struct {
	SentsContext
}

func NewSIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SIfContext {
	var p = new(SIfContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SIfContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

func (s *SIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type SSwitchContext struct {
	SentsContext
}

func NewSSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SSwitchContext {
	var p = new(SSwitchContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SSwitchContext) Switch_() ISwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchContext)
}

func (s *SSwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

type SentExprContext struct {
	SentsContext
}

func NewSentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentExprContext {
	var p = new(SentExprContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SPrintContext struct {
	SentsContext
}

func NewSPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SPrintContext {
	var p = new(SPrintContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SPrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SPrintContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *SPrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDeclAsigContext struct {
	SentsContext
	op antlr.Token
}

func NewSDeclAsigContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDeclAsigContext {
	var p = new(SDeclAsigContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SDeclAsigContext) GetOp() antlr.Token { return s.op }

func (s *SDeclAsigContext) SetOp(v antlr.Token) { s.op = v }

func (s *SDeclAsigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDeclAsigContext) Declar() IDeclarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarContext)
}

func (s *SDeclAsigContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SDeclAsigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSDeclAsig(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDeclContext struct {
	SentsContext
}

func NewSDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDeclContext {
	var p = new(SDeclContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDeclContext) Declar() IDeclarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarContext)
}

func (s *SDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Sents() (localctx ISentsContext) {
	localctx = NewSentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TSParser_rulesParserRULE_sents)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.If_()
		}

	case 2:
		localctx = NewSSwitchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Switch_()
		}

	case 3:
		localctx = NewSDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(31)
			p.Declar()
		}

	case 4:
		localctx = NewSDeclAsigContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(32)
			p.Declar()
		}
		{
			p.SetState(33)

			var _m = p.Match(TSParser_rulesParserT__0)

			localctx.(*SDeclAsigContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.expr(0)
		}

	case 5:
		localctx = NewSPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(36)
			p.Print_()
		}

	case 6:
		localctx = NewSentExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(37)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EVStringContext struct {
	ExprContext
}

func NewEVStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EVStringContext {
	var p = new(EVStringContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EVStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EVStringContext) VSTRING() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVSTRING, 0)
}

func (s *EVStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEVString(s)

	default:
		return t.VisitChildren(s)
	}
}

type EIDContext struct {
	ExprContext
}

func NewEIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EIDContext {
	var p = new(EIDContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIDContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *EIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEID(s)

	default:
		return t.VisitChildren(s)
	}
}

type EVBOOLContext struct {
	ExprContext
}

func NewEVBOOLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EVBOOLContext {
	var p = new(EVBOOLContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EVBOOLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EVBOOLContext) VBOOL() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVBOOL, 0)
}

func (s *EVBOOLContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEVBOOL(s)

	default:
		return t.VisitChildren(s)
	}
}

type EMulDivContext struct {
	ExprContext
	op antlr.Token
}

func NewEMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EMulDivContext {
	var p = new(EMulDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EMulDivContext) GetOp() antlr.Token { return s.op }

func (s *EMulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *EMulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EMulDivContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EMulDivContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EMulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type ERelOrContext struct {
	ExprContext
	op antlr.Token
}

func NewERelOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ERelOrContext {
	var p = new(ERelOrContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ERelOrContext) GetOp() antlr.Token { return s.op }

func (s *ERelOrContext) SetOp(v antlr.Token) { s.op = v }

func (s *ERelOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ERelOrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ERelOrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ERelOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitERelOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EVIntegerContext struct {
	ExprContext
}

func NewEVIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EVIntegerContext {
	var p = new(EVIntegerContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EVIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EVIntegerContext) VINTEGER() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVINTEGER, 0)
}

func (s *EVIntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEVInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

type ESubAddContext struct {
	ExprContext
	op antlr.Token
}

func NewESubAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ESubAddContext {
	var p = new(ESubAddContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ESubAddContext) GetOp() antlr.Token { return s.op }

func (s *ESubAddContext) SetOp(v antlr.Token) { s.op = v }

func (s *ESubAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ESubAddContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ESubAddContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ESubAddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitESubAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENegContext struct {
	ExprContext
	op antlr.Token
}

func NewENegContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENegContext {
	var p = new(ENegContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ENegContext) GetOp() antlr.Token { return s.op }

func (s *ENegContext) SetOp(v antlr.Token) { s.op = v }

func (s *ENegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENegContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ENegContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitENeg(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAddSubContext struct {
	ExprContext
	op antlr.Token
}

func NewEAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAddSubContext {
	var p = new(EAddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EAddSubContext) GetOp() antlr.Token { return s.op }

func (s *EAddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *EAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EAddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EAddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAsAddContext struct {
	ExprContext
	op antlr.Token
}

func NewEAsAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAsAddContext {
	var p = new(EAsAddContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EAsAddContext) GetOp() antlr.Token { return s.op }

func (s *EAsAddContext) SetOp(v antlr.Token) { s.op = v }

func (s *EAsAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAsAddContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EAsAddContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EAsAddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEAsAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENotContext struct {
	ExprContext
	op antlr.Token
}

func NewENotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENotContext {
	var p = new(ENotContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ENotContext) GetOp() antlr.Token { return s.op }

func (s *ENotContext) SetOp(v antlr.Token) { s.op = v }

func (s *ENotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENotContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ENotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitENot(s)

	default:
		return t.VisitChildren(s)
	}
}

type EModuleContext struct {
	ExprContext
	op antlr.Token
}

func NewEModuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EModuleContext {
	var p = new(EModuleContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EModuleContext) GetOp() antlr.Token { return s.op }

func (s *EModuleContext) SetOp(v antlr.Token) { s.op = v }

func (s *EModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EModuleContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EModuleContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EModuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEModule(s)

	default:
		return t.VisitChildren(s)
	}
}

type EVFloatContext struct {
	ExprContext
}

func NewEVFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EVFloatContext {
	var p = new(EVFloatContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EVFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EVFloatContext) VFLOAT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVFLOAT, 0)
}

func (s *EVFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEVFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type EParentContext struct {
	ExprContext
}

func NewEParentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EParentContext {
	var p = new(EParentContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EParentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EParentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EParentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEParent(s)

	default:
		return t.VisitChildren(s)
	}
}

type ERelContext struct {
	ExprContext
	op antlr.Token
}

func NewERelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ERelContext {
	var p = new(ERelContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ERelContext) GetOp() antlr.Token { return s.op }

func (s *ERelContext) SetOp(v antlr.Token) { s.op = v }

func (s *ERelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ERelContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ERelContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ERelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitERel(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAssignContext struct {
	ExprContext
	op antlr.Token
}

func NewEAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAssignContext {
	var p = new(EAssignContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EAssignContext) GetOp() antlr.Token { return s.op }

func (s *EAssignContext) SetOp(v antlr.Token) { s.op = v }

func (s *EAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAssignContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EAssignContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type ERelAndContext struct {
	ExprContext
	op antlr.Token
}

func NewERelAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ERelAndContext {
	var p = new(ERelAndContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ERelAndContext) GetOp() antlr.Token { return s.op }

func (s *ERelAndContext) SetOp(v antlr.Token) { s.op = v }

func (s *ERelAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ERelAndContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ERelAndContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ERelAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitERelAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *TSParser_rulesParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, TSParser_rulesParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSParser_rulesParserT__1:
		localctx = NewENegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(41)

			var _m = p.Match(TSParser_rulesParserT__1)

			localctx.(*ENegContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.expr(17)
		}

	case TSParser_rulesParserT__2:
		localctx = NewENotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(43)

			var _m = p.Match(TSParser_rulesParserT__2)

			localctx.(*ENotContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.expr(16)
		}

	case TSParser_rulesParserT__13:
		localctx = NewEParentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.Match(TSParser_rulesParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.expr(0)
		}
		{
			p.SetState(47)
			p.Match(TSParser_rulesParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSParser_rulesParserVSTRING:
		localctx = NewEVStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(TSParser_rulesParserVSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSParser_rulesParserVINTEGER:
		localctx = NewEVIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.Match(TSParser_rulesParserVINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSParser_rulesParserVFLOAT:
		localctx = NewEVFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(51)
			p.Match(TSParser_rulesParserVFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSParser_rulesParserVBOOL:
		localctx = NewEVBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(TSParser_rulesParserVBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSParser_rulesParserID:
		localctx = NewEIDContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(53)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEModuleContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(57)

					var _m = p.Match(TSParser_rulesParserT__3)

					localctx.(*EModuleContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(58)
					p.expr(16)
				}

			case 2:
				localctx = NewEMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(60)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EMulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TSParser_rulesParserT__4 || _la == TSParser_rulesParserT__5) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EMulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(61)
					p.expr(15)
				}

			case 3:
				localctx = NewEAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(63)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EAddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TSParser_rulesParserT__1 || _la == TSParser_rulesParserT__6) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EAddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(64)
					p.expr(14)
				}

			case 4:
				localctx = NewERelContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(66)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ERelContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16128) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ERelContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(67)
					p.expr(13)
				}

			case 5:
				localctx = NewERelAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(69)

					var _m = p.Match(TSParser_rulesParserT__15)

					localctx.(*ERelAndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(70)
					p.expr(11)
				}

			case 6:
				localctx = NewERelOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(72)

					var _m = p.Match(TSParser_rulesParserT__16)

					localctx.(*ERelOrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(73)
					p.expr(10)
				}

			case 7:
				localctx = NewEAssignContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(75)

					var _m = p.Match(TSParser_rulesParserT__0)

					localctx.(*EAssignContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(76)
					p.expr(8)
				}

			case 8:
				localctx = NewEAsAddContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(78)

					var _m = p.Match(TSParser_rulesParserT__17)

					localctx.(*EAsAddContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(79)
					p.expr(7)
				}

			case 9:
				localctx = NewESubAddContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(81)

					var _m = p.Match(TSParser_rulesParserT__18)

					localctx.(*ESubAddContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(82)
					p.expr(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarContext is an interface to support dynamic dispatch.
type IDeclarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclarContext differentiates from other interfaces.
	IsDeclarContext()
}

type DeclarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarContext() *DeclarContext {
	var p = new(DeclarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_declar
	return p
}

func InitEmptyDeclarContext(p *DeclarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_declar
}

func (*DeclarContext) IsDeclarContext() {}

func NewDeclarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarContext {
	var p = new(DeclarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_declar

	return p
}

func (s *DeclarContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarContext) CopyAll(ctx *DeclarContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SDStrContext struct {
	DeclarContext
}

func NewSDStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDStrContext {
	var p = new(SDStrContext)

	InitEmptyDeclarContext(&p.DeclarContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarContext))

	return p
}

func (s *SDStrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDStrContext) VAR() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVAR, 0)
}

func (s *SDStrContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *SDStrContext) STRING() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserSTRING, 0)
}

func (s *SDStrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSDStr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDBoolContext struct {
	DeclarContext
}

func NewSDBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDBoolContext {
	var p = new(SDBoolContext)

	InitEmptyDeclarContext(&p.DeclarContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarContext))

	return p
}

func (s *SDBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDBoolContext) VAR() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVAR, 0)
}

func (s *SDBoolContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *SDBoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserBOOL, 0)
}

func (s *SDBoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSDBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDFltContext struct {
	DeclarContext
}

func NewSDFltContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDFltContext {
	var p = new(SDFltContext)

	InitEmptyDeclarContext(&p.DeclarContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarContext))

	return p
}

func (s *SDFltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDFltContext) VAR() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVAR, 0)
}

func (s *SDFltContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *SDFltContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserFLOAT, 0)
}

func (s *SDFltContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSDFlt(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDChrContext struct {
	DeclarContext
}

func NewSDChrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDChrContext {
	var p = new(SDChrContext)

	InitEmptyDeclarContext(&p.DeclarContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarContext))

	return p
}

func (s *SDChrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDChrContext) VAR() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVAR, 0)
}

func (s *SDChrContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *SDChrContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserCHARACTER, 0)
}

func (s *SDChrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSDChr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDIntContext struct {
	DeclarContext
}

func NewSDIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDIntContext {
	var p = new(SDIntContext)

	InitEmptyDeclarContext(&p.DeclarContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarContext))

	return p
}

func (s *SDIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDIntContext) VAR() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVAR, 0)
}

func (s *SDIntContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *SDIntContext) INT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserINT, 0)
}

func (s *SDIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSDInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Declar() (localctx IDeclarContext) {
	localctx = NewDeclarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TSParser_rulesParserRULE_declar)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSDStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(TSParser_rulesParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSDIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(TSParser_rulesParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSDFltContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(TSParser_rulesParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewSDBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(100)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(TSParser_rulesParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewSDChrContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(104)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(TSParser_rulesParserCHARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTerm returns the term token.
	GetTerm() antlr.Token

	// SetTerm sets the term token.
	SetTerm(antlr.Token)

	// Getter signatures
	Lsents() ILsentsContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	term   antlr.Token
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) GetTerm() antlr.Token { return s.term }

func (s *BlockContext) SetTerm(v antlr.Token) { s.term = v }

func (s *BlockContext) Lsents() ILsentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILsentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILsentsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TSParser_rulesParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)

		var _m = p.Match(TSParser_rulesParserT__20)

		localctx.(*BlockContext).term = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Lsents()
	}
	{
		p.SetState(112)
		p.Match(TSParser_rulesParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfContext is an interface to support dynamic dispatch.
type IIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfContext differentiates from other interfaces.
	IsIfContext()
}

type IfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfContext() *IfContext {
	var p = new(IfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_if
	return p
}

func InitEmptyIfContext(p *IfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_if
}

func (*IfContext) IsIfContext() {}

func NewIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfContext {
	var p = new(IfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_if

	return p
}

func (s *IfContext) GetParser() antlr.Parser { return s.parser }

func (s *IfContext) CopyAll(ctx *IfContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RIfEIfContext struct {
	IfContext
}

func NewRIfEIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RIfEIfContext {
	var p = new(RIfEIfContext)

	InitEmptyIfContext(&p.IfContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfContext))

	return p
}

func (s *RIfEIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RIfEIfContext) IF() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserIF, 0)
}

func (s *RIfEIfContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RIfEIfContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RIfEIfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserELSE, 0)
}

func (s *RIfEIfContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

func (s *RIfEIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitRIfEIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type RIfElseContext struct {
	IfContext
}

func NewRIfElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RIfElseContext {
	var p = new(RIfElseContext)

	InitEmptyIfContext(&p.IfContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfContext))

	return p
}

func (s *RIfElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RIfElseContext) IF() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserIF, 0)
}

func (s *RIfElseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RIfElseContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *RIfElseContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RIfElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserELSE, 0)
}

func (s *RIfElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitRIfElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type RIfContext struct {
	IfContext
}

func NewRIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RIfContext {
	var p = new(RIfContext)

	InitEmptyIfContext(&p.IfContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfContext))

	return p
}

func (s *RIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RIfContext) IF() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserIF, 0)
}

func (s *RIfContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RIfContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitRIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) If_() (localctx IIfContext) {
	localctx = NewIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TSParser_rulesParserRULE_if)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Match(TSParser_rulesParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.expr(0)
		}
		{
			p.SetState(116)
			p.Block()
		}

	case 2:
		localctx = NewRIfElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(TSParser_rulesParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.expr(0)
		}
		{
			p.SetState(120)
			p.Block()
		}
		{
			p.SetState(121)
			p.Match(TSParser_rulesParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Block()
		}

	case 3:
		localctx = NewRIfEIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Match(TSParser_rulesParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.expr(0)
		}
		{
			p.SetState(126)
			p.Block()
		}
		{
			p.SetState(127)
			p.Match(TSParser_rulesParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.If_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchContext is an interface to support dynamic dispatch.
type ISwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInput returns the input rule contexts.
	GetInput() IExprContext

	// SetInput sets the input rule contexts.
	SetInput(IExprContext)

	// Getter signatures
	SWITCH() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCASE() []antlr.TerminalNode
	CASE(i int) antlr.TerminalNode
	AllLsents() []ILsentsContext
	Lsents(i int) ILsentsContext
	Default_() IDefaultContext

	// IsSwitchContext differentiates from other interfaces.
	IsSwitchContext()
}

type SwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	input  IExprContext
}

func NewEmptySwitchContext() *SwitchContext {
	var p = new(SwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_switch
	return p
}

func InitEmptySwitchContext(p *SwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_switch
}

func (*SwitchContext) IsSwitchContext() {}

func NewSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchContext {
	var p = new(SwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_switch

	return p
}

func (s *SwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchContext) GetInput() IExprContext { return s.input }

func (s *SwitchContext) SetInput(v IExprContext) { s.input = v }

func (s *SwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserSWITCH, 0)
}

func (s *SwitchContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SwitchContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchContext) AllCASE() []antlr.TerminalNode {
	return s.GetTokens(TSParser_rulesParserCASE)
}

func (s *SwitchContext) CASE(i int) antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserCASE, i)
}

func (s *SwitchContext) AllLsents() []ILsentsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILsentsContext); ok {
			len++
		}
	}

	tst := make([]ILsentsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILsentsContext); ok {
			tst[i] = t.(ILsentsContext)
			i++
		}
	}

	return tst
}

func (s *SwitchContext) Lsents(i int) ILsentsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILsentsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILsentsContext)
}

func (s *SwitchContext) Default_() IDefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultContext)
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Switch_() (localctx ISwitchContext) {
	localctx = NewSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TSParser_rulesParserRULE_switch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(TSParser_rulesParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)

		var _x = p.expr(0)

		localctx.(*SwitchContext).input = _x
	}
	{
		p.SetState(134)
		p.Match(TSParser_rulesParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TSParser_rulesParserCASE {
		{
			p.SetState(135)
			p.Match(TSParser_rulesParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.expr(0)
		}
		{
			p.SetState(137)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Lsents()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSParser_rulesParserDEFAULT {
		{
			p.SetState(144)
			p.Default_()
		}

	}
	{
		p.SetState(147)
		p.Match(TSParser_rulesParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultContext is an interface to support dynamic dispatch.
type IDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	Lsents() ILsentsContext

	// IsDefaultContext differentiates from other interfaces.
	IsDefaultContext()
}

type DefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultContext() *DefaultContext {
	var p = new(DefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_default
	return p
}

func InitEmptyDefaultContext(p *DefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_default
}

func (*DefaultContext) IsDefaultContext() {}

func NewDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultContext {
	var p = new(DefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_default

	return p
}

func (s *DefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserDEFAULT, 0)
}

func (s *DefaultContext) Lsents() ILsentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILsentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILsentsContext)
}

func (s *DefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Default_() (localctx IDefaultContext) {
	localctx = NewDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TSParser_rulesParserRULE_default)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(TSParser_rulesParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(TSParser_rulesParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Lsents()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserPRINT, 0)
}

func (s *PrintContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TSParser_rulesParserRULE_print)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(TSParser_rulesParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(TSParser_rulesParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.expr(0)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSParser_rulesParserT__22 {
		{
			p.SetState(156)
			p.Match(TSParser_rulesParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.expr(0)
		}

		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(163)
		p.Match(TSParser_rulesParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *TSParser_rulesParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TSParser_rulesParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
