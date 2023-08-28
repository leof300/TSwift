package NTExpression

import (
	"TSwiftCompiler/ast"
	"TSwiftCompiler/ast/TExpression"
)

type ILSentences struct {
	ast.TSExpression
	Sentences []ast.TSExpressioner
}

func NewILSentences() *ILSentences {
	return &ILSentences{Sentences: make([]ast.TSExpressioner, 0)}
}

func (I ILSentences) interpret(ctx *ast.TSContext) TExpression.TValue {
	return nil
}
