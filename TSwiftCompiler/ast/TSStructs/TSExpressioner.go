package TSStructs

import "TSwiftCompiler/ast/TExpression"

// intefaz en Go
type TSExpressioner interface {
	Interpret(ctx *TSContext) *TExpression.TSValue
}

type TSExpression struct {
	Line     int
	Position int
	TSlog    []string
}

//func (I TSExpression) SafeInterpret(ctx *TSContext) *TExpression.TSValue {
//	a, ok := I.Interpret(ctx).(TExpression.TSValue)
//	if !ok {
//		//TODO: agregar error semantico
//		//I.TSlog = append(I.TSlog, "Semantic exception")
//		fmt.Print("nodo inválido")
//	}
//	return a
//}

func (I TSExpression) Interpret(ctx *TSContext) *TExpression.TSValue {
	return nil
}
