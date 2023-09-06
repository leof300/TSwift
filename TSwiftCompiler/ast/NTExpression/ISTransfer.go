package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type IContinue struct {
	TSStructs.TSExpression
}

func NewIContinue(Line int, Position int) *IContinue {
	return &IContinue{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
	}
}

func (I IContinue) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	return TExpression.NewTNil()
}

type IBreak struct {
	TSStructs.TSExpression
}

func NewIBreak(Line int, Position int) *IBreak {
	return &IBreak{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
	}
}

func (I IBreak) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	return TExpression.NewTNil()
}

type IReturn struct {
	TSStructs.TSExpression
	expr TSStructs.TSExpressioner
}

func NewIReturn(Line int, Position int, expr TSStructs.TSExpressioner) *IReturn {
	return &IReturn{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
		expr:         expr,
	}
}

func (I IReturn) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	expr := I.expr.Interpret(ctx)

	var returnValue *TExpression.TSValue

	switch expr.TSType {
	case TExpression.INTEGER:
		returnValue = TExpression.NewTInt(expr.Ivalue)
	case TExpression.FLOAT:
		returnValue = TExpression.NewTFloat(expr.Fvalue)
	case TExpression.STRING:
		returnValue = TExpression.NewTString(expr.Svalue)
	case TExpression.CHARACTER:
		returnValue = TExpression.NewTString(expr.Svalue)
	case TExpression.BOOL:
		returnValue = TExpression.NewTBoolean(expr.Bvalue)
	case TExpression.NIL:
		returnValue = TExpression.NewTNil()
	}

	returnValue.IsReturn = true

	return returnValue
}

type INoReturn struct {
	TSStructs.TSExpression
}

func NewINoReturn(Line int, Position int) *INoReturn {
	return &INoReturn{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
	}
}

func (I INoReturn) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	return TExpression.NewTNil()
}
