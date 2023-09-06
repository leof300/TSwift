package NTExpression

import (
	"TSwiftCompiler/ast/TExpression"
	"TSwiftCompiler/ast/TSStructs"
)

type ISwitch struct {
	TSStructs.TSExpression
	expr    TSStructs.TSExpressioner
	Icases  []ICase
	Idfault TSStructs.TSExpressioner
}

func NewISwitch(Line int, Position int, expr TSStructs.TSExpressioner, idfault TSStructs.TSExpressioner) *ISwitch {
	return &ISwitch{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		expr, make([]ICase, 0), idfault,
	}
}

func (I ISwitch) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	expr := I.expr.Interpret(ctx)

	if expr.IsUndefined() {
		ctx.AddException("SWITCH: La expresión no es válida.", I.Line, I.Position)
		//TODO: SI LA EXPRESION ES INVALIDA NO SE EJECUTA EL SWITCH
		return TExpression.NewTNil()
	}

	wasEvaluated := false
	for _, icase := range I.Icases {

		exprCase := icase.expr.Interpret(ctx)

		err := TSStructs.RelationalValidationsBefore("Switch", expr, exprCase)

		if err != nil {
			ctx.AddException(err.Error(), icase.Line, icase.Position)
			return TExpression.NewTNil()
		}

		switch expr.TSType {
		case TExpression.INTEGER:
			if expr.Ivalue == exprCase.Ivalue {
				icase.lsents.Interpret(ctx)
				wasEvaluated = true
			}
		case TExpression.FLOAT:
			if expr.Fvalue == exprCase.Fvalue {
				icase.lsents.Interpret(ctx)
				wasEvaluated = true
			}
		case TExpression.BOOL:
			if expr.Bvalue == exprCase.Bvalue {
				icase.lsents.Interpret(ctx)
				wasEvaluated = true
			}
		case TExpression.STRING:
			if expr.Svalue == exprCase.Svalue {
				icase.lsents.Interpret(ctx)
				wasEvaluated = true
			}
		case TExpression.CHARACTER:
			if expr.Svalue == exprCase.Svalue {
				icase.lsents.Interpret(ctx)
				wasEvaluated = true
			}
		}

		if wasEvaluated {
			break
		}
	}

	//si no hizo match con ninguno se pasa al default
	if !wasEvaluated && I.Idfault != nil {
		I.Idfault.Interpret(ctx)
	}
	return TExpression.NewTNil()
}

type ICase struct {
	TSStructs.TSExpression

	expr   TSStructs.TSExpressioner
	lsents TSStructs.TSExpressioner
}

func NewICase(Line int, Position int, expr TSStructs.TSExpressioner, lsents TSStructs.TSExpressioner) *ICase {
	return &ICase{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		expr, lsents,
	}
}

func (I ICase) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	I.lsents.Interpret(ctx)
	return TExpression.NewTNil()
}

type IDefault struct {
	TSStructs.TSExpression
	lsents TSStructs.TSExpressioner
}

func NewIDefault(Line int, Position int, lsents TSStructs.TSExpressioner) *IDefault {
	return &IDefault{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		lsents,
	}
}

func (I IDefault) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
	I.lsents.Interpret(ctx)
	return TExpression.NewTNil()
}
