package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type IOpDivision struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIDivision(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IOpDivision {
	return &IOpDivision{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IOpDivision) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	dividend := I.op1.Interpret(ctx)
	divisor := I.op2.Interpret(ctx)

	ctx.Log += "MULTX: {" + dividend.ToString() + "} + {" + divisor.ToString() + "}\n"

	//nil no se puede operar
	if dividend.IsNil || divisor.IsNil {
		ctx.AddException("No se puede realizar division, tipos incompatibles.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	//división dentro de cero
	if divisor.TSType == TExpression.INTEGER || divisor.TSType == TExpression.FLOAT {
		if divisor.Fvalue == 0 && divisor.Ivalue == 0 {
			ctx.AddException("No se puede realizar dividir dentro de cero.", I.Line, I.Position)
			return TExpression.NewTNil()
		}
	}

	switch dividend.TSType {
	case TExpression.INTEGER:
		switch divisor.TSType {
		case TExpression.INTEGER:
			return TExpression.NewTInt(dividend.Ivalue / divisor.Ivalue)
		case TExpression.FLOAT:
			return TExpression.NewTFloat(float64(dividend.Ivalue) / divisor.Fvalue)
		}
	case TExpression.FLOAT:
		switch divisor.TSType {
		case TExpression.INTEGER:
			return TExpression.NewTFloat(dividend.Fvalue / float64(divisor.Ivalue))
		case TExpression.FLOAT:
			return TExpression.NewTFloat(dividend.Fvalue / divisor.Fvalue)
		}
	}

	ctx.AddException("No se puede realizar division, tipos incompatibles.", I.Line, I.Position)
	return TExpression.NewTNil()
}
