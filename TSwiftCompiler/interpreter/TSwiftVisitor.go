package interpreter

import (
	TSExceptions "TSwiftCompiler/ast/Exceptions"
	"TSwiftCompiler/ast/NTExpression"
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
	"TSwiftCompiler/visitor"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"reflect"
	"strconv"
)

type TSwiftVisitor struct {
	antlr.ParseTreeVisitor
	Start TSStructs.TSExpressioner
}

func (T TSwiftVisitor) VisitSDecl(ctx *TSVisitor.SDeclContext) interface{} {
	return ctx.Declar().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitSDeclAsig(ctx *TSVisitor.SDeclAsigContext) interface{} {
	variable := ctx.Declar().Accept(T).(TSStructs.TSExpressioner)
	content := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIAssignation(ctx.GetOp().GetLine(), 0, variable, content)
}

func (T TSwiftVisitor) VisitSDStr(ctx *TSVisitor.SDStrContext) interface{} {
	return NTExpression.NewIBoxCreation(ctx.ID().GetSymbol().GetLine(), 0, nil, TExpression.STRING, ctx.ID().GetText())
}

func (T TSwiftVisitor) VisitSDInt(ctx *TSVisitor.SDIntContext) interface{} {
	return NTExpression.NewIBoxCreation(ctx.ID().GetSymbol().GetLine(), 0, nil, TExpression.INTEGER, ctx.ID().GetText())

}

func (T TSwiftVisitor) VisitSDFlt(ctx *TSVisitor.SDFltContext) interface{} {
	return NTExpression.NewIBoxCreation(ctx.ID().GetSymbol().GetLine(), 0, nil, TExpression.FLOAT, ctx.ID().GetText())

}

func (T TSwiftVisitor) VisitSDBool(ctx *TSVisitor.SDBoolContext) interface{} {
	return NTExpression.NewIBoxCreation(ctx.ID().GetSymbol().GetLine(), 0, nil, TExpression.BOOL, ctx.ID().GetText())

}

func (T TSwiftVisitor) VisitSDChr(ctx *TSVisitor.SDChrContext) interface{} {
	//TODO CAMBIAR LO DEL CHAR, NO ES STRING
	return NTExpression.NewIBoxCreation(ctx.ID().GetSymbol().GetLine(), 0, nil, TExpression.STRING, ctx.ID().GetText())
}

// expr : VBOOL
func (T TSwiftVisitor) VisitEVBOOL(ctx *TSVisitor.EVBOOLContext) interface{} {
	//todo: errores de conversion de datos
	bvalue, _ := strconv.ParseBool(ctx.VBOOL().GetText())
	return NTExpression.NewIBooleanCreation(ctx.VBOOL().GetSymbol().GetLine(),
		ctx.VBOOL().GetSymbol().GetColumn(),
		bvalue)
}

// start : sents+ EOF
func (T TSwiftVisitor) VisitLsents(ctx *TSVisitor.LsentsContext) interface{} {
	lsentences := NTExpression.NewILSentences()
	sentences := ctx.AllSents()
	for _, sentence := range sentences {
		//fmt.Println(reflect.TypeOf(sentence).String())
		//reflect.TypeOf(sentence).Key() != TSVisitor.T
		//No vamos a visitar los nodos que vienen vacios NLContext
		if reflect.TypeOf(sentence).String() != "*TSVisitor.SentNLContext" {
			nSentence := sentence.Accept(T).(TSStructs.TSExpressioner)
			lsentences.Sentences = append(lsentences.Sentences, nSentence)
		}
	}
	return lsentences
}

// sents : expr NL #SentExpr
func (T TSwiftVisitor) VisitSentExpr(ctx *TSVisitor.SentExprContext) interface{} {
	return ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitSentNL(ctx *TSVisitor.SentNLContext) interface{} {
	return nil
}

// //////////////////////////////// START////////////////////////////////////////////////////////////////////

func (T TSwiftVisitor) VisitStart(ctx *TSVisitor.StartContext) interface{} {
	T.Start = ctx.Lsents().Accept(T).(TSStructs.TSExpressioner)
	return T.Start
}

/*//////////////////////////////////////////////////////////////////////////////////////////////////////////*/

/************* VALUES *****************/
// expr ->  VFLOAT
func (T TSwiftVisitor) VisitEVFloat(ctx *TSVisitor.EVFloatContext) interface{} {
	//todo: errores de conversion de datos
	fvalue, _ := strconv.ParseFloat(ctx.VFLOAT().GetText(), 64)
	return NTExpression.NewIFloatCreation(ctx.VFLOAT().GetSymbol().GetLine(),
		ctx.VFLOAT().GetSymbol().GetColumn(),
		fvalue)
}

// expr ->  VSTRING
func (T TSwiftVisitor) VisitEVString(ctx *TSVisitor.EVStringContext) interface{} {
	return NTExpression.NewIStringCreation(ctx.VSTRING().GetSymbol().GetLine(),
		ctx.VSTRING().GetSymbol().GetColumn(),
		ctx.VSTRING().GetText())
}

// expr ->  VInteger
func (T TSwiftVisitor) VisitEVInteger(ctx *TSVisitor.EVIntegerContext) interface{} {
	//todo: errores de conversion de datos
	ivalue, _ := strconv.Atoi(ctx.VINTEGER().GetText())
	return NTExpression.NewIIntCreation(ctx.VINTEGER().GetSymbol().GetLine(),
		ctx.VINTEGER().GetSymbol().GetColumn(),
		ivalue)
}

// expr -> ID
func (T TSwiftVisitor) VisitEID(ctx *TSVisitor.EIDContext) interface{} {
	//return NTExpression.NewIBox(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), ctx.ID().GetText())
	return NTExpression.NewIBoxGetting(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), ctx.ID().GetText())
}

func (T TSwiftVisitor) VisitEParent(ctx *TSVisitor.EParentContext) interface{} {
	return ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
}

/************* ASIGNACION ****************

 */

func (T TSwiftVisitor) VisitEAssign(ctx *TSVisitor.EAssignContext) interface{} {
	variable := ctx.Expr(0).Accept(T).(TSStructs.TSExpressioner)
	content := ctx.Expr(1).Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIAssignation(ctx.GetOp().GetLine(), 0, variable, content)
}

func (T TSwiftVisitor) VisitESubAdd(ctx *TSVisitor.ESubAddContext) interface{} {
	variable := ctx.Expr(0).Accept(T).(TSStructs.TSExpressioner)
	content := ctx.Expr(1).Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIAssignSub(ctx.GetOp().GetLine(), 0, variable, content)
}

func (T TSwiftVisitor) VisitENeg(ctx *TSVisitor.ENegContext) interface{} {
	e := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIOpNegation(ctx.GetOp().GetLine(), 0, e)
}

func (T TSwiftVisitor) VisitEAsAdd(ctx *TSVisitor.EAsAddContext) interface{} {
	variable := ctx.Expr(0).Accept(T).(TSStructs.TSExpressioner)
	content := ctx.Expr(1).Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIAssignAdd(ctx.GetOp().GetLine(), 0, variable, content)
}

/*
************** OPERACIONES ****************
 */
func (T TSwiftVisitor) VisitEMulDiv(ctx *TSVisitor.EMulDivContext) interface{} {

	e1 := ctx.Expr(0).Accept(T).(TSStructs.TSExpressioner)
	e2 := ctx.Expr(1).Accept(T).(TSStructs.TSExpressioner)
	if ctx.GetOp().GetText() == "*" {
		return NTExpression.NewIMultiplication(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
	}

	return NTExpression.NewIOpDivision(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
}

func (T TSwiftVisitor) VisitEAddSub(ctx *TSVisitor.EAddSubContext) interface{} {

	e1 := ctx.Expr(0).Accept(T).(TSStructs.TSExpressioner)
	e2 := ctx.Expr(1).Accept(T).(TSStructs.TSExpressioner)
	if ctx.GetOp().GetText() == "+" {
		return NTExpression.NewIAdd(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
	}

	return NTExpression.NewISubtraction(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
}

func (T TSwiftVisitor) VisitEModule(ctx *TSVisitor.EModuleContext) interface{} {

	e1 := ctx.Expr(0).Accept(T).(TSStructs.TSExpressioner)
	e2 := ctx.Expr(1).Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIModulo(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
}

/**
************** OPERACIONES RELACIONALES ****************
 */

func (T TSwiftVisitor) VisitERel(ctx *TSVisitor.ERelContext) interface{} {
	e1 := ctx.Expr(0).Accept(T).(TSStructs.TSExpressioner)
	e2 := ctx.Expr(1).Accept(T).(TSStructs.TSExpressioner)
	if ctx.GetOp().GetText() == "<" {
		return NTExpression.NewIRelLess(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
	}
	if ctx.GetOp().GetText() == ">" {
		return NTExpression.NewIRelGreater(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
	}
	if ctx.GetOp().GetText() == "<=" {
		return NTExpression.NewIRelLessEqual(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
	}
	if ctx.GetOp().GetText() == ">=" {
		return NTExpression.NewIRelGreaterEqual(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
	}
	if ctx.GetOp().GetText() == "==" {
		return NTExpression.NewIRelEqual(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
	}
	if ctx.GetOp().GetText() == "!=" {
		return NTExpression.NewIRelNotEqual(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
	}
	return NTExpression.NewIRelLess(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
}

func (T TSwiftVisitor) VisitENot(ctx *TSVisitor.ENotContext) interface{} {
	e1 := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIBitNot(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1)
}

func (T TSwiftVisitor) VisitERelOr(ctx *TSVisitor.ERelOrContext) interface{} {
	e1 := ctx.Expr(0).Accept(T).(TSStructs.TSExpressioner)
	e2 := ctx.Expr(1).Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIBitOr(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
}

func (T TSwiftVisitor) VisitERelAnd(ctx *TSVisitor.ERelAndContext) interface{} {
	e1 := ctx.Expr(0).Accept(T).(TSStructs.TSExpressioner)
	e2 := ctx.Expr(1).Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIBitAnd(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), e1, e2)
}

// ////////////////////////////////////////////////////////////////////////////////////
func NewTSwiftVisitor() TSVisitor.TSParser_rulesVisitor {
	return &TSwiftVisitor{ParseTreeVisitor: &TSVisitor.BaseTSParser_rulesVisitor{}}
}

func (v TSwiftVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		//log.Fatal(val.GetText())
		return TSExceptions.NewTSException(val.GetText(), 0, 0)
	default:
		//comprobamos el el nodo sea valido
		//TODO
		nodoIntepreter, ok := tree.Accept(v).(TSStructs.TSExpressioner)
		if ok {
			return nodoIntepreter
		}
		fmt.Print("nodo desconocido")
		return TSExceptions.NewTSException("ERROR NODO NO VALIDO", 0, 0)
	}
}
