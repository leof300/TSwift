package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
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

func (I IRelEqual) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	op1 := I.op1.Interpret(ctx)
	op2 := I.op2.Interpret(ctx)

	////valores invalidos
	//if op1.IsNil || op2.IsNil {
	//	ctx.AddException("Igualdad: valor NIL invalido para comparación.", I.Line, I.Position)
	//	return TExpression.NewTNil()
	//}
	//
	//op1Type := op1.TSType
	//op2Type := op2.TSType
	//
	////error tipos no iguales
	//if op1Type != op2Type {
	//	ctx.AddException("Igualdad: tipos distintos para comparación.", I.Line, I.Position)
	//	return TExpression.NewTNil()
	//}

	err := TSStructs.RelationalValidationsBefore("RelMenorQ", op1, op2)

	if err != nil {
		ctx.AddException(err.Error(), I.Line, I.Position)
		return TExpression.NewTNil()
	}

	switch op1.TSType {
	case TExpression.INTEGER:
		return TExpression.NewTBoolean(op1.Ivalue == op2.Ivalue)
	case TExpression.FLOAT:
		return TExpression.NewTBoolean(op1.Fvalue == op2.Fvalue)
	case TExpression.BOOL:
		return TExpression.NewTBoolean(op1.Bvalue == op2.Bvalue)
	case TExpression.STRING:
		return TExpression.NewTBoolean(op1.Svalue == op2.Svalue)
	case TExpression.CHARACTER:
		return TExpression.NewTBoolean(op1.Svalue == op2.Svalue)
	}

	ctx.AddException("Igualdad: no permitidos para esta operación.", I.Line, I.Position)
	return TExpression.NewTNil()
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

func (I IRelNotEqual) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	op1 := I.op1.Interpret(ctx)
	op2 := I.op2.Interpret(ctx)

	err := TSStructs.RelationalValidationsBefore("RelMenorQ", op1, op2)

	if err != nil {
		ctx.AddException(err.Error(), I.Line, I.Position)
		return TExpression.NewTNil()
	}

	switch op1.TSType {
	case TExpression.INTEGER:
		return TExpression.NewTBoolean(op1.Ivalue != op2.Ivalue)
	case TExpression.FLOAT:
		v := op1.Fvalue != op2.Fvalue
		return TExpression.NewTBoolean(v)
	case TExpression.BOOL:
		return TExpression.NewTBoolean(op1.Bvalue != op2.Bvalue)
	case TExpression.STRING:
		return TExpression.NewTBoolean(op1.Svalue != op2.Svalue)
	case TExpression.CHARACTER:
		return TExpression.NewTBoolean(op1.Svalue != op2.Svalue)
	}

	ctx.AddException("Desigualdad: tipos no permitidos para esta operación.", I.Line, I.Position)
	return TExpression.NewTNil()
}
