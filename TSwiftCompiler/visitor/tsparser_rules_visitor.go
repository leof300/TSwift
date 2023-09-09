// Code generated from TSParser_rules.g4 by ANTLR 4.13.1. DO NOT EDIT.

package TSVisitor // TSParser_rules
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TSParser_rulesParser.
type TSParser_rulesVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TSParser_rulesParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#lins.
	VisitLins(ctx *LinsContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#IFunc.
	VisitIFunc(ctx *IFuncContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#IDecl.
	VisitIDecl(ctx *IDeclContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ICons.
	VisitICons(ctx *IConsContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SCallFunction.
	VisitSCallFunction(ctx *SCallFunctionContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#IPrint.
	VisitIPrint(ctx *IPrintContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#lsents.
	VisitLsents(ctx *LsentsContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SIf.
	VisitSIf(ctx *SIfContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SSwitch.
	VisitSSwitch(ctx *SSwitchContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SWhile.
	VisitSWhile(ctx *SWhileContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SGuard.
	VisitSGuard(ctx *SGuardContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SFor.
	VisitSFor(ctx *SForContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SDecl.
	VisitSDecl(ctx *SDeclContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SCons.
	VisitSCons(ctx *SConsContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SPrint.
	VisitSPrint(ctx *SPrintContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SentExpr.
	VisitSentExpr(ctx *SentExprContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SentTrans.
	VisitSentTrans(ctx *SentTransContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVString.
	VisitEVString(ctx *EVStringContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EID.
	VisitEID(ctx *EIDContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVBOOL.
	VisitEVBOOL(ctx *EVBOOLContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EMulDiv.
	VisitEMulDiv(ctx *EMulDivContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EFunction.
	VisitEFunction(ctx *EFunctionContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ERelOr.
	VisitERelOr(ctx *ERelOrContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVInteger.
	VisitEVInteger(ctx *EVIntegerContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ESubAdd.
	VisitESubAdd(ctx *ESubAddContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ENeg.
	VisitENeg(ctx *ENegContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ENIL.
	VisitENIL(ctx *ENILContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EAddSub.
	VisitEAddSub(ctx *EAddSubContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EAsAdd.
	VisitEAsAdd(ctx *EAsAddContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ENot.
	VisitENot(ctx *ENotContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EModule.
	VisitEModule(ctx *EModuleContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVArray.
	VisitEVArray(ctx *EVArrayContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EVFloat.
	VisitEVFloat(ctx *EVFloatContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EParent.
	VisitEParent(ctx *EParentContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ERel.
	VisitERel(ctx *ERelContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#EAssign.
	VisitEAssign(ctx *EAssignContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ETSFunctions.
	VisitETSFunctions(ctx *ETSFunctionsContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ERelAnd.
	VisitERelAnd(ctx *ERelAndContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SDType.
	VisitSDType(ctx *SDTypeContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SDecAssign.
	VisitSDecAssign(ctx *SDecAssignContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SDecTAssign.
	VisitSDecTAssign(ctx *SDecTAssignContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#arrayValue.
	VisitArrayValue(ctx *ArrayValueContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SConsAss.
	VisitSConsAss(ctx *SConsAssContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#SConsTypeAss.
	VisitSConsTypeAss(ctx *SConsTypeAssContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#tstypes.
	VisitTstypes(ctx *TstypesContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#RIf.
	VisitRIf(ctx *RIfContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#RIfElse.
	VisitRIfElse(ctx *RIfElseContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#RIfEIf.
	VisitRIfEIf(ctx *RIfEIfContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#switch.
	VisitSwitch(ctx *SwitchContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#default.
	VisitDefault(ctx *DefaultContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#while.
	VisitWhile(ctx *WhileContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#guard.
	VisitGuard(ctx *GuardContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#for.
	VisitFor(ctx *ForContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#rango.
	VisitRango(ctx *RangoContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#strans.
	VisitStrans(ctx *StransContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#tsprint.
	VisitTsprint(ctx *TsprintContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#functions.
	VisitFunctions(ctx *FunctionsContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#callFunction.
	VisitCallFunction(ctx *CallFunctionContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#callParameter.
	VisitCallParameter(ctx *CallParameterContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ConvertString.
	VisitConvertString(ctx *ConvertStringContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ConvertInt.
	VisitConvertInt(ctx *ConvertIntContext) interface{}

	// Visit a parse tree produced by TSParser_rulesParser#ConvertFloat.
	VisitConvertFloat(ctx *ConvertFloatContext) interface{}
}
