package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
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

func (I ITSIf) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	expr := I.expr.Interpret(ctx)

	if expr.IsNil || expr.TSType != TExpression.BOOL {
		ctx.AddException("IF: condición no se puede evaluar.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	if expr.Bvalue {
		I.block.Interpret(ctx)
	}

	return TExpression.NewTNil()
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

func (I ITSIfElse) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	expr := I.expr.Interpret(ctx)

	if expr.IsNil || expr.TSType != TExpression.BOOL {
		ctx.AddException("IF: condición no se puede evaluar.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	if expr.Bvalue {
		I.blockIf.Interpret(ctx)
	} else {
		I.blockElse.Interpret(ctx)
	}

	return TExpression.NewTNil()
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

func (I ITSIfElseIf) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	expr := I.expr.Interpret(ctx)

	if expr.IsNil || expr.TSType != TExpression.BOOL {
		ctx.AddException("IF: condición no se puede evaluar.", I.Line, I.Position)
		return TExpression.NewTNil()
	}

	if expr.Bvalue {
		I.blockIf.Interpret(ctx)
	} else {
		I.blockElseIf.Interpret(ctx)
	}

	return TExpression.NewTNil()
}
