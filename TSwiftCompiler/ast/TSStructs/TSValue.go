package TSStructs

import (
	"fmt"
	"strconv"
)

type TSPTYPES uint

const (
	NIL       TSPTYPES = 0
	BOOL      TSPTYPES = 1
	INTEGER   TSPTYPES = 2
	FLOAT     TSPTYPES = 3
	STRING    TSPTYPES = 4
	UNDEFINED TSPTYPES = 5
	BOX       TSPTYPES = 6
	CHARACTER TSPTYPES = 7
)

func (t TSPTYPES) String() string {
	switch t {
	case NIL:
		return "NIL"
	case INTEGER:
		return "Integer"
	case FLOAT:
		return "Float"
	case STRING:
		return "String"
	case BOOL:
		return "Bool"
	case UNDEFINED:
		return "Undefined"
	case BOX:
		return "Box"
	case CHARACTER:
		return "Character"
	default:
		return "Undefined"
	}
}

type TSValuer interface {
	ToString()
}

type TSValue struct {
	Fvalue float64
	Svalue string
	Bvalue bool
	Ivalue int
	TSType TSPTYPES

	IsNil bool

	IsBox  bool
	BoxTag string

	IsReturn   bool
	IsBreak    bool
	IsContinue bool

	IsConstant bool

	IsFunction         bool
	FuncValue          TSExpressioner
	FuncParameters     map[string]*TSValue //apuntado a los parametros, o variables creadas en el ámbito
	FuncParameterAlias string
}

func (T TSValue) ToString() string {
	switch T.TSType {
	case INTEGER:
		return strconv.Itoa(T.Ivalue)
	case FLOAT:
		return strconv.FormatFloat(T.Fvalue, 'g', -1, 64)
	case STRING:
		return T.Svalue
	case CHARACTER:
		return T.Svalue
	case BOOL:
		return strconv.FormatBool(T.Bvalue)
	case NIL:
		return "nil"
	default:
		return ""
	}
}

func (T TSValue) ToComplexString() string {
	svalue := ""

	if T.IsBox {
		svalue = fmt.Sprintf("{Tag: %s, Type: %s,Value: %s}", T.BoxTag, T.TSType.String(), T.ToString())
	}

	return svalue
}

func (T TSValue) IsUndefined() bool {
	if T.TSType == NIL {
		return true
	}
	return false
}

func NewTInt(value int) *TSValue {
	return &TSValue{
		Ivalue: value,
		TSType: INTEGER,
	}
}

func NewTFloat(value float64) *TSValue {
	return &TSValue{
		Fvalue: value,
		TSType: FLOAT,
	}
}

func NewTString(value string) *TSValue {
	return &TSValue{
		Svalue: value,
		TSType: STRING,
	}
}

func NewTBoolean(value bool) *TSValue {
	return &TSValue{
		Bvalue: value,
		TSType: BOOL,
	}
}

func NewTNil() *TSValue {
	return &TSValue{
		TSType: NIL,
		IsNil:  true,
	}
}

func NewTUndefined() *TSValue {
	return &TSValue{
		TSType: UNDEFINED,
		IsNil:  true,
	}
}

func NewTBox(tag string, tstype TSPTYPES) *TSValue {
	return &TSValue{
		TSType: tstype,
		BoxTag: tag,
		IsBox:  true,
		IsNil:  true,
	}
}

func NewFunction(tag string, tstype TSPTYPES) *TSValue {
	return &TSValue{
		TSType:         tstype,
		BoxTag:         tag,
		IsBox:          true,
		IsNil:          true,
		IsFunction:     true,
		FuncParameters: make(map[string]*TSValue),
	}
}

func NewConstant(tag string, tstype TSPTYPES) *TSValue {
	return &TSValue{
		TSType:     tstype,
		BoxTag:     tag,
		IsConstant: true,
		IsNil:      true,
	}
}
