package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
)

type IFParameter struct {
	TSStructs.TSExpression
	Alias       string
	Key         string
	PType       TSStructs.TSExpressioner
	PassByValue bool
}

func (I IFParameter) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	ptype := I.PType.Interpret(ctx)

	var box = TSStructs.NewTBox(I.Key, ptype.TSType)
	box.IsBox = true //si es una box

	box.FuncParameterAlias = I.Alias

	box.IsConstant = I.PassByValue

	return box
}

func NewIFParameter(Line int, Position int, alias string, key string, ptype TSStructs.TSExpressioner, passByValue bool) *IFParameter {
	return &IFParameter{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		alias, key, ptype, passByValue,
	}
}
