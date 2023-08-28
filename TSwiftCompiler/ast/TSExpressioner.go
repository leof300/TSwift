package ast

import (
	"TSwiftCompiler/ast/TExpression"
	"fmt"
)

// intefaz en Go
type TSExpressioner interface {
	Interpret(ctx *TSContext) TExpression.TValue
}

type TSExpression struct {
	Line     int
	Position int
	TSlog    []string
}

func (I TSExpression) SafeInterpret(ctx *TSContext) TExpression.TValue {
	a, ok := I.Interpret(ctx).(TExpression.TValue)
	if !ok {
		//TODO: agregar error semantico
		//I.TSlog = append(I.TSlog, "Semantic exception")
		fmt.Print("nodo inválido")
	}
	return a
}

func (I TSExpression) Interpret(ctx *TSContext) TExpression.TValue {
	return nil
}
