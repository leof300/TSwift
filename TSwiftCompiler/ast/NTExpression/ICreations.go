package NTExpression

import (
	"TSwiftCompiler/ast"
	"TSwiftCompiler/ast/TExpression"
)

type IIntCreation struct {
	ast.TSExpression
	value int
}

func NewIIntCreation(Line int, Position int, value int) *IIntCreation {
	return &IIntCreation{
		ast.TSExpression{Line, Position, make([]string, 0)},
		value,
	}
}

func (I IIntCreation) interpret(ctx *ast.TSContext) TExpression.TValue {
	return TExpression.NewTInt(I.value)
}

type IFloatCreation struct {
	ast.TSExpression
	value float64
}

func NewIFloatCreation(Line int, Position int, value float64) *IFloatCreation {
	return &IFloatCreation{
		ast.TSExpression{Line, Position, make([]string, 0)},
		value,
	}
}

func (I IFloatCreation) interpret(ctx *ast.TSContext) TExpression.TValue {
	return TExpression.NewTFloat(I.value)
}

type IBooleanCreation struct {
	ast.TSExpression
	value bool
}

func NewIBooleanCreation(Line int, Position int, value bool) *IBooleanCreation {
	return &IBooleanCreation{
		ast.TSExpression{Line, Position, make([]string, 0)},
		value,
	}
}
func (I IBooleanCreation) interpret(ctx *ast.TSContext) TExpression.TValue {
	return TExpression.NewTBoolean(I.value)
}

type IStringCreation struct {
	ast.TSExpression
	value string
}

func NewIStringCreation(Line int, Position int, value string) *IStringCreation {
	return &IStringCreation{
		ast.TSExpression{Line, Position, make([]string, 0)},
		value,
	}
}

func (I IStringCreation) interpret(ctx *ast.TSContext) TExpression.TValue {
	return TExpression.NewTString(I.value)
}

type IBox struct {
	ast.TSExpression
	key string
}

func NewIBox(Line int, Position int, key string) *IBox {
	return &IBox{
		ast.TSExpression{Line, Position, make([]string, 0)},
		key,
	}
}

func (I IBox) interpret(ctx *ast.TSContext) TExpression.TValue {
	//TODO: BUSCAR LA VARIABLE Y SI NO ESTA ENVIAR UN ERROR DE EJECUCION
	return TExpression.NewTBox(I.key)
}
