package interpreter

import (
	"TSwiftCompiler/ast"
	TSExceptions "TSwiftCompiler/ast/Exceptions"
	"TSwiftCompiler/ast/NTExpression"
	"TSwiftCompiler/visitor"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"reflect"
	"strconv"
)

type TSwiftVisitor struct {
	antlr.ParseTreeVisitor
	Start ast.TSExpressioner
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
		//reflect.TypeOf(sentence).Name() != TSVisitor.T
		//No vamos a visitar los nodos que vienen vacios NLContext
		if reflect.TypeOf(sentence).String() != "*TSVisitor.SentNLContext" {
			nSentence := sentence.Accept(T).(ast.TSExpressioner)
			lsentences.Sentences = append(lsentences.Sentences, nSentence)
		}
	}
	return lsentences
}

// sents : expr NL #SentExpr
func (T TSwiftVisitor) VisitSentExpr(ctx *TSVisitor.SentExprContext) interface{} {
	return ctx.Expr().Accept(T).(ast.TSExpressioner)
}

func (T TSwiftVisitor) VisitSentNL(ctx *TSVisitor.SentNLContext) interface{} {
	return nil
}

// //////////////////////////////// START////////////////////////////////////////////////////////////////////

func (T TSwiftVisitor) VisitStart(ctx *TSVisitor.StartContext) interface{} {
	T.Start = ctx.Lsents().Accept(T).(ast.TSExpressioner)
	return T.Start
}

/*********************************************************************************************************************/

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
	return NTExpression.NewIBox(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), ctx.ID().GetText())
}

func (T TSwiftVisitor) VisitEParent(ctx *TSVisitor.EParentContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (T TSwiftVisitor) VisitEMulDiv(ctx *TSVisitor.EMulDivContext) interface{} {
	//TODO implement me
	var _ = TSVisitor.TSParser_rulesParserVSTRING //accedemos a las constantes de los lexers
	panic("implement me")
}

func (T TSwiftVisitor) VisitEAssign(ctx *TSVisitor.EAssignContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (T TSwiftVisitor) VisitEAddSub(ctx *TSVisitor.EAddSubContext) interface{} {
	//TODO implement me
	e1 := ctx.Expr(0).Accept(T).(ast.TSExpressioner)
	return e1
}

func (T TSwiftVisitor) VisitEModule(ctx *TSVisitor.EModuleContext) interface{} {
	//TODO implement me
	panic("implement me")
}

/*
constructor equivalente a:

public static class TSwiftVisitor extends BaseTSParser_rulesVisitor<T> {
*/
func NewTSwiftVisitor() TSVisitor.TSParser_rulesVisitor {
	return &TSwiftVisitor{ParseTreeVisitor: &TSVisitor.BaseTSParser_rulesVisitor{}}
}

func (v TSwiftVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		//log.Fatal(val.GetText())
		return TSExceptions.NewTSSemanticE(val.GetText(), 0, 0)
	default:
		//comprobamos el el nodo sea valido
		//TODO
		nodoIntepreter, ok := tree.Accept(v).(ast.TSExpressioner)
		if ok {
			return nodoIntepreter
		}
		fmt.Print("nodo desconocido")
		return TSExceptions.NewTSSemanticE("ERROR NODO NO VALIDO", 0, 0)
	}
}
