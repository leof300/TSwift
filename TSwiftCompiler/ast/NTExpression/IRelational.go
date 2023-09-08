package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
)

type IRelLess struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIRelLess(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IRelLess {
	return &IRelLess{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IRelLess) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	op1 := I.op1.Interpret(ctx)
	op2 := I.op2.Interpret(ctx)

	err := TSStructs.RelationalValidationsBefore("RelMenorQ", op1, op2)

	if err != nil {
		ctx.AddException(err.Error(), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch op1.TSType {
	case TSStructs.INTEGER:
		return TSStructs.NewTBoolean(op1.Ivalue < op2.Ivalue)
	case TSStructs.FLOAT:
		return TSStructs.NewTBoolean(op1.Fvalue < op2.Fvalue)
	case TSStructs.STRING:
		return TSStructs.NewTBoolean(op1.Svalue < op2.Svalue)
	case TSStructs.CHARACTER:
		return TSStructs.NewTBoolean(op1.Svalue < op2.Svalue)
	}

	ctx.AddException("RelMenorQ: tipos no permitidos para esta operación.", I.Line, I.Position)
	return TSStructs.NewTNil()
}

/**********************************************************************************/
type IRelGreater struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIRelGreater(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IRelGreater {
	return &IRelGreater{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IRelGreater) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	op1 := I.op1.Interpret(ctx)
	op2 := I.op2.Interpret(ctx)

	err := TSStructs.RelationalValidationsBefore("RelMayorQ", op1, op2)

	if err != nil {
		ctx.AddException(err.Error(), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch op1.TSType {
	case TSStructs.INTEGER:
		return TSStructs.NewTBoolean(op1.Ivalue > op2.Ivalue)
	case TSStructs.FLOAT:
		return TSStructs.NewTBoolean(op1.Fvalue > op2.Fvalue)
	case TSStructs.STRING:
		return TSStructs.NewTBoolean(op1.Svalue > op2.Svalue)
	case TSStructs.CHARACTER:
		return TSStructs.NewTBoolean(op1.Svalue > op2.Svalue)
	}

	ctx.AddException("RelMayorQ: tipos no permitidos para esta operación.", I.Line, I.Position)
	return TSStructs.NewTNil()
}

/**********************************************************************************/
type IRelLessEqual struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIRelLessEqual(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IRelLessEqual {
	return &IRelLessEqual{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IRelLessEqual) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	op1 := I.op1.Interpret(ctx)
	op2 := I.op2.Interpret(ctx)

	err := TSStructs.RelationalValidationsBefore("RelMayorQ", op1, op2)

	if err != nil {
		ctx.AddException(err.Error(), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch op1.TSType {
	case TSStructs.INTEGER:
		return TSStructs.NewTBoolean(op1.Ivalue <= op2.Ivalue)
	case TSStructs.FLOAT:
		return TSStructs.NewTBoolean(op1.Fvalue <= op2.Fvalue)
	case TSStructs.STRING:
		return TSStructs.NewTBoolean(op1.Svalue <= op2.Svalue)
	case TSStructs.CHARACTER:
		return TSStructs.NewTBoolean(op1.Svalue <= op2.Svalue)
	}

	ctx.AddException("RelMayorQ: tipos no permitidos para esta operación.", I.Line, I.Position)
	return TSStructs.NewTNil()
}

/**********************************************************************************/
type IRelGreaterEqual struct {
	TSStructs.TSExpression
	op1 TSStructs.TSExpressioner
	op2 TSStructs.TSExpressioner
}

func NewIRelGreaterEqual(Line int, Position int, op1 TSStructs.TSExpressioner, op2 TSStructs.TSExpressioner) *IRelGreaterEqual {
	return &IRelGreaterEqual{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IRelGreaterEqual) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	op1 := I.op1.Interpret(ctx)
	op2 := I.op2.Interpret(ctx)

	err := TSStructs.RelationalValidationsBefore("RelMayorQ", op1, op2)

	if err != nil {
		ctx.AddException(err.Error(), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch op1.TSType {
	case TSStructs.INTEGER:
		return TSStructs.NewTBoolean(op1.Ivalue >= op2.Ivalue)
	case TSStructs.FLOAT:
		return TSStructs.NewTBoolean(op1.Fvalue >= op2.Fvalue)
	case TSStructs.STRING:
		return TSStructs.NewTBoolean(op1.Svalue >= op2.Svalue)
	case TSStructs.CHARACTER:
		return TSStructs.NewTBoolean(op1.Svalue >= op2.Svalue)
	}

	ctx.AddException("RelMayorQ: tipos no permitidos para esta operación.", I.Line, I.Position)
	return TSStructs.NewTNil()
}

/**********************************************************************************/
