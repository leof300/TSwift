// Code generated from TSParser_rules.g4 by ANTLR 4.13.1. DO NOT EDIT.

package TSVisitor // TSParser_rules
import "github.com/antlr4-go/antlr/v4"

type BaseTSParser_rulesVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTSParser_rulesVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitLins(ctx *LinsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitIFunc(ctx *IFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitIDecl(ctx *IDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitICons(ctx *IConsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSCallFunction(ctx *SCallFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitIPrint(ctx *IPrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitLsents(ctx *LsentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSIf(ctx *SIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSSwitch(ctx *SSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSWhile(ctx *SWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSGuard(ctx *SGuardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSFor(ctx *SForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSDecl(ctx *SDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSCons(ctx *SConsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSPrint(ctx *SPrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSentExpr(ctx *SentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSentTrans(ctx *SentTransContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEVString(ctx *EVStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEID(ctx *EIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEVBOOL(ctx *EVBOOLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEMulDiv(ctx *EMulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEFunction(ctx *EFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitERelOr(ctx *ERelOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEVInteger(ctx *EVIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitESubAdd(ctx *ESubAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitENeg(ctx *ENegContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitENIL(ctx *ENILContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEAddSub(ctx *EAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEAsAdd(ctx *EAsAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitENot(ctx *ENotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEModule(ctx *EModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEVArray(ctx *EVArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEVFloat(ctx *EVFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEParent(ctx *EParentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitERel(ctx *ERelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEAssign(ctx *EAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitETSFunctions(ctx *ETSFunctionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitERelAnd(ctx *ERelAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSDType(ctx *SDTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSDecAssign(ctx *SDecAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSDecTAssign(ctx *SDecTAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitArrayValue(ctx *ArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSConsAss(ctx *SConsAssContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSConsTypeAss(ctx *SConsTypeAssContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitTstypes(ctx *TstypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitRIf(ctx *RIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitRIfElse(ctx *RIfElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitRIfEIf(ctx *RIfEIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSwitch(ctx *SwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitDefault(ctx *DefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitWhile(ctx *WhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitGuard(ctx *GuardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitFor(ctx *ForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitRango(ctx *RangoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitStrans(ctx *StransContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitTsprint(ctx *TsprintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitFunctions(ctx *FunctionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitCallFunction(ctx *CallFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitCallParameter(ctx *CallParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitConvertString(ctx *ConvertStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitConvertInt(ctx *ConvertIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitConvertFloat(ctx *ConvertFloatContext) interface{} {
	return v.VisitChildren(ctx)
}
