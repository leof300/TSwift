package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
)

type IOpSubtraction struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewISubtraction(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IOpSubtraction {
	return &IOpSubtraction{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IOpSubtraction) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	minuend := I.op1.Interpret(ctx)
	subtrahend := I.op2.Interpret(ctx)

	if minuend.IsNil || subtrahend.IsNil {
		ctx.AddException("No se puede realizar resta, tipos incompatibles.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch minuend.TSType {
	case TSStructs.INTEGER:
		switch subtrahend.TSType {
		case TSStructs.INTEGER:
			return TSStructs.NewTInt(minuend.Ivalue - subtrahend.Ivalue)
		case TSStructs.FLOAT:
			return TSStructs.NewTFloat(float64(minuend.Ivalue) - subtrahend.Fvalue)
		}
	case TSStructs.FLOAT:
		switch subtrahend.TSType {
		case TSStructs.INTEGER:
			return TSStructs.NewTFloat(minuend.Fvalue - float64(subtrahend.Ivalue))
		case TSStructs.FLOAT:
			return TSStructs.NewTFloat(minuend.Fvalue - subtrahend.Fvalue)
		}
	}

	ctx.AddException("No se puede realizar resta, tipos incompatibles.", I.Line, I.Position)
	return TSStructs.NewTNil()
}
