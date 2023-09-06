package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type ITSBlock struct {
	TSStructs.TSExpression
	sentences TSStructs.TSExpressioner
}

func NewITSBlock(Line int, Position int, sentences TSStructs.TSExpressioner) *ITSBlock {
	return &ITSBlock{
		TSExpression: TSStructs.TSExpression{Line, Position, make([]string, 0)},
		sentences:    sentences,
	}
}

func (I ITSBlock) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {

	ctx.AddScope()

	result := I.sentences.Interpret(ctx)

	ctx.RemoveScope()

	return result
}
