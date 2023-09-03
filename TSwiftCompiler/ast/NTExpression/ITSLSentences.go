package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
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
		sentence.Interpret(ctx)
	}

	return TExpression.NewTNil()
}
