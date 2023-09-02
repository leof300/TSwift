package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type ILSentences struct {
	TSStructs.TSExpression
	Sentences []TSStructs.TSExpressioner
}

func NewILSentences() *ILSentences {
	return &ILSentences{Sentences: make([]TSStructs.TSExpressioner, 0)}
}

func (I ILSentences) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {

	for _, sentence := range I.Sentences {
		sentence.Interpret(ctx)
	}

	return TExpression.NewTNil()
}
