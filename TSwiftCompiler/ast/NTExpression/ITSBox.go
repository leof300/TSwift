package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
	"fmt"
)

type ITSBoxCreation struct {
	TSStructs.TSExpression
	BKey       string
	BInitValue TSStructs.TSExpressioner
	IsArray    bool
}

func NewIBoxCreation(Line int, Position int, key string, initValue TSStructs.TSExpressioner, IsArray bool) *ITSBoxCreation {
	return &ITSBoxCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		key, initValue, IsArray,
	}
}

func (I ITSBoxCreation) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {

	ivalue := TSStructs.NewTNil()

	if I.BInitValue != nil {
		ivalue = I.BInitValue.Interpret(ctx)
	}

	var box *TSStructs.TSValue

	if I.IsArray {

		box = TSStructs.NewArray(I.BKey, ivalue.TSType)
		box.ArrayContent = ivalue.ArrayContent

	} else {
		box = TSStructs.NewTBox(I.BKey, ivalue.TSType)
		box.Svalue = ivalue.Svalue
		box.Bvalue = ivalue.Bvalue
		box.Ivalue = ivalue.Ivalue
		box.Fvalue = ivalue.Fvalue
		box.IsNil = ivalue.IsBox
	}

	ok := ctx.AddVariable(I.BKey, box, I.Line, I.Position)

	if !ok {
		ctx.AddException(fmt.Sprintf("BoxCreation: No se puede crear boxGetting.", I.BKey), I.Line, I.Position)
	}

	return box
}

// ////////////////////////////////////////////////////////////////
// /// GETING
// ////////////////////////////////////////////////////////////////
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

func (I IBoxGetting) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	box := ctx.GetVariable(I.key)

	if box == nil {
		ctx.AddException(fmt.Sprintf("ObtenerBox: No se encuentra la box %s", I.key), I.Line, I.Position)
		return TSStructs.NewTNil()
	}
	return box
}

// ////////////////////////////////////////////////////////////////
// /// CONST
// ////////////////////////////////////////////////////////////////
type IConstCreation struct {
	TSStructs.TSExpression
	boxGetting  TSStructs.TSExpressioner
	boxCreation TSStructs.TSExpressioner
}

func NewIConstCreation(Line int, Position int, boxGetting TSStructs.TSExpressioner, boxCreation TSStructs.TSExpressioner) *IConstCreation {
	return &IConstCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		boxGetting, boxCreation,
	}
}

func (I IConstCreation) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	I.boxCreation.Interpret(ctx)

	var box *TSStructs.TSValue = I.boxGetting.Interpret(ctx)

	if box.IsBox {
		box.IsConstant = true
	} else {
		ctx.AddException("ConstantCreation: No se puede crear la constante.", I.Line, I.Position)
	}

	return TSStructs.NewTNil()
}
