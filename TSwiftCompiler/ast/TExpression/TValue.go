package TExpression

import (
	"fmt"
	"strconv"
)

type TValue interface {
	toString() string
}

type TInt struct {
	Value int
}

func NewTInt(value int) *TInt {
	return &TInt{Value: value}
}

func (T TInt) toString() string {
	return strconv.Itoa(T.Value)
}

type TFloat struct {
	Value float64
}

func (T TFloat) toString() string {
	return fmt.Sprintf("%v", T.Value)
}

func NewTFloat(value float64) *TFloat {
	return &TFloat{Value: value}
}

type TString struct {
	Value string
}

func (T TString) toString() string {
	return T.Value
}

func NewTString(value string) *TString {
	return &TString{Value: value}
}

type TBox struct {
	Tag   string
	Value TValue
}

func (T TBox) toString() string {
	return T.Value.toString()
}

func NewTBox(tag string) *TBox {
	return &TBox{Tag: tag}
}

type TBoolean struct {
	Value bool
}

func (T TBoolean) toString() string {
	return fmt.Sprintf("%v", T.Value)
}

func NewTBoolean(value bool) *TBoolean {
	return &TBoolean{Value: value}
}

type TNil struct {
	Value bool
}

func (T TNil) toString() string {
	return fmt.Sprintf("%v", T.Value)
}
