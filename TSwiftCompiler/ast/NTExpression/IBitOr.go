package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type IBitOr struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIBitOr(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IBitOr {
	return &IBitOr{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IBitOr) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	op1 := I.op1.Interpret(ctx)
	op2 := I.op2.Interpret(ctx)

	if op1.IsNil || op2.IsNil {
		ctx.AddException("OperadorOr: no se puede operar, valor nulo", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	if op1.TSType != TExpression.BOOL || op2.TSType != TExpression.BOOL {
		ctx.AddException("OperadorOr: no se puede operar, valores distintos a bool.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	return TExpression.NewTBoolean(op1.Bvalue || op2.Bvalue)
}
