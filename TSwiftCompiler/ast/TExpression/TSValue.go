package TExpression

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
}

func (T TSValue) ToString() string {
	var tvalue string
	if T.IsBox {
		tvalue = fmt.Sprintf("Tag: %s alue: ", T.BoxTag)
	}
	switch T.TSType {
	case INTEGER:
		tvalue += strconv.Itoa(T.Ivalue)
	case FLOAT:
		tvalue += strconv.FormatFloat(T.Fvalue, 'g', -1, 64)
	case STRING:
		tvalue += T.Svalue
	case BOOL:
		tvalue += strconv.FormatBool(T.Bvalue)
	case NIL:
		tvalue += "nil"
	case UNDEFINED:
		tvalue += "undefined"
	default:
		return "UNDEFINED"
	}
	return fmt.Sprintf("type: %s value: %s", T.TSType.String(), tvalue)
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

func NewTBox(tag string, content *TSValue, tstype TSPTYPES) *TSValue {
	return &TSValue{
		TSType: tstype,
		BoxTag: tag,
		IsBox:  true,
		IsNil:  true,
	}
}
