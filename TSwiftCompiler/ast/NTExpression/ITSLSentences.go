package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
	"reflect"
)

type ITSLSentences struct {
	TSStructs.TSExpression
	Sentences []TSStructs.TSExpressioner
}

func NewILSentences() *ITSLSentences {
	return &ITSLSentences{Sentences: make([]TSStructs.TSExpressioner, 0)}
}

func (I ITSLSentences) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {

	for _, sentence := range I.Sentences {
		//TODO AQUI AGREGAR BREAK, RETURN, ETC

		if reflect.TypeOf(sentence).String() == "*NTExpression.IContinue" {
			returnValue := TExpression.NewTNil()
			returnValue.IsContinue = true
			return returnValue
		}

		if reflect.TypeOf(sentence).String() == "*NTExpression.IBreak" {
			returnValue := TExpression.NewTNil()
			returnValue.IsBreak = true
			return returnValue
		}

		if reflect.TypeOf(sentence).String() == "*NTExpression.INoReturn" {
			returnValue := TExpression.NewTNil()
			returnValue.IsReturn = true
			return returnValue
		}

		if reflect.TypeOf(sentence).String() == "*NTExpression.IReturn" {
			returnValue := sentence.Interpret(ctx)
			returnValue.IsReturn = true
			return returnValue
		}

		sentence.Interpret(ctx)

	}

	return TExpression.NewTNil()
}
