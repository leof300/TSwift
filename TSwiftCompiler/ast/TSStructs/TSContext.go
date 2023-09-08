package TSStructs

import (
	"TSwiftCompiler/ast/Exceptions"
	"fmt"
	"strconv"
)

type TSContext struct {
	Log        []string
	Console    []string                    `json:"Console"`
	Exceptions []*TSExceptions.TSException `json:"Exceptions"`
	Scope      *TSScope                    `json:"-"`
	TDS        []TDSymbol
}

func NewContext() *TSContext {
	c := &TSContext{
		Log:        make([]string, 0),
		Exceptions: make([]*TSExceptions.TSException, 0),
		Scope:      NewTSScope(nil),
		Console:    make([]string, 0),
		TDS:        make([]TDSymbol, 0),
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
const FUNCTION_PREFIX string = "F%"

/*
*
Agregar una funcion al contexto
*/
func (c *TSContext) AddFunction(key string, value *TSValue, line int, position int) bool {
	//TODO AGREGAR A LA TABLA DE SIMBOLOS AQUI, reporte
	//agregamos el prefijo F%
	ok := c.Scope.AddSymbol(FUNCTION_PREFIX+key, value, line, position)
	return ok
}

/**
*Buscar una variable
 */
func (c *TSContext) GetVariable(key string) *TSValue {
	variable := c.Scope.GetSymbolValue(key)
	if variable == nil {
		scopeAux := c.Scope.ParentScope
		for scopeAux != nil && !variable.IsFunction && !variable.IsArray {
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
	key = FUNCTION_PREFIX + key
	function := c.Scope.GetSymbolValue(key)
	if function == nil {
		scopeAux := c.Scope.ParentScope
		for scopeAux != nil {
			function = scopeAux.GetSymbolValue(key)
			if function != nil && function.IsFunction {
				return function
			}
			scopeAux = scopeAux.ParentScope
		}
	} else {
		return function
	}
	return nil

}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type TDSymbol struct {
	Scope       string
	Key         string
	Type        TSPTYPES
	Value       string
	Box         bool
	Constant    bool
	Array       bool
	Function    bool
	FParameters []TDSymbol
	Alias       string
	Line        int
	Position    int
}

func (c *TSContext) FillTDS() {
	for _, s := range c.Scope.variables {
		tds := convertTSValueTpTDSymbol(s.Value)
		tds.Key = s.Key
		tds.Line = s.Line
		tds.Position = s.Position
		tds.FParameters = make([]TDSymbol, 0)
		for _, p := range s.Value.FuncParameters {
			par := convertTSValueTpTDSymbol(p)
			tds.FParameters = append(tds.FParameters, par)
		}

		c.TDS = append(c.TDS, tds)
	}
}

func convertTSValueTpTDSymbol(s *TSValue) TDSymbol {
	var tds TDSymbol
	tds.Type = s.TSType
	switch tds.Type {
	case NIL:
		tds.Value = "nil"
	case INTEGER:
		tds.Value = strconv.Itoa(s.Ivalue)
	case FLOAT:
		tds.Value = strconv.FormatFloat(s.Fvalue, 'f', -1, 64)
	case BOOL:
		tds.Value = strconv.FormatBool(s.Bvalue)
	case STRING:
		tds.Value = s.Svalue
	case CHARACTER:
		tds.Value = s.Svalue
	}
	tds.Box = s.IsBox
	tds.Constant = s.IsConstant
	tds.Array = s.IsArray
	tds.Function = s.IsFunction
	tds.Alias = s.FuncParameterAlias
	return tds
}
