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
		"", "'%'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'='", "", "",
		"'var'", "'let'", "'nil'", "'String'", "'Int'", "'Float'", "'Bool'",
		"'Character'", "'if'", "'else'", "'switch'", "'default'", "'case'",
		"'while'", "'in'", "'for'", "'guard'", "'continue'", "'break'", "'return'",
		"'struct'", "'self'", "'mutating'", "'func'", "'inout'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "NL", "WS", "VAR", "LET", "NIL",
		"STRING", "INT", "FLOAT", "BOOL", "CHARACTER", "IF", "ELSE", "SWITCH",
		"DEFAULT", "CASE", "WHILE", "IN", "FOR", "GUARD", "CONTINUE", "BREAK",
		"RETURN", "STRUCT", "SELF", "MUTATING", "FUNC", "INOUT", "VBOOL", "VSTRING",
		"VFLOAT", "VINTEGER", "ID", "SL_COMMENT", "ML_COMMENT",
	}
	staticData.RuleNames = []string{
		"start", "lsents", "sents", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 53, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 1, 4, 1, 12, 8, 1, 11, 1, 12, 1, 13, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 22, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 34, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 48, 8, 3, 10, 3, 12, 3, 51, 9, 3, 1, 3, 0,
		1, 6, 4, 0, 2, 4, 6, 0, 2, 1, 0, 2, 3, 1, 0, 4, 5, 59, 0, 8, 1, 0, 0, 0,
		2, 11, 1, 0, 0, 0, 4, 21, 1, 0, 0, 0, 6, 33, 1, 0, 0, 0, 8, 9, 3, 2, 1,
		0, 9, 1, 1, 0, 0, 0, 10, 12, 3, 4, 2, 0, 11, 10, 1, 0, 0, 0, 12, 13, 1,
		0, 0, 0, 13, 11, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15,
		16, 5, 0, 0, 1, 16, 3, 1, 0, 0, 0, 17, 18, 3, 6, 3, 0, 18, 19, 5, 9, 0,
		0, 19, 22, 1, 0, 0, 0, 20, 22, 5, 9, 0, 0, 21, 17, 1, 0, 0, 0, 21, 20,
		1, 0, 0, 0, 22, 5, 1, 0, 0, 0, 23, 24, 6, 3, -1, 0, 24, 25, 5, 6, 0, 0,
		25, 26, 3, 6, 3, 0, 26, 27, 5, 7, 0, 0, 27, 34, 1, 0, 0, 0, 28, 34, 5,
		37, 0, 0, 29, 34, 5, 39, 0, 0, 30, 34, 5, 38, 0, 0, 31, 34, 5, 36, 0, 0,
		32, 34, 5, 40, 0, 0, 33, 23, 1, 0, 0, 0, 33, 28, 1, 0, 0, 0, 33, 29, 1,
		0, 0, 0, 33, 30, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 32, 1, 0, 0, 0, 34,
		49, 1, 0, 0, 0, 35, 36, 10, 10, 0, 0, 36, 37, 5, 1, 0, 0, 37, 48, 3, 6,
		3, 11, 38, 39, 10, 9, 0, 0, 39, 40, 7, 0, 0, 0, 40, 48, 3, 6, 3, 10, 41,
		42, 10, 8, 0, 0, 42, 43, 7, 1, 0, 0, 43, 48, 3, 6, 3, 9, 44, 45, 10, 6,
		0, 0, 45, 46, 5, 8, 0, 0, 46, 48, 3, 6, 3, 7, 47, 35, 1, 0, 0, 0, 47, 38,
		1, 0, 0, 0, 47, 41, 1, 0, 0, 0, 47, 44, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0,
		49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 7, 1, 0, 0, 0, 51, 49, 1, 0,
		0, 0, 5, 13, 21, 33, 47, 49,
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
	TSParser_rulesParserNL         = 9
	TSParser_rulesParserWS         = 10
	TSParser_rulesParserVAR        = 11
	TSParser_rulesParserLET        = 12
	TSParser_rulesParserNIL        = 13
	TSParser_rulesParserSTRING     = 14
	TSParser_rulesParserINT        = 15
	TSParser_rulesParserFLOAT      = 16
	TSParser_rulesParserBOOL       = 17
	TSParser_rulesParserCHARACTER  = 18
	TSParser_rulesParserIF         = 19
	TSParser_rulesParserELSE       = 20
	TSParser_rulesParserSWITCH     = 21
	TSParser_rulesParserDEFAULT    = 22
	TSParser_rulesParserCASE       = 23
	TSParser_rulesParserWHILE      = 24
	TSParser_rulesParserIN         = 25
	TSParser_rulesParserFOR        = 26
	TSParser_rulesParserGUARD      = 27
	TSParser_rulesParserCONTINUE   = 28
	TSParser_rulesParserBREAK      = 29
	TSParser_rulesParserRETURN     = 30
	TSParser_rulesParserSTRUCT     = 31
	TSParser_rulesParserSELF       = 32
	TSParser_rulesParserMUTATING   = 33
	TSParser_rulesParserFUNC       = 34
	TSParser_rulesParserINOUT      = 35
	TSParser_rulesParserVBOOL      = 36
	TSParser_rulesParserVSTRING    = 37
	TSParser_rulesParserVFLOAT     = 38
	TSParser_rulesParserVINTEGER   = 39
	TSParser_rulesParserID         = 40
	TSParser_rulesParserSL_COMMENT = 41
	TSParser_rulesParserML_COMMENT = 42
)

// TSParser_rulesParser rules.
const (
	TSParser_rulesParserRULE_start  = 0
	TSParser_rulesParserRULE_lsents = 1
	TSParser_rulesParserRULE_sents  = 2
	TSParser_rulesParserRULE_expr   = 3
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
		p.SetState(8)
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
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2130303779392) != 0) {
		{
			p.SetState(10)
			p.Sents()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(15)
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

func (p *TSParser_rulesParser) Sents() (localctx ISentsContext) {
	localctx = NewSentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TSParser_rulesParserRULE_sents)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSParser_rulesParserT__5, TSParser_rulesParserVBOOL, TSParser_rulesParserVSTRING, TSParser_rulesParserVFLOAT, TSParser_rulesParserVINTEGER, TSParser_rulesParserID:
		localctx = NewSentExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(17)
			p.expr(0)
		}
		{
			p.SetState(18)
			p.Match(TSParser_rulesParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSParser_rulesParserNL:
		localctx = NewSentNLContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(20)
			p.Match(TSParser_rulesParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
}

func NewEMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EMulDivContext {
	var p = new(EMulDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

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

type EAssignContext struct {
	ExprContext
}

func NewEAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAssignContext {
	var p = new(EAssignContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

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

type EAddSubContext struct {
	ExprContext
}

func NewEAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAddSubContext {
	var p = new(EAddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

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

type EModuleContext struct {
	ExprContext
}

func NewEModuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EModuleContext {
	var p = new(EModuleContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSParser_rulesParserT__5:
		localctx = NewEParentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(24)
			p.Match(TSParser_rulesParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.expr(0)
		}
		{
			p.SetState(26)
			p.Match(TSParser_rulesParserT__6)
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
			p.SetState(28)
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
			p.SetState(29)
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
			p.SetState(30)
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
			p.SetState(31)
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
			p.SetState(32)
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
	p.SetState(49)
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
			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEModuleContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(36)
					p.Match(TSParser_rulesParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(37)
					p.expr(11)
				}

			case 2:
				localctx = NewEMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(39)
					_la = p.GetTokenStream().LA(1)

					if !(_la == TSParser_rulesParserT__1 || _la == TSParser_rulesParserT__2) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(40)
					p.expr(10)
				}

			case 3:
				localctx = NewEAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(42)
					_la = p.GetTokenStream().LA(1)

					if !(_la == TSParser_rulesParserT__3 || _la == TSParser_rulesParserT__4) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(43)
					p.expr(9)
				}

			case 4:
				localctx = NewEAssignContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(45)
					p.Match(TSParser_rulesParserT__7)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(46)
					p.expr(7)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(51)
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
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
