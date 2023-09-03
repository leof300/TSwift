package TSStructs

import (
	"TSwiftCompiler/ast/TExpression"
)

//if reflect.TypeOf(tvalue).String() != "TExpression.TSBox" {
//	return tvalue
//}

func Unbox(tvalue *TExpression.TSValue) *TExpression.TSValue {
	//
	//if !tvalue.IsBox {
	//	return tvalue
	//}
	//
	//return tvalue.BoxContent
	return tvalue
}
