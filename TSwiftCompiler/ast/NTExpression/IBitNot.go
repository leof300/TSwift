package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type IBitNot struct {
	TSStructs.TSExpression
	op TSStructs.TSExpressioner
}

func NewIBitNot(Line int, Position int, op TSStructs.TSExpressioner) *IBitNot {
	return &IBitNot{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op,
	}
}

func (I IBitNot) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	op := I.op.Interpret(ctx)

	if op.IsNil {
		ctx.AddException("OperadorNot: no se puede operar, valor nulo", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	if op.TSType != TExpression.BOOL {
		ctx.AddException("OperadorNot: no se puede operar, valor distinto a bool.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	return TExpression.NewTBoolean(!op.Bvalue)
}
