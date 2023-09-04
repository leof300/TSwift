// Code generated from TSParser_rules.g4 by ANTLR 4.13.0. DO NOT EDIT.

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
		"'-='", "':'", "", "", "'var'", "'let'", "'nil'", "'String'", "'Int'",
		"'Float'", "'Bool'", "'Character'", "'if'", "'else'", "'switch'", "'default'",
		"'case'", "'while'", "'in'", "'for'", "'guard'", "'continue'", "'break'",
		"'return'", "'struct'", "'self'", "'mutating'", "'func'", "'inout'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "NL", "WS", "VAR", "LET", "NIL", "STRING", "INT", "FLOAT",
		"BOOL", "CHARACTER", "IF", "ELSE", "SWITCH", "DEFAULT", "CASE", "WHILE",
		"IN", "FOR", "GUARD", "CONTINUE", "BREAK", "RETURN", "STRUCT", "SELF",
		"MUTATING", "FUNC", "INOUT", "VBOOL", "VSTRING", "VFLOAT", "VINTEGER",
		"ID", "SL_COMMENT", "ML_COMMENT",
	}
	staticData.RuleNames = []string{
		"start", "lsents", "sents", "expr", "declar",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 101, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 1, 4, 1, 14, 8, 1, 11, 1, 12, 1, 15, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 29, 8, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 3, 3, 45, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 74, 8, 3, 10, 3, 12, 3, 77, 9,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 99, 8, 4, 1, 4,
		0, 1, 6, 5, 0, 2, 4, 6, 8, 0, 3, 1, 0, 5, 6, 2, 0, 2, 2, 7, 7, 1, 0, 8,
		13, 119, 0, 10, 1, 0, 0, 0, 2, 13, 1, 0, 0, 0, 4, 28, 1, 0, 0, 0, 6, 44,
		1, 0, 0, 0, 8, 98, 1, 0, 0, 0, 10, 11, 3, 2, 1, 0, 11, 1, 1, 0, 0, 0, 12,
		14, 3, 4, 2, 0, 13, 12, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 13, 1, 0, 0,
		0, 15, 16, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 18, 5, 0, 0, 1, 18, 3, 1,
		0, 0, 0, 19, 20, 3, 6, 3, 0, 20, 21, 5, 21, 0, 0, 21, 29, 1, 0, 0, 0, 22,
		29, 5, 21, 0, 0, 23, 29, 3, 8, 4, 0, 24, 25, 3, 8, 4, 0, 25, 26, 5, 1,
		0, 0, 26, 27, 3, 6, 3, 0, 27, 29, 1, 0, 0, 0, 28, 19, 1, 0, 0, 0, 28, 22,
		1, 0, 0, 0, 28, 23, 1, 0, 0, 0, 28, 24, 1, 0, 0, 0, 29, 5, 1, 0, 0, 0,
		30, 31, 6, 3, -1, 0, 31, 32, 5, 2, 0, 0, 32, 45, 3, 6, 3, 17, 33, 34, 5,
		3, 0, 0, 34, 45, 3, 6, 3, 16, 35, 36, 5, 14, 0, 0, 36, 37, 3, 6, 3, 0,
		37, 38, 5, 15, 0, 0, 38, 45, 1, 0, 0, 0, 39, 45, 5, 49, 0, 0, 40, 45, 5,
		51, 0, 0, 41, 45, 5, 50, 0, 0, 42, 45, 5, 48, 0, 0, 43, 45, 5, 52, 0, 0,
		44, 30, 1, 0, 0, 0, 44, 33, 1, 0, 0, 0, 44, 35, 1, 0, 0, 0, 44, 39, 1,
		0, 0, 0, 44, 40, 1, 0, 0, 0, 44, 41, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44,
		43, 1, 0, 0, 0, 45, 75, 1, 0, 0, 0, 46, 47, 10, 15, 0, 0, 47, 48, 5, 4,
		0, 0, 48, 74, 3, 6, 3, 16, 49, 50, 10, 14, 0, 0, 50, 51, 7, 0, 0, 0, 51,
		74, 3, 6, 3, 15, 52, 53, 10, 13, 0, 0, 53, 54, 7, 1, 0, 0, 54, 74, 3, 6,
		3, 14, 55, 56, 10, 12, 0, 0, 56, 57, 7, 2, 0, 0, 57, 74, 3, 6, 3, 13, 58,
		59, 10, 10, 0, 0, 59, 60, 5, 16, 0, 0, 60, 74, 3, 6, 3, 11, 61, 62, 10,
		9, 0, 0, 62, 63, 5, 17, 0, 0, 63, 74, 3, 6, 3, 10, 64, 65, 10, 8, 0, 0,
		65, 66, 5, 1, 0, 0, 66, 74, 3, 6, 3, 8, 67, 68, 10, 7, 0, 0, 68, 69, 5,
		18, 0, 0, 69, 74, 3, 6, 3, 7, 70, 71, 10, 6, 0, 0, 71, 72, 5, 19, 0, 0,
		72, 74, 3, 6, 3, 6, 73, 46, 1, 0, 0, 0, 73, 49, 1, 0, 0, 0, 73, 52, 1,
		0, 0, 0, 73, 55, 1, 0, 0, 0, 73, 58, 1, 0, 0, 0, 73, 61, 1, 0, 0, 0, 73,
		64, 1, 0, 0, 0, 73, 67, 1, 0, 0, 0, 73, 70, 1, 0, 0, 0, 74, 77, 1, 0, 0,
		0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 7, 1, 0, 0, 0, 77, 75, 1,
		0, 0, 0, 78, 79, 5, 23, 0, 0, 79, 80, 5, 52, 0, 0, 80, 81, 5, 20, 0, 0,
		81, 99, 5, 26, 0, 0, 82, 83, 5, 23, 0, 0, 83, 84, 5, 52, 0, 0, 84, 85,
		5, 20, 0, 0, 85, 99, 5, 27, 0, 0, 86, 87, 5, 23, 0, 0, 87, 88, 5, 52, 0,
		0, 88, 89, 5, 20, 0, 0, 89, 99, 5, 28, 0, 0, 90, 91, 5, 23, 0, 0, 91, 92,
		5, 52, 0, 0, 92, 93, 5, 20, 0, 0, 93, 99, 5, 29, 0, 0, 94, 95, 5, 23, 0,
		0, 95, 96, 5, 52, 0, 0, 96, 97, 5, 20, 0, 0, 97, 99, 5, 30, 0, 0, 98, 78,
		1, 0, 0, 0, 98, 82, 1, 0, 0, 0, 98, 86, 1, 0, 0, 0, 98, 90, 1, 0, 0, 0,
		98, 94, 1, 0, 0, 0, 99, 9, 1, 0, 0, 0, 6, 15, 28, 44, 73, 75, 98,
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
	TSParser_rulesParserNL         = 21
	TSParser_rulesParserWS         = 22
	TSParser_rulesParserVAR        = 23
	TSParser_rulesParserLET        = 24
	TSParser_rulesParserNIL        = 25
	TSParser_rulesParserSTRING     = 26
	TSParser_rulesParserINT        = 27
	TSParser_rulesParserFLOAT      = 28
	TSParser_rulesParserBOOL       = 29
	TSParser_rulesParserCHARACTER  = 30
	TSParser_rulesParserIF         = 31
	TSParser_rulesParserELSE       = 32
	TSParser_rulesParserSWITCH     = 33
	TSParser_rulesParserDEFAULT    = 34
	TSParser_rulesParserCASE       = 35
	TSParser_rulesParserWHILE      = 36
	TSParser_rulesParserIN         = 37
	TSParser_rulesParserFOR        = 38
	TSParser_rulesParserGUARD      = 39
	TSParser_rulesParserCONTINUE   = 40
	TSParser_rulesParserBREAK      = 41
	TSParser_rulesParserRETURN     = 42
	TSParser_rulesParserSTRUCT     = 43
	TSParser_rulesParserSELF       = 44
	TSParser_rulesParserMUTATING   = 45
	TSParser_rulesParserFUNC       = 46
	TSParser_rulesParserINOUT      = 47
	TSParser_rulesParserVBOOL      = 48
	TSParser_rulesParserVSTRING    = 49
	TSParser_rulesParserVFLOAT     = 50
	TSParser_rulesParserVINTEGER   = 51
	TSParser_rulesParserID         = 52
	TSParser_rulesParserSL_COMMENT = 53
	TSParser_rulesParserML_COMMENT = 54
)

// TSParser_rulesParser rules.
const (
	TSParser_rulesParserRULE_start  = 0
	TSParser_rulesParserRULE_lsents = 1
	TSParser_rulesParserRULE_sents  = 2
	TSParser_rulesParserRULE_expr   = 3
	TSParser_rulesParserRULE_declar = 4
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lsents() ILsentsContext

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
		p.SetState(10)
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

// ILsentsContext is an interface to support dynamic dispatch.
type ILsentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
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

func (s *LsentsContext) EOF() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserEOF, 0)
}

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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8725724288532492) != 0) {
		{
			p.SetState(12)
			p.Sents()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(17)
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

func (s *SentExprContext) NL() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserNL, 0)
}

func (s *SentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SentNLContext struct {
	SentsContext
}

func NewSentNLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentNLContext {
	var p = new(SentNLContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SentNLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentNLContext) NL() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserNL, 0)
}

func (s *SentNLContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSentNL(s)

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
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSentExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
			p.expr(0)
		}
		{
			p.SetState(20)
			p.Match(TSParser_rulesParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSentNLContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(22)
			p.Match(TSParser_rulesParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(23)
			p.Declar()
		}

	case 4:
		localctx = NewSDeclAsigContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(24)
			p.Declar()
		}
		{
			p.SetState(25)

			var _m = p.Match(TSParser_rulesParserT__0)

			localctx.(*SDeclAsigContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
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
	p.SetState(44)
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
			p.SetState(31)

			var _m = p.Match(TSParser_rulesParserT__1)

			localctx.(*ENegContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.expr(17)
		}

	case TSParser_rulesParserT__2:
		localctx = NewENotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(33)

			var _m = p.Match(TSParser_rulesParserT__2)

			localctx.(*ENotContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.expr(16)
		}

	case TSParser_rulesParserT__13:
		localctx = NewEParentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(35)
			p.Match(TSParser_rulesParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.expr(0)
		}
		{
			p.SetState(37)
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
			p.SetState(39)
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
			p.SetState(40)
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
			p.SetState(41)
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
			p.SetState(42)
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
			p.SetState(43)
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
	p.SetState(75)
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
			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEModuleContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(47)

					var _m = p.Match(TSParser_rulesParserT__3)

					localctx.(*EModuleContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(48)
					p.expr(16)
				}

			case 2:
				localctx = NewEMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(49)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(50)

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
					p.SetState(51)
					p.expr(15)
				}

			case 3:
				localctx = NewEAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(53)

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
					p.SetState(54)
					p.expr(14)
				}

			case 4:
				localctx = NewERelContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(55)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(56)

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
					p.SetState(57)
					p.expr(13)
				}

			case 5:
				localctx = NewERelAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(58)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(59)

					var _m = p.Match(TSParser_rulesParserT__15)

					localctx.(*ERelAndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(60)
					p.expr(11)
				}

			case 6:
				localctx = NewERelOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(62)

					var _m = p.Match(TSParser_rulesParserT__16)

					localctx.(*ERelOrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(63)
					p.expr(10)
				}

			case 7:
				localctx = NewEAssignContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(65)

					var _m = p.Match(TSParser_rulesParserT__0)

					localctx.(*EAssignContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(66)
					p.expr(8)
				}

			case 8:
				localctx = NewEAsAddContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(68)

					var _m = p.Match(TSParser_rulesParserT__17)

					localctx.(*EAsAddContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(69)
					p.expr(7)
				}

			case 9:
				localctx = NewESubAddContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(71)

					var _m = p.Match(TSParser_rulesParserT__18)

					localctx.(*ESubAddContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(72)
					p.expr(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(77)
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
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSDStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
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
			p.SetState(82)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
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
			p.SetState(86)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
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
			p.SetState(90)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
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
			p.SetState(94)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
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
