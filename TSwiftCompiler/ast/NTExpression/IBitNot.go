package NTExpression

import (
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

func (I IBitNot) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	op := I.op.Interpret(ctx)

	if op.IsNil {
		ctx.AddException("OperadorNot: no se puede operar, valor nulo", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	if op.TSType != TSStructs.BOOL {
		ctx.AddException("OperadorNot: no se puede operar, valor distinto a bool.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	return TSStructs.NewTBoolean(!op.Bvalue)
}
