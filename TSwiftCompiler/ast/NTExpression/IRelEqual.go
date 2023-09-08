package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
)

type IRelEqual struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIRelEqual(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IRelEqual {
	return &IRelEqual{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IRelEqual) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	op1 := I.op1.Interpret(ctx)
	op2 := I.op2.Interpret(ctx)

	err := TSStructs.RelationalValidationsBefore("RelMenorQ", op1, op2)

	if err != nil {
		ctx.AddException(err.Error(), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch op1.TSType {
	case TSStructs.INTEGER:
		return TSStructs.NewTBoolean(op1.Ivalue == op2.Ivalue)
	case TSStructs.FLOAT:
		return TSStructs.NewTBoolean(op1.Fvalue == op2.Fvalue)
	case TSStructs.BOOL:
		return TSStructs.NewTBoolean(op1.Bvalue == op2.Bvalue)
	case TSStructs.STRING:
		return TSStructs.NewTBoolean(op1.Svalue == op2.Svalue)
	case TSStructs.CHARACTER:
		return TSStructs.NewTBoolean(op1.Svalue == op2.Svalue)
	}

	ctx.AddException("Igualdad: no permitidos para esta operación.", I.Line, I.Position)
	return TSStructs.NewTNil()
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////////////////////////////////////////////////////////////////////////////////
type IRelNotEqual struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIRelNotEqual(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IRelNotEqual {
	return &IRelNotEqual{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IRelNotEqual) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	op1 := I.op1.Interpret(ctx)
	op2 := I.op2.Interpret(ctx)

	err := TSStructs.RelationalValidationsBefore("RelMenorQ", op1, op2)

	if err != nil {
		ctx.AddException(err.Error(), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch op1.TSType {
	case TSStructs.INTEGER:
		return TSStructs.NewTBoolean(op1.Ivalue != op2.Ivalue)
	case TSStructs.FLOAT:
		v := op1.Fvalue != op2.Fvalue
		return TSStructs.NewTBoolean(v)
	case TSStructs.BOOL:
		return TSStructs.NewTBoolean(op1.Bvalue != op2.Bvalue)
	case TSStructs.STRING:
		return TSStructs.NewTBoolean(op1.Svalue != op2.Svalue)
	case TSStructs.CHARACTER:
		return TSStructs.NewTBoolean(op1.Svalue != op2.Svalue)
	}

	ctx.AddException("Desigualdad: tipos no permitidos para esta operación.", I.Line, I.Position)
	return TSStructs.NewTNil()
}
