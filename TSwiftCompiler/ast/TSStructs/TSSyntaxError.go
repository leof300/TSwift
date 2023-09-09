package TSStructs

type TSSyntaxError struct {
	TSExpression
	msg string
}

func NewTSSyntaxError(Line int, Position int, msg string) *TSSyntaxError {
	return &TSSyntaxError{
		TSExpression: TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
		msg:          msg,
	}
}

func (N *TSSyntaxError) Interpret(ctx *TSContext) *TSValue {
	ctx.AddException(N.msg, N.Line, N.Position)
	return NewTNil()
}
