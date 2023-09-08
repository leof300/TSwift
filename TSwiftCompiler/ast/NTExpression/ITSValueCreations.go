package NTExpression

import (
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

func (I IIntCreation) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	return TSStructs.NewTInt(I.value)
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

func (I IFloatCreation) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	return TSStructs.NewTFloat(I.value)
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
func (I IBooleanCreation) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	return TSStructs.NewTBoolean(I.value)
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

func (I IStringCreation) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	return TSStructs.NewTString(I.value)
}

type IUndefinedCreation struct {
	TSStructs.TSExpression
}

func NewIUndefinedCreation(Line int, Position int) *IUndefinedCreation {
	return &IUndefinedCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
	}
}

func (I IUndefinedCreation) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	return TSStructs.NewTUndefined()
}

type INilCreation struct {
	TSStructs.TSExpression
}

func NewINilCreation(Line int, Position int) *INilCreation {
	return &INilCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
	}
}

func (I INilCreation) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	return TSStructs.NewTNil()
}

/*
&&&&&&&&&&&&&&&&&&&&&&&&&&         TYPES         &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
*/
type ITSType struct {
	TSStructs.TSExpression
	tstype TSStructs.TSPTYPES
}

func NewITSType(Line int, Position int, etype TSStructs.TSPTYPES) *ITSType {
	return &ITSType{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		etype,
	}
}

func (I ITSType) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	ereturn := TSStructs.NewTNil()
	ereturn.TSType = I.tstype
	return ereturn
}
