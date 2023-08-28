// Code generated from TSParser_rules.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TSVisitor // TSParser_rules
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TSParser_rulesParser.
type TSParser_rulesVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TSParser_rulesParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#lsents.
	VisitLsents(ctx *LsentsContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SentExpr.
	VisitSentExpr(ctx *SentExprContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SentNL.
	VisitSentNL(ctx *SentNLContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVFloat.
	VisitEVFloat(ctx *EVFloatContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVString.
	VisitEVString(ctx *EVStringContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EID.
	VisitEID(ctx *EIDContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EParent.
	VisitEParent(ctx *EParentContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVBOOL.
	VisitEVBOOL(ctx *EVBOOLContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EMulDiv.
	VisitEMulDiv(ctx *EMulDivContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EAssign.
	VisitEAssign(ctx *EAssignContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVInteger.
	VisitEVInteger(ctx *EVIntegerContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EAddSub.
	VisitEAddSub(ctx *EAddSubContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EModule.
	VisitEModule(ctx *EModuleContext) interface{}
}
