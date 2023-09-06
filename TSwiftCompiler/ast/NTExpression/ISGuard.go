package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type ISGuard struct {
	TSStructs.TSExpression
	expr  TSStructs.TSExpressioner
	block TSStructs.TSExpressioner
}

func NewISGuard(Line int, Position int, expr TSStructs.TSExpressioner, block TSStructs.TSExpressioner) *ISGuard {
	return &ISGuard{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
		expr:         expr, block: block,
	}
}

func (I ISGuard) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	result := TExpression.NewTNil()

	expr := I.expr.Interpret(ctx)

	if expr.IsNil || expr.TSType != TExpression.BOOL {
		ctx.AddException("GUARD: condición no se puede evaluar.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	for !expr.Bvalue {
		result := I.block.Interpret(ctx)

		if result.IsReturn || result.IsBreak {
			return result
		}

		expr = I.expr.Interpret(ctx)
	}

	return result
}
