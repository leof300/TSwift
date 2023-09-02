package TSStructs

import (
	"TSwiftCompiler/ast/Exceptions"
	"TSwiftCompiler/ast/TExpression"
)

type TSContext struct {
	Console    string
	Exceptions []TSExceptions.TSException
	Memory     *TSMemory
}

func NewContext() *TSContext {
	c := &TSContext{
		Console:    "",
		Exceptions: make([]TSExceptions.TSException, 0),
		Memory:     NewTSMemory(nil),
	}
	return c
}

/*
* Agregar un error o excepción
 */
func (c TSContext) AddException(s string, line int, position int) bool {
	exception := TSExceptions.NewTSException(s, line, position)
	c.Exceptions = append(c.Exceptions, *exception)
	return true
}

/*
*
Agregar una variable al contexto
*/
func (c TSContext) AddVariable(key string, value *TExpression.TSValue, line int, position int) bool {
	//TODO AGREGAR A LA TABLA DE SIMBOLOS AQUI

	ok := c.Memory.AddSymbol(key, value, line, position)

	return ok
}

/**
*Buscar una variable
 */
func (c TSContext) GetVariable(key string) *TExpression.TSValue {
	variable := c.Memory.GetSymbolValue(key)
	if variable == nil {
		auxMem := c.Memory.ParentMemory
		for auxMem != nil {
			variable = c.Memory.GetSymbolValue(key)
			if variable != nil {
				return variable
			}
			auxMem = auxMem.ParentMemory
		}
	} else {
		return variable
	}
	return nil
}
