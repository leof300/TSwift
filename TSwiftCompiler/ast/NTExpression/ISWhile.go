package NTExpression

import (
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

func (I ISWhile) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	ereturn := TSStructs.NewTNil()

	expr := I.expr.Interpret(ctx)

	if expr.IsNil || expr.TSType != TSStructs.BOOL {
		ctx.AddException("WHILE: condición no se puede evaluar.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	for expr.Bvalue {
		ereturn := I.block.Interpret(ctx)

		if ereturn.IsBreak || ereturn.IsReturn {
			return ereturn
		}

		expr = I.expr.Interpret(ctx)
	}

	return ereturn
}
