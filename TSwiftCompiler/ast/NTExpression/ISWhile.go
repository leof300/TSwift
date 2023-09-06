package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type ISWhile struct {
	TSStructs.TSExpression
	expr  TSStructs.TSExpressioner
	block TSStructs.TSExpressioner
}

func NewISWhile(Line int, Position int, expr TSStructs.TSExpressioner, block TSStructs.TSExpressioner) *ISWhile {
	return &ISWhile{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
		expr:         expr, block: block,
	}
}

func (I ISWhile) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	expr := I.expr.Interpret(ctx)

	if expr.IsNil || expr.TSType != TExpression.BOOL {
		ctx.AddException("WHILE: condición no se puede evaluar.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	for expr.Bvalue {
		finishCycle := I.block.Interpret(ctx)

		if finishCycle.Bvalue {
			break
		}

		expr = I.expr.Interpret(ctx)
	}

	return TExpression.NewTNil()
}
