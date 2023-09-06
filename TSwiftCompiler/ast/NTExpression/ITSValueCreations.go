package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type IIntCreation struct {
	TSStructs.TSExpression
	value int
}

func NewIIntCreation(Line int, Position int, value int) *IIntCreation {
	return &IIntCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		value,
	}
}

func (I IIntCreation) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	return TExpression.NewTInt(I.value)
}

type IFloatCreation struct {
	TSStructs.TSExpression
	value float64
}

func NewIFloatCreation(Line int, Position int, value float64) *IFloatCreation {
	return &IFloatCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		value,
	}
}

func (I IFloatCreation) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	return TExpression.NewTFloat(I.value)
}

type IBooleanCreation struct {
	TSStructs.TSExpression
	value bool
}

func NewIBooleanCreation(Line int, Position int, value bool) *IBooleanCreation {
	return &IBooleanCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		value,
	}
}
func (I IBooleanCreation) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	return TExpression.NewTBoolean(I.value)
}

type IStringCreation struct {
	TSStructs.TSExpression
	value string
}

func NewIStringCreation(Line int, Position int, value string) *IStringCreation {
	return &IStringCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		value,
	}
}

func (I IStringCreation) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	return TExpression.NewTString(I.value)
}

type IUndefinedCreation struct {
	TSStructs.TSExpression
}

func NewIUndefinedCreation(Line int, Position int) *IUndefinedCreation {
	return &IUndefinedCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
	}
}

func (I IUndefinedCreation) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	return TExpression.NewTUndefined()
}

type INilCreation struct {
	TSStructs.TSExpression
}

func NewINilCreation(Line int, Position int) *INilCreation {
	return &INilCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
	}
}

func (I INilCreation) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	return TExpression.NewTNil()
}
