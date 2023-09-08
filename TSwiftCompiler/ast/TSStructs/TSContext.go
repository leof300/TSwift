package TSStructs

import (
	"TSwiftCompiler/ast/Exceptions"
	"fmt"
)

type TSContext struct {
	Log        string
	Console    []string
	Exceptions []*TSExceptions.TSException
	Scope      *TSScope
}

func NewContext() *TSContext {
	c := &TSContext{
		Log:        "",
		Exceptions: make([]*TSExceptions.TSException, 0),
		Scope:      NewTSScope(nil),
		Console:    make([]string, 0),
	}
	return c
}

/*
* Agregar un error o excepción
 */
func (c *TSContext) AddException(s string, line int, position int) bool {
	c.Exceptions = append(c.Exceptions, TSExceptions.NewTSException(s, line, position))
	return true
}

/*
* crea un nuevo ámbito
 */
func (c *TSContext) AddScope() {
	//la referencia del padre la tiene el scope original
	c.Scope = NewTSScope(c.Scope)
}

/*
* elimina el ámbito actual
 */
func (c *TSContext) RemoveScope() {
	if c.Scope != nil {
		c.Scope = c.Scope.ParentScope
	} else {
		fmt.Print("erro con scopes")
	}
}

/*
*
Agregar una variable al contexto
*/
func (c *TSContext) AddConsole(msg string) bool {
	c.Console = append(c.Console, msg)
	return true
}

/*
*
Agregar una variable al contexto
*/
func (c *TSContext) AddVariable(key string, value *TSValue, line int, position int) bool {
	//TODO AGREGAR A LA TABLA DE SIMBOLOS AQUI, reporte
	ok := c.Scope.AddSymbol(key, value, line, position)
	return ok
}

// esta constante es para el nombre de funciones
const functionPrefix string = "F%"

/*
*
Agregar una funcion al contexto
*/
func (c *TSContext) AddFunction(key string, value *TSValue, line int, position int) bool {
	//TODO AGREGAR A LA TABLA DE SIMBOLOS AQUI, reporte
	//agregamos el prefijo F%
	ok := c.Scope.AddSymbol(functionPrefix+key, value, line, position)
	return ok
}

/**
*Buscar una variable
 */
func (c *TSContext) GetVariable(key string) *TSValue {
	variable := c.Scope.GetSymbolValue(key)
	if variable == nil {
		scopeAux := c.Scope.ParentScope
		for scopeAux != nil {
			variable = scopeAux.GetSymbolValue(key)
			if variable != nil {
				return variable
			}
			scopeAux = scopeAux.ParentScope
		}
	} else {
		return variable
	}
	return nil
}

/**
*Buscar una variable
 */
func (c *TSContext) GetFunction(key string) *TSValue {
	//eliminamos el prefijo F_
	//key = key[2:]

	//key = strings.TrimLeft(key, functionPrefix)

	return c.GetVariable(functionPrefix + key)
}
