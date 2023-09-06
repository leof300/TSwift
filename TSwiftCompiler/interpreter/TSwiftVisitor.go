package interpreter

import (
	TSExceptions "TSwiftCompiler/ast/Exceptions"
	"TSwiftCompiler/ast/NTExpression"
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
	"TSwiftCompiler/visitor"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type TSwiftVisitor struct {
	antlr.ParseTreeVisitor
	Start TSStructs.TSExpressioner
}

func (T TSwiftVisitor) VisitBlock(ctx *TSVisitor.BlockContext) interface{} {
	sentences := ctx.Lsents().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewITSBlock(ctx.GetTerm().GetLine(), 0, sentences)
}

// start : sents+ EOF
func (T TSwiftVisitor) VisitLsents(ctx *TSVisitor.LsentsContext) interface{} {
	lsentences := NTExpression.NewILSentences()
	sentences := ctx.AllSents()
	for _, sentence := range sentences {
		//fmt.Println(reflect.TypeOf(sentence).String())
		//reflect.TypeOf(sentence).Key() != TSVisitor.T
		//No vamos a visitar los nodos que vienen vacios NLContext
		//if reflect.TypeOf(sentence).String() != "*TSVisitor.SentNLContext" {
		nSentence := sentence.Accept(T).(TSStructs.TSExpressioner)
		lsentences.Sentences = append(lsentences.Sentences, nSentence)
	}
	return lsentences
}

// sents : expr NL #SentExpr
func (T TSwiftVisitor) VisitSentExpr(ctx *TSVisitor.SentExprContext) interface{} {
	return ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
}

//func (T TSwiftVisitor) VisitSentNL(ctx *TSVisitor.SentNLContext) interface{} {
//	return nil
//}

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

// expr : VBOOL
func (T TSwiftVisitor) VisitEVBOOL(ctx *TSVisitor.EVBOOLContext) interface{} {
	//todo: errores de conversion de datos
	bvalue, _ := strconv.ParseBool(ctx.VBOOL().GetText())
	return NTExpression.NewIBooleanCreation(ctx.VBOOL().GetSymbol().GetLine(),
		ctx.VBOOL().GetSymbol().GetColumn(),
		bvalue)
}

// expr -> ID
func (T TSwiftVisitor) VisitEID(ctx *TSVisitor.EIDContext) interface{} {
	//return NTExpression.NewIBox(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), ctx.ID().GetText())
	return NTExpression.NewIBoxGetting(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), ctx.ID().GetText())
}

func (T TSwiftVisitor) VisitEParent(ctx *TSVisitor.EParentContext) interface{} {
	return ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
}

/*
************ DECLARATIONS ******************
 */
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

func (T TSwiftVisitor) VisitSDecl(ctx *TSVisitor.SDeclContext) interface{} {
	return ctx.Declar().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitSDeclAsig(ctx *TSVisitor.SDeclAsigContext) interface{} {
	variable := ctx.Declar().Accept(T).(TSStructs.TSExpressioner)
	content := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIAssignation(ctx.GetOp().GetLine(), 0, variable, content)
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

/**
************** ESTRUCTURAS CONDICIONALES IF ****************
 */

func (T TSwiftVisitor) VisitSIf(ctx *TSVisitor.SIfContext) interface{} {
	return ctx.If_().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitRIf(ctx *TSVisitor.RIfContext) interface{} {
	condition := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	block := ctx.Block().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewITSIf(ctx.IF().GetSymbol().GetLine(), 0, condition, block)
}

func (T TSwiftVisitor) VisitRIfElse(ctx *TSVisitor.RIfElseContext) interface{} {
	condition := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	blockIf := ctx.Block(0).Accept(T).(TSStructs.TSExpressioner)
	blockElse := ctx.Block(1).Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewITSIfElse(ctx.IF().GetSymbol().GetLine(), 0, condition, blockIf, blockElse)
}

func (T TSwiftVisitor) VisitRIfEIf(ctx *TSVisitor.RIfEIfContext) interface{} {
	condition := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	blockIf := ctx.Block().Accept(T).(TSStructs.TSExpressioner)
	blockElse := ctx.If_().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewITSIfElseIf(ctx.IF().GetSymbol().GetLine(), 0, condition, blockIf, blockElse)

}

/**
************** ESTRUCTURAS CONDICIONALES SWITCH ****************
 */
func (T TSwiftVisitor) VisitSSwitch(ctx *TSVisitor.SSwitchContext) interface{} {
	return ctx.Switch_().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitSwitch(ctx *TSVisitor.SwitchContext) interface{} {

	expr := ctx.GetInput().Accept(T).(TSStructs.TSExpressioner)

	var cdefault TSStructs.TSExpressioner

	if ctx.Default_() != nil {
		cdefault = ctx.Default_().Accept(T).(TSStructs.TSExpressioner)
	}

	nswitch := NTExpression.NewISwitch(ctx.SWITCH().GetSymbol().GetLine(), ctx.SWITCH().GetSymbol().GetColumn(), expr, cdefault)

	d := ctx.AllCASE()
	for i := 0; i < len(d); i++ {
		gline := ctx.AllCASE()[i].GetSymbol().GetLine()
		gcolumn := ctx.AllCASE()[i].GetSymbol().GetLine()
		cexpr := ctx.Expr(i + 1).Accept(T).(TSStructs.TSExpressioner) //se ignora la expresion a comparar
		csents := ctx.Expr(i + 1).Accept(T).(TSStructs.TSExpressioner)

		ccase := NTExpression.NewICase(gline, gcolumn, cexpr, csents)

		nswitch.Icases = append(nswitch.Icases, *ccase)
	}

	return nswitch

}

func (T TSwiftVisitor) VisitDefault(ctx *TSVisitor.DefaultContext) interface{} {
	sentences := ctx.Lsents().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIDefault(ctx.DEFAULT().GetSymbol().GetLine(), ctx.DEFAULT().GetSymbol().GetColumn(), sentences)
}

/**
************** FUNCIONES DE TSWIFT ****************
 */

func (T TSwiftVisitor) VisitSPrint(ctx *TSVisitor.SPrintContext) interface{} {
	return ctx.Print_().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitPrint(ctx *TSVisitor.PrintContext) interface{} {
	expressions := ctx.AllExpr()
	sPrint := NTExpression.NewIFPrint(ctx.PRINT().GetSymbol().GetLine(), ctx.PRINT().GetSymbol().GetColumn())
	for _, expr := range expressions {
		nExpr := expr.Accept(T).(TSStructs.TSExpressioner)
		sPrint.Exprs = append(sPrint.Exprs, nExpr)
	}
	return sPrint
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
