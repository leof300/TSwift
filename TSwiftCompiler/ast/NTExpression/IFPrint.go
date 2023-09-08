package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
)

/*
*

	type IFPrint struct {
		TSStructs.TSExpression
		expr TSStructs.TSExpressioner
	}

	func NewIFPrint(Line int, Position int, expr TSStructs.TSExpressioner) *IFPrint {
		return &IFPrint{
			TSStructs.TSExpression{Line, Position, make([]string, 0)},
			op1, op2,
		}
	}

	func (I IFPrint) Interpret(ctx *TSStructs.TSContext) *TExpression.TSValue {
		return nil
	}
*/
type IFPrint struct {
	TSStructs.TSExpression
	Exprs []TSStructs.TSExpressioner
}

func NewIFPrint(Line int, Position int) *IFPrint {
	return &IFPrint{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		make([]TSStructs.TSExpressioner, 0),
	}
}

func (I IFPrint) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	expressions := I.Exprs

	msg := ""
	for _, e := range expressions {
		expr := e.Interpret(ctx)

		if expr.IsUndefined() {
			ctx.AddException("Print: La expresión no es válida.", I.Line, I.Position)
			return TSStructs.NewTNil()
		}

		if expr.IsNil {
			msg += "nil "
			continue
		}

		msg += expr.ToString() + " "
	}

	ctx.AddConsole(msg)

	return TSStructs.NewTNil()
}
