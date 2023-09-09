package TSStructs

type TSExpressioner interface {
	Interpret(ctx *TSContext) *TSValue
}

type TSExpression struct {
	Line     int
	Position int
	TSlog    []string
}

func (I TSExpression) Interpret(ctx *TSContext) *TSValue {
	return nil
}
