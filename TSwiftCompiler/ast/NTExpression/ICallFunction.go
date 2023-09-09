package NTExpression

import (
	"TSwiftCompiler/ast/TSStructs"
	"fmt"
)

type IFCallFunction struct {
	TSStructs.TSExpression
	Name       string
	Parameters []TSStructs.TSExpressioner
}

func NewIFCallFunction(Line int, Position int, name string) *IFCallFunction {
	return &IFCallFunction{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		name, make([]TSStructs.TSExpressioner, 0),
	}
}

func (I IFCallFunction) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	function := ctx.GetFunction(I.Name)

	if function == nil {
		ctx.AddException(fmt.Sprintf("Llamada Funcion: No se encuentra la fncion %s", I.Name), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	if !function.IsFunction {
		ctx.AddException(fmt.Sprintf("Llamada Funcion: No se encuentra la funcion %s", I.Name), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	if len(I.Parameters) != len(function.FuncParameters) {
		ctx.AddException(fmt.Sprintf("Llamada Funcion: El número de parámetros que se envio %d es incorrecto, se necesitan %d ", len(I.Parameters), len(function.FuncParameters)), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	//solo se ejecuta la función, ya que no se envian parametros
	if len(I.Parameters) == 0 {
		ctx.AddScope()

		result := function.FuncValue.Interpret(ctx)

		ctx.RemoveScope()

		//valor devuelto no coincide
		if result.TSType != function.TSType {
			ctx.AddException(fmt.Sprintf("Llamada Funcion: El valor de retorno es incorrecto: %s.", result.TSType), I.Line, I.Position)
			return TSStructs.NewTNil()
		}

		return result
	}

	//generamos los parametros, como un listado de ExprValues
	var paramValueList []*TSStructs.TSValue
	areParamsWithName := false

	for i := 0; i < len(I.Parameters); i++ {
		paramValue := I.Parameters[i].Interpret(ctx)

		if paramValue.FuncParameterAlias != "" {
			areParamsWithName = true
		}

		paramValueList = append(paramValueList, paramValue)
	}

	//validamos tipos y nombre
	i := 0
	for key := range function.FuncParameters {
		if function.FuncParameters[key].TSType != paramValueList[i].TSType {
			ctx.AddException(fmt.Sprintf("Llamada Funcion: El valor enviado es incorrecto, se esperaba: %s.", function.FuncParameters[key].TSType), I.Line, I.Position)
			return TSStructs.NewTNil()
		}

		if areParamsWithName && function.FuncParameters[key].FuncParameterAlias != paramValueList[i].FuncParameterAlias {
			ctx.AddException(fmt.Sprintf("Llamada Funcion: el alias especificado no coincide con el definido en la función: %s.", paramValueList[i].FuncParameterAlias), I.Line, I.Position)
			return TSStructs.NewTNil()
		}

		i++
	}

	//se usa en el orde que se declaran en la función

	ctx.AddScope()

	i = 0
	for key := range function.FuncParameters {
		paramValueList[i].IsBox = true
		//la variable se trata como una variable por valor o referencia
		if function.FuncParameters[key].IsConstant {
			copy := TSStructs.CopyTSValue(paramValueList[i])
			copy.IsConstant = true
			ctx.AddVariable(function.FuncParameters[key].BoxTag, &copy, I.Line, I.Position)
			i++
			continue
		}
		ctx.AddVariable(function.FuncParameters[key].BoxTag, paramValueList[i], I.Line, I.Position)
		i++
	}

	result := function.FuncValue.Interpret(ctx)

	ctx.RemoveScope()

	//valor devuelto no coincide
	if result.TSType != function.TSType {
		ctx.AddException(fmt.Sprintf("Llamada Funcion: El valor de retorna un valor incorrecto: %s, espperado %s.", result.TSType, function.TSType), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	return result

}

/******************************************************************************/
/*****************************  PARAMETER *************************************/
/******************************************************************************/
type IFCallParameter struct {
	TSStructs.TSExpression
	Alias           string
	Value           TSStructs.TSExpressioner
	PassByReference bool
}

func (I IFCallParameter) Interpret(ctx *TSStructs.TSContext) *TSStructs.TSValue {
	value := I.Value.Interpret(ctx)

	//el valor debe de ser una variable
	if I.PassByReference && !value.IsBox && !value.IsConstant {
		ctx.AddException(fmt.Sprintf("LLamadaFuncion: los parámetros por referencia deben ser variables."), I.Line, I.Position)
		return TSStructs.NewTNil()
	}

	if I.Alias != "" {
		value.FuncParameterAlias = I.Alias
	}

	//si es por referencia mandamos la variable
	if I.PassByReference {
		return value
	}

	//Creamos una copia del valor
	switch value.TSType {
	case TSStructs.STRING:
		return TSStructs.NewTString(value.Svalue)
	case TSStructs.INTEGER:
		return TSStructs.NewTInt(value.Ivalue)
	case TSStructs.FLOAT:
		return TSStructs.NewTFloat(value.Fvalue)
	case TSStructs.BOOL:
		return TSStructs.NewTBoolean(value.Bvalue)
	case TSStructs.CHARACTER:
		return TSStructs.NewTString(value.Svalue)
	case TSStructs.NIL:
		return TSStructs.NewTNil()
	}

	return TSStructs.NewTNil()
}

func NewIFCallParameter(Line int, Position int, alias string, value TSStructs.TSExpressioner, passByValue bool) *IFCallParameter {
	return &IFCallParameter{
		TSStructs.TSExpression{Line, Position, make([]string, 0)},
		alias, value, passByValue,
	}
}
