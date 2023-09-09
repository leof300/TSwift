package NTExpression

import "TSwiftCompiler/ast/TSStructs"

type IArrayCreation struct {
	TSStructs.TSExpression
	ExprValues []TSStructs.TSExpressioner
}

func NewIArrayCreation(Line int, Position int) *IArrayCreation {
	return &IArrayCreation{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		make([]TSStructs.TSExpressioner, 0),
	}
}
func (I IArrayCreation) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	typePivot := TSStructs.NIL

	result := TSStructs.NewTNil()
	result.IsArray = true
	result.ArrayContent = make([]TSStructs.TSValue, 0)

	for _, expr := range I.ExprValues {
		vresult := expr.Interpret(ctx)

		if typePivot != TSStructs.NIL && typePivot != vresult.TSType {
			ctx.AddException("ArrayCreation: tipos de datos distintos.", I.Line, I.Position)
		}

		result.ArrayContent = append(result.ArrayContent, *TSStructs.CopyTSValueNew(vresult))
		typePivot = vresult.TSType
	}

	result.TSType = typePivot
	result.IsNil = false
	result.IsBox = true

	return result
}
