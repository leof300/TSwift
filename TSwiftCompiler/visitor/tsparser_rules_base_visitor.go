// Code generated from TSParser_rules.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TSVisitor // TSParser_rules
import "github.com/antlr4-go/antlr/v4"

type BaseTSParser_rulesVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTSParser_rulesVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitLsents(ctx *LsentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSentExpr(ctx *SentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSentNL(ctx *SentNLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSDecl(ctx *SDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSDeclAsig(ctx *SDeclAsigContext) interface{} {
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

func (v *BaseTSParser_rulesVisitor) VisitEVInteger(ctx *EVIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitESubAdd(ctx *ESubAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitENeg(ctx *ENegContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEAddSub(ctx *EAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEAsAdd(ctx *EAsAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEModule(ctx *EModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEVFloat(ctx *EVFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEParent(ctx *EParentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitEAssign(ctx *EAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSDStr(ctx *SDStrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSDInt(ctx *SDIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSDFlt(ctx *SDFltContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSDBool(ctx *SDBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSParser_rulesVisitor) VisitSDChr(ctx *SDChrContext) interface{} {
	return v.VisitChildren(ctx)
}
