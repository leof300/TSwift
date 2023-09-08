package interpreter

import (
	TSExceptions "TSwiftCompiler/ast/Exceptions"
	"TSwiftCompiler/ast/NTExpression"
	"TSwiftCompiler/ast/TSStructs"
	"TSwiftCompiler/visitor"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
	"strings"
)

type TSwiftVisitor struct {
	antlr.ParseTreeVisitor
	Start TSStructs.TSExpressioner
}

/*-----------------------------------------------------------------------------------------------------*/
func (T TSwiftVisitor) VisitBlock(ctx *TSVisitor.BlockContext) interface{} {
	sentences := ctx.Lsents().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewITSBlock(ctx.GetTerm().GetLine(), 0, sentences)
}

// start : sents+ EOF
func (T TSwiftVisitor) VisitLsents(ctx *TSVisitor.LsentsContext) interface{} {
	lsentences := NTExpression.NewILSentences()
	sentences := ctx.AllSents()
	for _, sentence := range sentences {

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

//////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////// START////////////////////////////////////////////////////////////////////

func (T TSwiftVisitor) VisitStart(ctx *TSVisitor.StartContext) interface{} {
	T.Start = ctx.Lins().Accept(T).(TSStructs.TSExpressioner)
	return T.Start
}

func (T TSwiftVisitor) VisitLins(ctx *TSVisitor.LinsContext) interface{} {
	lsentences := NTExpression.NewILSentences()
	sentences := ctx.AllIns()
	for _, sentence := range sentences {
		nSentence := sentence.Accept(T).(TSStructs.TSExpressioner)
		lsentences.Sentences = append(lsentences.Sentences, nSentence)
	}
	return lsentences
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%      VALUES       %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
*/

func (T TSwiftVisitor) VisitEParent(ctx *TSVisitor.EParentContext) interface{} {
	return ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
}

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
		strings.Replace(ctx.VSTRING().GetText(), `"`, ``, -1))
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

func (T TSwiftVisitor) VisitENIL(ctx *TSVisitor.ENILContext) interface{} {
	return NTExpression.NewINilCreation(ctx.NIL().GetSymbol().GetLine(), ctx.NIL().GetSymbol().GetColumn())
}

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%      DECLARATIONS      %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
*/

func (T TSwiftVisitor) VisitIDecl(ctx *TSVisitor.IDeclContext) interface{} {
	return ctx.Declar().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitSDecl(ctx *TSVisitor.SDeclContext) interface{} {
	return ctx.Declar().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitSDType(ctx *TSVisitor.SDTypeContext) interface{} {
	initValue := ctx.Tstypes().Accept(T).(TSStructs.TSExpressioner)
	isArray := false

	if ctx.GetArr() != nil && ctx.GetArr().GetText() != "" {
		isArray = true
	}

	return NTExpression.NewIBoxCreation(ctx.VAR().GetSymbol().GetLine(), 0, ctx.ID().GetText(), initValue, isArray)
}

func (T TSwiftVisitor) VisitSDecAssign(ctx *TSVisitor.SDecAssignContext) interface{} {
	line := ctx.VAR().GetSymbol().GetLine()
	position := ctx.VAR().GetSymbol().GetColumn()

	initValue := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)

	return NTExpression.NewIBoxCreation(line, position, ctx.ID().GetText(), initValue, false)

}

func (T TSwiftVisitor) VisitSDecTAssign(ctx *TSVisitor.SDecTAssignContext) interface{} {
	tstype := ctx.Tstypes().Accept(T).(TSStructs.TSExpressioner)
	initValue := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)

	isArray := false
	if ctx.GetArr() != nil && ctx.GetArr().GetText() != "" {
		isArray = true
	}

	variable := NTExpression.NewIBoxCreation(ctx.VAR().GetSymbol().GetLine(), ctx.VAR().GetSymbol().GetColumn(), ctx.ID().GetText(), tstype, isArray)
	return NTExpression.NewIAssignation(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), variable, initValue)
}

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%        TYPES         %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
*/

func (T TSwiftVisitor) VisitTstypes(ctx *TSVisitor.TstypesContext) interface{} {
	etype := NTExpression.NewITSType(0, 0, TSStructs.UNDEFINED)

	if ctx.STRING() != nil {
		etype = NTExpression.NewITSType(ctx.STRING().GetSymbol().GetLine(), ctx.STRING().GetSymbol().GetColumn(), TSStructs.STRING)

	}
	if ctx.INT() != nil {
		etype = NTExpression.NewITSType(ctx.INT().GetSymbol().GetLine(), ctx.INT().GetSymbol().GetColumn(), TSStructs.INTEGER)

	}
	if ctx.FLOAT() != nil {
		etype = NTExpression.NewITSType(ctx.FLOAT().GetSymbol().GetLine(), ctx.FLOAT().GetSymbol().GetColumn(), TSStructs.FLOAT)

	}
	if ctx.BOOL() != nil {
		etype = NTExpression.NewITSType(ctx.BOOL().GetSymbol().GetLine(), ctx.BOOL().GetSymbol().GetColumn(), TSStructs.BOOL)

	}
	if ctx.CHARACTER() != nil {
		etype = NTExpression.NewITSType(ctx.CHARACTER().GetSymbol().GetLine(), ctx.CHARACTER().GetSymbol().GetColumn(), TSStructs.CHARACTER)

	}

	return etype

}

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%      CONSTANTS      %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
*/

func (T TSwiftVisitor) VisitICons(ctx *TSVisitor.IConsContext) interface{} {
	return ctx.Decons().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitSCons(ctx *TSVisitor.SConsContext) interface{} {

	return ctx.Decons().Accept(T).(TSStructs.TSExpressioner)

}

func (T TSwiftVisitor) VisitSConsAss(ctx *TSVisitor.SConsAssContext) interface{} {
	line := ctx.LET().GetSymbol().GetLine()
	position := ctx.LET().GetSymbol().GetColumn()

	initValue := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)

	boxCreation := NTExpression.NewIBoxCreation(line, position, ctx.ID().GetText(), initValue, false)
	boxGetting := NTExpression.NewIBoxGetting(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), ctx.ID().GetText())

	return NTExpression.NewIConstCreation(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), boxGetting, boxCreation)
}

func (T TSwiftVisitor) VisitSConsTypeAss(ctx *TSVisitor.SConsTypeAssContext) interface{} {
	tstype := ctx.Tstypes().Accept(T).(TSStructs.TSExpressioner)
	initValue := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)

	variable := NTExpression.NewIBoxCreation(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), ctx.ID().GetText(), tstype, false)
	boxCreation := NTExpression.NewIAssignation(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), variable, initValue)
	boxGetting := NTExpression.NewIBoxGetting(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), ctx.ID().GetText())
	return NTExpression.NewIConstCreation(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), boxGetting, boxCreation)
}

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%      ASIGNACION     %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
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

/***********************************************************************************************
4.5 Operadores Aritmeticos
************************************************************************************************/
/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%  OPERACIONES ARITMETICAS %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
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

/***********************************************************************************************
4.6 Operaciones de Comparación
************************************************************************************************/
/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%  OPERACIONES RELACIONALES %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
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

/***********************************************************************************************
4.8 SENTENCIAS DE CONTROL DE FLUJO
************************************************************************************************/
/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%  Estructuras Condicionales:    IF    %%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
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

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%  Estructuras Condicionales:    SWITCH    %%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
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

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%  EEstructuras Condicionales:   WHILE   %%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
*/

func (T TSwiftVisitor) VisitSWhile(ctx *TSVisitor.SWhileContext) interface{} {
	return ctx.While().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitWhile(ctx *TSVisitor.WhileContext) interface{} {
	condition := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	block := ctx.Block().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewISWhile(ctx.WHILE().GetSymbol().GetLine(), 0, condition, block)
}

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%  Estructuras Condicionales:  GUARD    %%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
*/
func (T TSwiftVisitor) VisitSGuard(ctx *TSVisitor.SGuardContext) interface{} {
	return ctx.Guard().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitGuard(ctx *TSVisitor.GuardContext) interface{} {
	expr := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	block := ctx.Block().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewISGuard(ctx.GUARD().GetSymbol().GetLine(), ctx.GUARD().GetSymbol().GetColumn(), expr, block)
}

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%  Estructuras Condicionales:  FOR   %%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
*/

func (T TSwiftVisitor) VisitSFor(ctx *TSVisitor.SForContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (T TSwiftVisitor) VisitFor(ctx *TSVisitor.ForContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (T TSwiftVisitor) VisitRango(ctx *TSVisitor.RangoContext) interface{} {
	//TODO implement me
	panic("implement me")
}

/***********************************************************************************************
4.9 SENTENCIAS DE TRANSFERENCIA
************************************************************************************************/

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%  Estructuras TRANSICION:  CONTINUE, BREAK, RETURN   %%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
*/

func (T TSwiftVisitor) VisitSentTrans(ctx *TSVisitor.SentTransContext) interface{} {
	return ctx.Strans().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitStrans(ctx *TSVisitor.StransContext) interface{} {
	if ctx.CONTINUE() != nil {
		return NTExpression.NewIContinue(ctx.CONTINUE().GetSymbol().GetLine(), ctx.CONTINUE().GetSymbol().GetColumn())
	}
	if ctx.BREAK() != nil {
		return NTExpression.NewIBreak(ctx.BREAK().GetSymbol().GetLine(), ctx.BREAK().GetSymbol().GetColumn())
	}

	a := ctx.RETURN()
	b := ctx.Expr()
	if a != nil && b != nil {
		expr := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
		return NTExpression.NewIReturn(ctx.RETURN().GetSymbol().GetLine(), ctx.RETURN().GetSymbol().GetColumn(), expr)
	}

	return NTExpression.NewINoReturn(ctx.RETURN().GetSymbol().GetLine(), ctx.RETURN().GetSymbol().GetColumn())
}

/***********************************************************************************************
4.9 FUNCIONES EMBEBIDAS
************************************************************************************************/

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%       PRINT             %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
*/
func (T TSwiftVisitor) VisitSPrint(ctx *TSVisitor.SPrintContext) interface{} {
	return ctx.Tsprint().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitIPrint(ctx *TSVisitor.IPrintContext) interface{} {
	return ctx.Tsprint().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitTsprint(ctx *TSVisitor.TsprintContext) interface{} {
	expressions := ctx.AllExpr()
	sPrint := NTExpression.NewIFPrint(ctx.PRINT().GetSymbol().GetLine(), ctx.PRINT().GetSymbol().GetColumn())
	for _, expr := range expressions {
		nExpr := expr.Accept(T).(TSStructs.TSExpressioner)
		sPrint.Exprs = append(sPrint.Exprs, nExpr)
	}
	return sPrint
}

/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%       CONVERT           %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
*/
func (T TSwiftVisitor) VisitETSFunctions(ctx *TSVisitor.ETSFunctionsContext) interface{} {
	return ctx.Tsfunctions().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitConvertString(ctx *TSVisitor.ConvertStringContext) interface{} {
	expr := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIFString(ctx.STRING().GetSymbol().GetLine(), ctx.STRING().GetSymbol().GetColumn(), expr)
}

func (T TSwiftVisitor) VisitConvertInt(ctx *TSVisitor.ConvertIntContext) interface{} {
	expr := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIFInt(ctx.INT().GetSymbol().GetLine(), ctx.INT().GetSymbol().GetColumn(), expr)

}

func (T TSwiftVisitor) VisitConvertFloat(ctx *TSVisitor.ConvertFloatContext) interface{} {
	expr := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	return NTExpression.NewIFFloat(ctx.FLOAT().GetSymbol().GetLine(), ctx.FLOAT().GetSymbol().GetColumn(), expr)

}

/***********************************************************************************************
7.0 FUNCIONES
************************************************************************************************/
/*
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%       FUNCIONES         %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
*/
func (T TSwiftVisitor) VisitIFunc(ctx *TSVisitor.IFuncContext) interface{} {
	return ctx.Functions().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitFunctions(ctx *TSVisitor.FunctionsContext) interface{} {

	line := ctx.FUNC().GetSymbol().GetLine()
	position := ctx.FUNC().GetSymbol().GetColumn()

	id := ctx.ID().GetText()
	fbody := ctx.Lsents().Accept(T).(TSStructs.TSExpressioner)

	var ftype TSStructs.TSExpressioner

	if ctx.Tstypes() != nil {
		ftype = ctx.Tstypes().Accept(T).(TSStructs.TSExpressioner)
	}

	function := NTExpression.NewIFunction(line, position, id, ftype, fbody)

	for _, p := range ctx.AllParameter() {
		fparameter := p.Accept(T).(TSStructs.TSExpressioner)
		function.FParameters = append(function.FParameters, fparameter)
	}

	return function
}

func (T TSwiftVisitor) VisitParameter(ctx *TSVisitor.ParameterContext) interface{} {
	alias := ""
	id := ""
	line := 0
	position := 0

	if len(ctx.AllID()) > 1 {
		alias = ctx.ID(0).GetText()
		id = ctx.ID(1).GetText()
		line = ctx.ID(1).GetSymbol().GetLine()
		position = ctx.ID(1).GetSymbol().GetLine()

	} else {
		id = ctx.ID(0).GetText()
		alias = ctx.ID(0).GetText()
		line = ctx.ID(0).GetSymbol().GetLine()
		position = ctx.ID(0).GetSymbol().GetLine()
	}
	ftype := ctx.Tstypes().Accept(T).(TSStructs.TSExpressioner)

	return NTExpression.NewIFParameter(line, position, alias, id, ftype, ctx.INOUT() != nil)

}

func (T TSwiftVisitor) VisitSCallFunction(ctx *TSVisitor.SCallFunctionContext) interface{} {
	return ctx.CallFunction().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitEFunction(ctx *TSVisitor.EFunctionContext) interface{} {
	return ctx.CallFunction().Accept(T).(TSStructs.TSExpressioner)
}

func (T TSwiftVisitor) VisitCallFunction(ctx *TSVisitor.CallFunctionContext) interface{} {
	fcall := NTExpression.NewIFCallFunction(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), ctx.ID().GetText())
	for _, p := range ctx.AllCallParameter() {
		parameter := p.Accept(T).(TSStructs.TSExpressioner)
		fcall.Parameters = append(fcall.Parameters, parameter)
	}

	return fcall
}

func (T TSwiftVisitor) VisitCallParameter(ctx *TSVisitor.CallParameterContext) interface{} {
	expr := ctx.Expr().Accept(T).(TSStructs.TSExpressioner)
	isReference := false
	if ctx.GetOp() != nil {
		isReference = true
	}
	alias := ""
	if ctx.ID() != nil {
		alias = ctx.ID().GetText()
	}

	return NTExpression.NewIFCallParameter(0, 0, alias, expr, isReference)
}

// ////////////////////////////////////////////////////////////////////////////////////// ////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////// ////////////////////////////////////////////////////////////////////////////////////
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

// ////////////////////////////////////////////////////////////////////////////////////// ////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////// ////////////////////////////////////////////////////////////////////////////////////
