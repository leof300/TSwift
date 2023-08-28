package NTExpression

import (
	"TSwiftCompiler/ast"
	"TSwiftCompiler/ast/TExpression"
)

type IAdd struct {
	ast.TSExpression
	op1 ast.TSExpression
	op2 ast.TSExpression
}

func NewIAdd(Line int, Position int, op1 ast.TSExpression, op2 ast.TSExpression) *IAdd {
	return &IAdd{
		ast.TSExpression{Line, Position, make([]string, 0)},
		op1, op2,
	}
}

func (I IAdd) interpret(ctx *ast.TSContext) TExpression.TValue {
	
	//return TExpression.NewTString(I.value)
	return nil
}
