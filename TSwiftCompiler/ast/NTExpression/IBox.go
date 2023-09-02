package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
	"fmt"
)

type IBoxCreation struct {
	TSStructs.TSExpression
	BKey       string
	BInitValue TSStructs.TSExpressioner
	BType      TExpression.TSPTYPES
}

func NewIBoxCreation(Line int, Position int, initValue TSStructs.TSExpressioner, btype TExpression.TSPTYPES, key string) *IBoxCreation {
	return &IBoxCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		key, initValue, btype,
	}
}

func (I IBoxCreation) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {

	ivalue := TExpression.NewTNil()

	if I.BInitValue != nil {
		ivalue = I.BInitValue.Interpret(ctx)
	}

	var box = TExpression.NewTBox(I.BKey, ivalue, I.BType)

	ok := ctx.AddVariable(I.BKey, box, I.Line, I.Position)

	if !ok {
		ctx.AddException(fmt.Sprintf("No se puede agregar la variable %s", I.BKey), I.Line, I.Position)
	}
	return box
}

// /// GETING
type IBoxGetting struct {
	TSStructs.TSExpression
	key string
}

func NewIBoxGetting(Line int, Position int, key string) *IBoxGetting {
	return &IBoxGetting{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		key,
	}
}

func (I IBoxGetting) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	box := ctx.GetVariable(I.key)

	if box == nil {
		ctx.AddException(fmt.Sprintf("No se encuentra la variable %s", I.key), I.Line, I.Position)
	}
	return box
}
