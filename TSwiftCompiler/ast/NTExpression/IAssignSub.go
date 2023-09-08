package NTExpression

import (
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

func (I IAssignSub) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	var box *TSStructs.TSValue = I.box.Interpret(ctx)
	content := I.content.Interpret(ctx)

	//error, no era una variable
	if !box.IsBox {
		ctx.AddException(fmt.Sprintf("AsignaciónResta: esta no es una variable valida {%s}", box.ToString()), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	if box.IsConstant {
		ctx.AddException(fmt.Sprintf("AsignaciónResta: este valor no se puede modificar %s", box.BoxTag), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	boxType := box.TSType

	//error, solo valores permitidos
	if boxType != TSStructs.FLOAT && boxType != TSStructs.INTEGER {
		ctx.AddException(fmt.Sprintf("AsignaciónResta: tipo invalido para operación  {%s}", boxType), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	//TODO: error no se puede asignar nil
	if content.IsBox && content.IsNil {
		ctx.AddException(fmt.Sprintf("AsignaciónResta: la variable no ha sido inicializada  {%s}", content.ToString()), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	//si el contenido es nulo
	if content.IsNil {
		return TSStructs.NewTNil()
	}

	contentType := content.TSType

	//conversion de INT a FLOAT si se asigna a un FLOAT
	if boxType == TSStructs.FLOAT && contentType == TSStructs.INTEGER {
		box.IsNil = false
		box.Fvalue -= float64(content.Ivalue)
		return TSStructs.NewTNil()
	}

	if boxType != contentType {
		msg := fmt.Sprintf("AsignaciónResta: valores incompatibles {%s} != {%s}", boxType, contentType)
		ctx.AddException(msg, I.Line, I.Position)
		return TSStructs.NewTNil()

	}

	box.IsNil = false

	switch boxType {
	case TSStructs.INTEGER:
		box.Ivalue -= content.Ivalue
	case TSStructs.FLOAT:
		box.Fvalue -= content.Fvalue

	}

	return TSStructs.NewTNil()
}
