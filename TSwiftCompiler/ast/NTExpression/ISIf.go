package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
)

type ITSIf struct {
	TSStructs.TSExpression
	expr  TSStructs.TSExpressioner
	block TSStructs.TSExpressioner
}

func NewITSIf(Line int, Position int, expr TSStructs.TSExpressioner, block TSStructs.TSExpressioner) *ITSIf {
	return &ITSIf{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
		expr:         expr, block: block,
	}
}

func (I ITSIf) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	erturn := TSStructs.NewTNil()
	expr := I.expr.Interpret(ctx)

	if expr.IsNil || expr.TSType != TSStructs.BOOL {
		ctx.AddException("IF: condición no se puede evaluar.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	if expr.Bvalue {
		erturn = I.block.Interpret(ctx)
	}

	return erturn
}

type ITSIfElse struct {
	TSStructs.TSExpression
	expr      TSStructs.TSExpressioner
	blockIf   TSStructs.TSExpressioner
	blockElse TSStructs.TSExpressioner
}

func NewITSIfElse(Line int, Position int, expr TSStructs.TSExpressioner, blockIf TSStructs.TSExpressioner, blockElse TSStructs.TSExpressioner) *ITSIfElse {
	return &ITSIfElse{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
		expr:         expr, blockIf: blockIf, blockElse: blockElse,
	}
}

func (I ITSIfElse) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	erturn := TSStructs.NewTNil()
	expr := I.expr.Interpret(ctx)

	if expr.IsNil || expr.TSType != TSStructs.BOOL {
		ctx.AddException("IF: condición no se puede evaluar.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	if expr.Bvalue {
		erturn = I.blockIf.Interpret(ctx)
	} else {
		erturn = I.blockElse.Interpret(ctx)
	}

	return erturn
}

type ITSIfElseIf struct {
	TSStructs.TSExpression
	expr        TSStructs.TSExpressioner
	blockIf     TSStructs.TSExpressioner
	blockElseIf TSStructs.TSExpressioner
}

func NewITSIfElseIf(Line int, Position int, expr TSStructs.TSExpressioner, blockIf TSStructs.TSExpressioner, blockElse TSStructs.TSExpressioner) *ITSIfElseIf {
	return &ITSIfElseIf{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
		expr:         expr, blockIf: blockIf, blockElseIf: blockElse,
	}
}

func (I ITSIfElseIf) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	erturn := TSStructs.NewTNil()
	expr := I.expr.Interpret(ctx)

	if expr.IsNil || expr.TSType != TSStructs.BOOL {
		ctx.AddException("IF: condición no se puede evaluar.", I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	if expr.Bvalue {
		erturn = I.blockIf.Interpret(ctx)
	} else {
		erturn = I.blockElseIf.Interpret(ctx)
	}

	return erturn
}
