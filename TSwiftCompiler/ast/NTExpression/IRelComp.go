package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
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

func (I IOpNegation) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	operand := I.op.Interpret(ctx)

	//error, valor nulo
	if operand.IsNil {
		ctx.AddException("Negación: No se puede realizar negación, valor nulo.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	operandType := operand.TSType

	//error, solo valores permitidos
	if operandType != TExpression.FLOAT && operandType != TExpression.INTEGER {
		ctx.AddException(fmt.Sprintf("Negación: tipo invalido para operación  {%s}", operandType), I.Line, I.Position)
		return TExpression.NewTNil()
	}

	switch operandType {
	case TExpression.INTEGER:
		return TExpression.NewTInt(operand.Ivalue * -1)
	case TExpression.FLOAT:
		return TExpression.NewTFloat(operand.Fvalue * -1)
	}

	return TExpression.NewTNil()
}
