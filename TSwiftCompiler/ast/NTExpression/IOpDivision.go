package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
)

type IOpDivision struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIOpDivision(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IOpDivision {
	return &IOpDivision{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IOpDivision) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	dividend := I.op1.Interpret(ctx)
	divisor := I.op2.Interpret(ctx)


	//nil no se puede operar
	if dividend.IsNil || divisor.IsNil {
		ctx.AddException("No se puede realizar division, tipos incompatibles.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	//división dentro de cero
	if divisor.TSType == TSStructs.INTEGER || divisor.TSType == TSStructs.FLOAT {
		if divisor.Fvalue == 0 && divisor.Ivalue == 0 {
			ctx.AddException("No se puede realizar dividir dentro de cero.", I.Line, I.Position)
			return TSStructs.NewTNil()
		}
	}

	switch dividend.TSType {
	case TSStructs.INTEGER:
		switch divisor.TSType {
		case TSStructs.INTEGER:
			return TSStructs.NewTInt(dividend.Ivalue / divisor.Ivalue)
		case TSStructs.FLOAT:
			return TSStructs.NewTFloat(float64(dividend.Ivalue) / divisor.Fvalue)
		}
	case TSStructs.FLOAT:
		switch divisor.TSType {
		case TSStructs.INTEGER:
			return TSStructs.NewTFloat(dividend.Fvalue / float64(divisor.Ivalue))
		case TSStructs.FLOAT:
			return TSStructs.NewTFloat(dividend.Fvalue / divisor.Fvalue)
		}
	}

	ctx.AddException("No se puede realizar division, tipos incompatibles.", I.Line, I.Position)
	return TSStructs.NewTNil()
}
