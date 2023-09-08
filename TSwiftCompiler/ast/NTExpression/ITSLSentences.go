package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
)

type ITSLSentences struct {
	TSStructs.TSExpression
	Sentences []TSStructs.TSExpressioner
}

func NewILSentences() *ITSLSentences {
	return &ITSLSentences{Sentences: make([]TSStructs.TSExpressioner, 0)}
}

func (I ITSLSentences) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {

	for _, sentence := range I.Sentences {
		//TODO AQUI AGREGAR BREAK, RETURN, ETC

		returnValue := sentence.Interpret(ctx)

		if returnValue.IsReturn || returnValue.IsBreak || returnValue.IsContinue {
			return returnValue
		}

	}

	return TSStructs.NewTNil()
}
