package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
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

func (I IOpSubtraction) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	minuend := I.op1.Interpret(ctx)
	subtrahend := I.op2.Interpret(ctx)

	//ctx.Console += fmt.Sprint("Suma: " + fmt.Sprint(addend1) + "+" + fmt.Sprint(addend2) + "\n")
	ctx.Log += "RESTA: {" + minuend.ToString() + "} + {" + subtrahend.ToString() + "}\n"

	if minuend.IsNil || subtrahend.IsNil {
		ctx.AddException("No se puede realizar resta, tipos incompatibles.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	switch minuend.TSType {
	case TExpression.INTEGER:
		switch subtrahend.TSType {
		case TExpression.INTEGER:
			return TExpression.NewTInt(minuend.Ivalue - subtrahend.Ivalue)
		case TExpression.FLOAT:
			return TExpression.NewTFloat(float64(minuend.Ivalue) - subtrahend.Fvalue)
		}
	case TExpression.FLOAT:
		switch subtrahend.TSType {
		case TExpression.INTEGER:
			return TExpression.NewTFloat(minuend.Fvalue - float64(subtrahend.Ivalue))
		case TExpression.FLOAT:
			return TExpression.NewTFloat(minuend.Fvalue - subtrahend.Fvalue)
		}
	}

	ctx.AddException("No se puede realizar resta, tipos incompatibles.", I.Line, I.Position)
	return TExpression.NewTNil()
}
