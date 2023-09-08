package NTExpression

import (
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

func (I ISwitch) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	ereturn := TSStructs.NewTNil()
	expr := I.expr.Interpret(ctx)

	if expr.IsUndefined() {
		ctx.AddException("SWITCH: La expresión no es válida.", I.Line, I.Position)
		//TODO: SI LA EXPRESION ES INVALIDA NO SE EJECUTA EL SWITCH
		return TSStructs.NewTNil()
	}

	wasEvaluated := false
	for _, icase := range I.Icases {

		exprCase := icase.expr.Interpret(ctx)

		err := TSStructs.RelationalValidationsBefore("Switch", expr, exprCase)

		if err != nil {
			ctx.AddException(err.Error(), icase.Line, icase.Position)
			return TSStructs.NewTNil()
		}

		switch expr.TSType {
		case TSStructs.INTEGER:
			if expr.Ivalue == exprCase.Ivalue {
				ereturn = icase.lsents.Interpret(ctx)
				wasEvaluated = true
			}
		case TSStructs.FLOAT:
			if expr.Fvalue == exprCase.Fvalue {
				ereturn = icase.lsents.Interpret(ctx)
				wasEvaluated = true
			}
		case TSStructs.BOOL:
			if expr.Bvalue == exprCase.Bvalue {
				ereturn = icase.lsents.Interpret(ctx)
				wasEvaluated = true
			}
		case TSStructs.STRING:
			if expr.Svalue == exprCase.Svalue {
				ereturn = icase.lsents.Interpret(ctx)
				wasEvaluated = true
			}
		case TSStructs.CHARACTER:
			if expr.Svalue == exprCase.Svalue {
				ereturn = icase.lsents.Interpret(ctx)
				wasEvaluated = true
			}
		}

		if wasEvaluated {
			break
		}
	}

	//si no hizo match con ninguno se pasa al default
	if !wasEvaluated && I.Idfault != nil {
		ereturn = I.Idfault.Interpret(ctx)
	}

	return ereturn
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

func (I ICase) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	I.lsents.Interpret(ctx)
	return TSStructs.NewTNil()
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

func (I IDefault) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	I.lsents.Interpret(ctx)
	return TSStructs.NewTNil()
}
