package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
)

type IOpAdd struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIAdd(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IOpAdd {
	return &IOpAdd{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IOpAdd) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	addend1 := I.op1.Interpret(ctx)
	addend2 := I.op2.Interpret(ctx)


	if addend1.IsNil || addend2.IsNil {
		ctx.AddException("Suma: No se puede realizar suma, tipos incompatibles.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch addend1.TSType {
	case TSStructs.INTEGER:
		switch addend2.TSType {
		case TSStructs.INTEGER:
			return TSStructs.NewTInt(addend1.Ivalue + addend2.Ivalue)
		case TSStructs.FLOAT:
			return TSStructs.NewTFloat(float64(addend1.Ivalue) + addend2.Fvalue)
		}
	case TSStructs.FLOAT:
		switch addend2.TSType {
		case TSStructs.INTEGER:
			return TSStructs.NewTFloat(addend1.Fvalue + float64(addend2.Ivalue))
		case TSStructs.FLOAT:
			return TSStructs.NewTFloat(addend1.Fvalue + addend2.Fvalue)
		}
	case TSStructs.STRING:
		if addend2.TSType == TSStructs.STRING {
			return TSStructs.NewTString(addend1.Svalue + addend2.Svalue)
		}
	}

	ctx.AddException("Suma: No se puede realizar suma, tipos incompatibles.", I.Line, I.Position)
	return TSStructs.NewTNil()
}
