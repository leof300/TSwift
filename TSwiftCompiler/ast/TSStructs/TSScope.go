package TSStructs

type Symbol struct {
	Key      string
	Value    *TSValue
	Line     int
	Position int

	//puede tener un tipo asociado
}

func NewSymbol(key string, value *TSValue, line int, position int) *Symbol {
	return &Symbol{
		key,
		value,
		line,
		position,
	}
}

type TSScope struct {
	variables   map[string]*Symbol //map / diccionario
	ParentScope *TSScope           //anterior
}

func NewTSScope(parent *TSScope) *TSScope {
	return &TSScope{
		variables:   make(map[string]*Symbol),
		ParentScope: parent,
	}
}

func (T *TSScope) AddSymbol(key string, value *TSValue, line int, position int) bool {
	_, ok := T.variables[key]
	if ok {
		//ya existe la variable
		return false
	}
	T.variables[key] = NewSymbol(key, value, line, position)
	return true
}

func (T *TSScope) UpdateSymbol(key string, value *TSValue) bool {
	symbol, ok := T.variables[key]
	if !ok {
		// no existe la variable
		return false
	}
	symbol.Value = value
	return true
}

func (T *TSScope) Exist(key string) bool {
	_, ok := T.variables[key]
	return ok
}

func (T *TSScope) GetSymbolValue(key string) *TSValue {
	symbol, ok := T.variables[key]
	if ok {
		return symbol.Value
	}
	return nil
}
