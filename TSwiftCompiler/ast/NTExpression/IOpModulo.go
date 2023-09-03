package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type IOpModulo struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIModulo(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IOpModulo {
	return &IOpModulo{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IOpModulo) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	dividend := I.op1.Interpret(ctx)
	divisor := I.op2.Interpret(ctx)

	ctx.Log += "MULTX: {" + dividend.ToString() + "} + {" + divisor.ToString() + "}\n"

	//nil no se puede operar
	if dividend.IsNil || divisor.IsNil {
		ctx.AddException("No se puede realizar division, tipos Nil.", I.Line, I.Position)
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
			return TExpression.NewTInt(dividend.Ivalue % divisor.Ivalue)
		}
	}

	ctx.AddException("No se puede realizar division, tipos incompatibles.", I.Line, I.Position)
	return TExpression.NewTNil()
}
