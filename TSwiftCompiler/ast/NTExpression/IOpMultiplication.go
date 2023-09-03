package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type IOpMultiplication struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIMultiplication(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IOpMultiplication {
	return &IOpMultiplication{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IOpMultiplication) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	factor1 := I.op1.Interpret(ctx)
	factor2 := I.op2.Interpret(ctx)

	//ctx.Console += fmt.Sprint("Suma: " + fmt.Sprint(addend1) + "+" + fmt.Sprint(addend2) + "\n")
	ctx.Log += "MULTX: {" + factor1.ToString() + "} + {" + factor2.ToString() + "}\n"

	if factor1.IsNil || factor2.IsNil {
		ctx.AddException("No se puede realizar multiplicacion, tipos incompatibles.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	switch factor1.TSType {
	case TExpression.INTEGER:
		switch factor2.TSType {
		case TExpression.INTEGER:
			return TExpression.NewTInt(factor1.Ivalue * factor2.Ivalue)
		case TExpression.FLOAT:
			return TExpression.NewTFloat(float64(factor1.Ivalue) * factor2.Fvalue)
		}
	case TExpression.FLOAT:
		switch factor2.TSType {
		case TExpression.INTEGER:
			return TExpression.NewTFloat(factor1.Fvalue * float64(factor2.Ivalue))
		case TExpression.FLOAT:
			return TExpression.NewTFloat(factor1.Fvalue * factor2.Fvalue)
		}
	}

	ctx.AddException("No se puede realizar multiplicacion, tipos incompatibles.", I.Line, I.Position)
	return TExpression.NewTNil()
}
