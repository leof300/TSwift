package TSStructs

import "TSwiftCompiler/ast/TExpression"

type Symbol struct {
	Key      string
	Value    *TExpression.TSValue
	Line     int
	Position int

	//puede tener un tipo asociado
}

func NewSymbol(key string, value *TExpression.TSValue, line int, position int) *Symbol {
	return &Symbol{
		key,
		value,
		line,
		position,
	}
}

type TSMemory struct {
	variables    map[string]*Symbol //map / diccionario
	ParentMemory *TSMemory          //anterior
}

func NewTSMemory(parent *TSMemory) *TSMemory {
	return &TSMemory{
		variables:    make(map[string]*Symbol),
		ParentMemory: parent,
	}
}

func (T *TSMemory) AddSymbol(key string, value *TExpression.TSValue, line int, position int) bool {
	_, ok := T.variables[key]
	if ok {
		//ya existe la variable
		return false
	}
	T.variables[key] = NewSymbol(key, value, line, position)
	return true
}

func (T *TSMemory) UpdateSymbol(key string, value *TExpression.TSValue) bool {
	symbol, ok := T.variables[key]
	if !ok {
		// no existe la variable
		return false
	}
	symbol.Value = value
	return true
}

func (T *TSMemory) Exist(key string) bool {
	_, ok := T.variables[key]
	return ok
}

func (T *TSMemory) GetSymbolValue(key string) *TExpression.TSValue {
	symbol, ok := T.variables[key]
	if ok {
		return symbol.Value
	}
	return nil
}
