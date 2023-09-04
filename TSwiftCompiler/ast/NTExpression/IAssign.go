package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
	"fmt"
)

type IAssignation struct {
	TSStructs.TSExpression
	box     TSStructs.TSExpressioner
	content TSStructs.TSExpressioner
}

func NewIAssignation(Line int, Position int, box TSStructs.TSExpressioner, content TSStructs.TSExpressioner) *IAssignation {
	return &IAssignation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		box, content,
	}
}

func (I IAssignation) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	var box *TExpression.TSValue = I.box.Interpret(ctx)
	content := I.content.Interpret(ctx)

	//error, no era una variable
	if !box.IsBox {
		ctx.AddException(fmt.Sprintf("Asignación: esta no es una variable valida {%s}", box.ToString()), I.Line, I.Position)
		return TExpression.NewTNil()
	}

	//TODO: error no se puede asignar nil
	if content.IsBox && content.IsNil {
		ctx.AddException(fmt.Sprintf("Asignación: la variable no ha sido inicializada  {%s}", content.ToString()), I.Line, I.Position)
		return TExpression.NewTNil()
	}

	//si el contenido es nulo
	if content.IsNil {
		box.Svalue = ""
		box.Ivalue = 0
		box.Fvalue = 0
		box.Bvalue = false
		box.IsNil = true
		return TExpression.NewTNil()
	}

	boxType := box.TSType
	contentType := content.TSType

	//conversion de INT a FLOAT si se asigna a un FLOAT
	if boxType == TExpression.FLOAT && contentType == TExpression.INTEGER {
		box.IsNil = false
		box.Fvalue = float64(content.Ivalue)
		return TExpression.NewTNil()
	}

	if boxType != contentType {
		msg := fmt.Sprintf("Asignación: valores incompatibles {%s} != {%s}", boxType, contentType)
		ctx.AddException(msg, I.Line, I.Position)
		return TExpression.NewTNil()

	}

	box.IsNil = false

	switch boxType {
	case TExpression.STRING:
		box.Svalue = content.Svalue
	case TExpression.INTEGER:
		box.Ivalue = content.Ivalue
	case TExpression.FLOAT:
		box.Fvalue = content.Fvalue
	case TExpression.BOOL:
		box.Bvalue = content.Bvalue
	case TExpression.CHARACTER:
		//TODO: VALIDAR SI EL CONTENIDO DE LA CADENA ES MAYOR A UN CARACTER
		box.Svalue = content.Svalue
	}

	return TExpression.NewTNil()
}

func SetNil(box *TExpression.TSValue) *TExpression.TSValue {
	box.Svalue = ""
	box.Ivalue = 0
	box.Fvalue = 0
	box.Bvalue = false
	box.IsNil = true
	return box
}
