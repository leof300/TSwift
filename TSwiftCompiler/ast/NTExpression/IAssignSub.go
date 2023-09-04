package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
	"fmt"
)

type IAssignSub struct {
	TSStructs.TSExpression
	box     TSStructs.TSExpressioner
	content TSStructs.TSExpressioner
}

func NewIAssignSub(Line int, Position int, box TSStructs.TSExpressioner, content TSStructs.TSExpressioner) *IAssignSub {
	return &IAssignSub{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		box, content,
	}
}

func (I IAssignSub) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	var box *TExpression.TSValue = I.box.Interpret(ctx)
	content := I.content.Interpret(ctx)

	//error, no era una variable
	if !box.IsBox {
		ctx.AddException(fmt.Sprintf("AsignaciónResta: esta no es una variable valida {%s}", box.ToString()), I.Line, I.Position)
		return TExpression.NewTNil()
	}

	boxType := box.TSType

	//error, solo valores permitidos
	if boxType != TExpression.FLOAT && boxType != TExpression.INTEGER {
		ctx.AddException(fmt.Sprintf("AsignaciónResta: tipo invalido para operación  {%s}", boxType), I.Line, I.Position)
		return TExpression.NewTNil()
	}

	//TODO: error no se puede asignar nil
	if content.IsBox && content.IsNil {
		ctx.AddException(fmt.Sprintf("AsignaciónResta: la variable no ha sido inicializada  {%s}", content.ToString()), I.Line, I.Position)
		return TExpression.NewTNil()
	}

	//si el contenido es nulo
	if content.IsNil {
		return TExpression.NewTNil()
	}

	contentType := content.TSType

	//conversion de INT a FLOAT si se asigna a un FLOAT
	if boxType == TExpression.FLOAT && contentType == TExpression.INTEGER {
		box.IsNil = false
		box.Fvalue -= float64(content.Ivalue)
		return TExpression.NewTNil()
	}

	if boxType != contentType {
		msg := fmt.Sprintf("AsignaciónResta: valores incompatibles {%s} != {%s}", boxType, contentType)
		ctx.AddException(msg, I.Line, I.Position)
		return TExpression.NewTNil()

	}

	box.IsNil = false

	switch boxType {
	case TExpression.INTEGER:
		box.Ivalue -= content.Ivalue
	case TExpression.FLOAT:
		box.Fvalue -= content.Fvalue

	}

	return TExpression.NewTNil()
}
