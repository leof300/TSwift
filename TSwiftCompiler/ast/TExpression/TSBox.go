package TExpression

import (
	"fmt"
	"strconv"
)

type TSBox struct {
	Tag string

	Value TSValue
}

func (T TSBox) toString() string {
	var tvalue string
	switch T.Value.TSType {
	case INTEGER:
		tvalue = strconv.Itoa(T.Value.Ivalue)
	case FLOAT:
		tvalue = strconv.FormatFloat(T.Value.Fvalue, 'g', -1, 64)
	case STRING:
		tvalue = T.Value.Svalue
	case BOOL:
		tvalue = strconv.FormatBool(T.Value.Bvalue)
	case NIL:
		tvalue = "nil"

	}
	return fmt.Sprintf("tag: %s value: %s", T.Tag, tvalue)
}
