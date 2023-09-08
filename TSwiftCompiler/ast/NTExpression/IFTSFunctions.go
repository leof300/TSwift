package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
	"strconv"
)

/*******************************************************/
type IFString struct {
	TSStructs.TSExpression
	expr TSStructs.TSExpressioner
}

func NewIFString(Line int, Position int, expr TSStructs.TSExpressioner) *IFString {
	return &IFString{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		expr,
	}
}

func (I IFString) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	expr := I.expr.Interpret(ctx)

	if expr.IsNil {
		ctx.AddException("ToString: La expresión no es válida para convertir.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch expr.TSType {
	case TSStructs.INTEGER:
		return TSStructs.NewTString(strconv.Itoa(expr.Ivalue))
	case TSStructs.FLOAT:
		return TSStructs.NewTString(strconv.FormatFloat(expr.Fvalue, 'f', -1, 64))
	case TSStructs.BOOL:
		return TSStructs.NewTString(strconv.FormatBool(expr.Bvalue))
	}

	ctx.AddException("ToString: Este tipo de valor no se puede convertir.", I.Line, I.Position)
	return TSStructs.NewTNil()
}

/*******************************************************/
type IFInt struct {
	TSStructs.TSExpression
	expr TSStructs.TSExpressioner
}

func NewIFInt(Line int, Position int, expr TSStructs.TSExpressioner) *IFInt {
	return &IFInt{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		expr,
	}
}

func (I IFInt) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	expr := I.expr.Interpret(ctx)

	if expr.IsNil {
		ctx.AddException("ToInteger: La expresión no es válida para convertir.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch expr.TSType {
	case TSStructs.FLOAT:
		return TSStructs.NewTInt(int(expr.Fvalue))
	case TSStructs.STRING:
		intvalue, err := strconv.Atoi(expr.Svalue)
		if err != nil {
			ctx.AddException("ToInteger: La expresión no es válida para convertir.", I.Line, I.Position)
			return TSStructs.NewTNil()
		}
		return TSStructs.NewTInt(intvalue)
	}

	ctx.AddException("ToInteger: Este tipo de valor no se puede convertir.", I.Line, I.Position)
	return TSStructs.NewTNil()
}

/*******************************************************/
type IFFloat struct {
	TSStructs.TSExpression
	expr TSStructs.TSExpressioner
}

func NewIFFloat(Line int, Position int, expr TSStructs.TSExpressioner) *IFFloat {
	return &IFFloat{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		expr,
	}
}

func (I IFFloat) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	expr := I.expr.Interpret(ctx)

	if expr.IsNil {
		ctx.AddException("TFloat: La expresión no es válida para convertir.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	switch expr.TSType {
	case TSStructs.STRING:
		valFloat, err := strconv.ParseFloat(expr.Svalue, 64)
		if err != nil {
			ctx.AddException("TFloat: La expresión no es válida para convertir.", I.Line, I.Position)
			return TSStructs.NewTNil()
		}
		return TSStructs.NewTFloat(valFloat)
	case TSStructs.INTEGER:
		return TSStructs.NewTFloat(float64(expr.Ivalue))
	}

	ctx.AddException("TFloat: Este tipo de valor no se puede convertir.", I.Line, I.Position)
	return TSStructs.NewTNil()
}
