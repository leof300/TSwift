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

	// Visit a parse tree produced by TSParser_rulesParser#SDecl.
	VisitSDecl(ctx *SDeclContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SDeclAsig.
	VisitSDeclAsig(ctx *SDeclAsigContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVString.
	VisitEVString(ctx *EVStringContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EID.
	VisitEID(ctx *EIDContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVBOOL.
	VisitEVBOOL(ctx *EVBOOLContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EMulDiv.
	VisitEMulDiv(ctx *EMulDivContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ERelOr.
	VisitERelOr(ctx *ERelOrContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVInteger.
	VisitEVInteger(ctx *EVIntegerContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ESubAdd.
	VisitESubAdd(ctx *ESubAddContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ENeg.
	VisitENeg(ctx *ENegContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EAddSub.
	VisitEAddSub(ctx *EAddSubContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EAsAdd.
	VisitEAsAdd(ctx *EAsAddContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ENot.
	VisitENot(ctx *ENotContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EModule.
	VisitEModule(ctx *EModuleContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVFloat.
	VisitEVFloat(ctx *EVFloatContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EParent.
	VisitEParent(ctx *EParentContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ERel.
	VisitERel(ctx *ERelContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EAssign.
	VisitEAssign(ctx *EAssignContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ERelAnd.
	VisitERelAnd(ctx *ERelAndContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SDStr.
	VisitSDStr(ctx *SDStrContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SDInt.
	VisitSDInt(ctx *SDIntContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SDFlt.
	VisitSDFlt(ctx *SDFltContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SDBool.
	VisitSDBool(ctx *SDBoolContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SDChr.
	VisitSDChr(ctx *SDChrContext) interface{}
}
