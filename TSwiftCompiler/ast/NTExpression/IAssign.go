package NTExpression

import (
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

func (I IAssignation) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	var box *TSStructs.TSValue = I.box.Interpret(ctx)
	content := I.content.Interpret(ctx)

	//error, no era una variable
	if !box.IsBox {
		ctx.AddException(fmt.Sprintf("Asignación: esta variable no existe %s", box.BoxTag), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	if box.IsConstant {
		ctx.AddException(fmt.Sprintf("Asignación: este valor no se puede modificar %s", box.BoxTag), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	//TODO: error no se puede asignar nil
	if content.IsBox && content.IsNil {
		ctx.AddException(fmt.Sprintf("Asignación: la variable no ha sido inicializada  {%s}", content.ToString()), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	//si el contenido es nulo
	if content.IsNil {
		box.Svalue = ""
		box.Ivalue = 0
		box.Fvalue = 0
		box.Bvalue = false
		box.IsNil = true
		return TSStructs.NewTNil()
	}

	boxType := box.TSType
	contentType := content.TSType

	//conversion de INT a FLOAT si se asigna a un FLOAT
	if boxType == TSStructs.FLOAT && contentType == TSStructs.INTEGER {
		box.IsNil = false
		box.Fvalue = float64(content.Ivalue)
		return TSStructs.NewTNil()
	}

	if boxType != contentType {
		msg := fmt.Sprintf("Asignación: valores incompatibles {%s} != {%s}", boxType, contentType)
		ctx.AddException(msg, I.Line, I.Position)
		return TSStructs.NewTNil()

	}

	box.IsNil = false

	switch boxType {
	case TSStructs.STRING:
		box.Svalue = content.Svalue
	case TSStructs.INTEGER:
		box.Ivalue = content.Ivalue
	case TSStructs.FLOAT:
		box.Fvalue = content.Fvalue
	case TSStructs.BOOL:
		box.Bvalue = content.Bvalue
	case TSStructs.CHARACTER:
		//TODO: VALIDAR SI EL CONTENIDO DE LA CADENA ES MAYOR A UN CARACTER
		box.Svalue = content.Svalue
	}

	if box.IsArray && content.IsArray {
		box.ArrayContent = *TSStructs.CopyTSArrayNew(content.ArrayContent)
	}

	return TSStructs.NewTNil()
}

func SetNil(box *TSStructs.TSValue) *TSStructs.TSValue {
	box.Svalue = ""
	box.Ivalue = 0
	box.Fvalue = 0
	box.Bvalue = false
	box.IsNil = true
	return box
}
