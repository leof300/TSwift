package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type IAdd struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIAdd(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IAdd {
	return &IAdd{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IAdd) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	addend1 := I.op1.Interpret(ctx)
	addend2 := I.op2.Interpret(ctx)

	//ctx.Console += fmt.Sprint("Suma: " + fmt.Sprint(addend1) + "+" + fmt.Sprint(addend2) + "\n")
	ctx.Console += "SUMA: {" + addend1.ToString() + "} + {" + addend2.ToString() + "}\n"

	if addend1.IsNil || addend2.IsNil {
		ctx.AddException("No se puede realizar suma, tipos incompatibles.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	switch addend1.TSType {
	case TExpression.INTEGER:
		switch addend2.TSType {
		case TExpression.INTEGER:
			return TExpression.NewTInt(addend1.Ivalue + addend2.Ivalue)
		case TExpression.FLOAT:
			return TExpression.NewTFloat(float64(addend1.Ivalue) + addend2.Fvalue)
		}
	case TExpression.FLOAT:
		switch addend2.TSType {
		case TExpression.INTEGER:
			return TExpression.NewTFloat(addend1.Fvalue + float64(addend2.Ivalue))
		case TExpression.FLOAT:
			return TExpression.NewTFloat(addend1.Fvalue + addend2.Fvalue)
		}
	case TExpression.STRING:
		if addend2.TSType == TExpression.STRING {
			return TExpression.NewTString(addend1.Svalue + addend2.Svalue)
		}
	}

	ctx.AddException("No se puede realizar suma, tipos incompatibles.", I.Line, I.Position)
	return TExpression.NewTNil()
}
