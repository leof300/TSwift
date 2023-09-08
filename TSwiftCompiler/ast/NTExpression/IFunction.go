package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
	"fmt"
)

type IFunction struct {
	TSStructs.TSExpression
	fKey        string
	FParameters []TSStructs.TSExpressioner
	fType       TSStructs.TSExpressioner
	fBody       TSStructs.TSExpressioner
}

func (I IFunction) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {

	//una variable es undefined si no devuelve nada
	ftype := TSStructs.NIL

	if I.fType != nil {
		fReturnType := I.fType.Interpret(ctx)
		ftype = fReturnType.TSType
	}

	nFunction := TSStructs.NewFunction(I.fKey, ftype)

	//agregamos todos los parametros
	for _, p := range I.FParameters {
		param := p.Interpret(ctx)

		//parametro o alias definido mas de una vez
		_, exists := nFunction.FuncParameters[param.FuncParameterAlias]
		if exists {
			ctx.AddException(fmt.Sprintf("FuncDeclaration: el parametro/alias %s ya existen para la función.", param.FuncParameterAlias), I.Line, I.Position)
			return TSStructs.NewTNil()
		}
		nFunction.FuncParameters[param.FuncParameterAlias] = param

	}

	nFunction.FuncValue = I.fBody

	ok := ctx.AddFunction(I.fKey, nFunction, I.Line, I.Position)

	if !ok {
		ctx.AddException(fmt.Sprintf("FuncDeclaration: esta funcion %s ya fue definida.", I.fKey), I.Line, I.Position)
	}

	return TSStructs.NewTNil()
}

func NewIFunction(Line int, Position int, fKey string, fType TSStructs.TSExpressioner, fBody TSStructs.TSExpressioner) *IFunction {
	return &IFunction{
		TSExpression: TSStructs.TSExpression{Line: Line, Position: Position, TSlog: make([]string, 0)},
		fKey:         fKey, fType: fType, fBody: fBody}
}
