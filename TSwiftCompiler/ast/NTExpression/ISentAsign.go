package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
	"fmt"
)

type ISentAsign struct {
	TSStructs.TSExpression
	box     TSStructs.TSExpressioner
	content TSStructs.TSExpressioner
}

func NewISentAsign(Line int, Position int, box TSStructs.TSExpressioner, content TSStructs.TSExpressioner) *ISentAsign {
	return &ISentAsign{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		box, content,
	}
}

func (I ISentAsign) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	var box *TExpression.TSValue = I.box.Interpret(ctx)
	content := I.content.Interpret(ctx)

	if !box.IsBox {
		ctx.AddException(fmt.Sprint("No se puede asignar a %s", box.ToString()), I.Line, I.Position)
	}

	if content.IsBox && content.TSType == TExpression.UNDEFINED {
		ctx.AddException(fmt.Sprint("El valor que se desea asignar no es correcto %s", I.box), I.Line, I.Position)
	}

	content = TSStructs.Unbox(content)

	boxType := box.TSType

	if boxType != content.TSType {
		ctx.AddException(fmt.Sprint("El valor de tipo (%s) que se desea asignar a %s no corresponde al tipo %s", content.TSType, box.BoxTag, boxType), I.Line, I.Position)
	}

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
		box.Svalue = content.Svalue
	}

	//TODO cambiar el valor de IsNul al final
	box.IsNil = false

	return box
}
