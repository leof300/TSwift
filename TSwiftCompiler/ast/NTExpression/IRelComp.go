package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
	"fmt"
)

type IOpNegation struct {
	TSStructs.TSExpression
	op TSStructs.TSExpressioner
}

func NewIOpNegation(Line int, Position int, op TSStructs.TSExpressioner) *IOpNegation {
	return &IOpNegation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op,
	}
}

func (I IOpNegation) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	operand := I.op.Interpret(ctx)

	//error, valor nulo
	if operand.IsNil {
		ctx.AddException("Negación: No se puede realizar negación, valor nulo.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	operandType := operand.TSType

	//error, solo valores permitidos
	if operandType != TSStructs.FLOAT && operandType != TSStructs.INTEGER {
		ctx.AddException(fmt.Sprintf("Negación: tipo invalido para operación  {%s}", operandType), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch operandType {
	case TSStructs.INTEGER:
		return TSStructs.NewTInt(operand.Ivalue * -1)
	case TSStructs.FLOAT:
		return TSStructs.NewTFloat(operand.Fvalue * -1)
	}

	return TSStructs.NewTNil()
}
