package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
)

// *************************************
type IContinue struct {
	TSStructs.TSExpression
}

func NewIContinue(Line int, Position int) *IContinue {
	return &IContinue{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
	}
}

func (I IContinue) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	result := TSStructs.NewTNil()
	result.IsContinue = true
	return result
}

// *************************************
type IBreak struct {
	TSStructs.TSExpression
}

func NewIBreak(Line int, Position int) *IBreak {
	return &IBreak{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
	}
}

func (I IBreak) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	result := TSStructs.NewTNil()
	result.IsBreak = true
	return result
}

/***************************************/
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

func (I IReturn) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {

	expr := I.expr.Interpret(ctx)

	var returnValue *TSStructs.TSValue

	switch expr.TSType {
	case TSStructs.INTEGER:
		returnValue = TSStructs.NewTInt(expr.Ivalue)
	case TSStructs.FLOAT:
		returnValue = TSStructs.NewTFloat(expr.Fvalue)
	case TSStructs.STRING:
		returnValue = TSStructs.NewTString(expr.Svalue)
	case TSStructs.CHARACTER:
		returnValue = TSStructs.NewTString(expr.Svalue)
	case TSStructs.BOOL:
		returnValue = TSStructs.NewTBoolean(expr.Bvalue)
	case TSStructs.NIL:
		returnValue = TSStructs.NewTNil()
	}

	returnValue.IsReturn = true

	return returnValue
}

/**************************************/
type INoReturn struct {
	TSStructs.TSExpression
}

func NewINoReturn(Line int, Position int) *INoReturn {
	return &INoReturn{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
	}
}

func (I INoReturn) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	result := TSStructs.NewTNil()
	result.IsReturn = true
	return result
}
