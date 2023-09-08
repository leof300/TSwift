package TSStructs

import (
	"fmt"
)

//if reflect.TypeOf(tvalue).String() != "TExpression.TSBox" {
//	return tvalue
//}

func Unbox(tvalue *TSValue) *TSValue {
	//
	//if !tvalue.IsBox {
	//	return tvalue
	//}
	//
	//return tvalue.BoxContent
	return tvalue
}

func RelationalValidationsBefore(operation string, op1 *TSValue, op2 *TSValue) error {
	msg := operation + ": "
	if op1.IsNil || op2.IsNil {
		msg += "no se puede operar, valor nulo"
		return fmt.Errorf(msg)
	}

	op1Type := op1.TSType
	op2Type := op2.TSType

	//si los tipos que deseamos asignar son distintos tipos, posiblemente INT y FLOAT o FLOAT e INT
	if op1Type != op2Type {
		//si es una variable, se hace comparación de tipos
		//si es entero y su valor es 0 entonces se puede considerar como float
		if !op1.IsBox && op1Type == INTEGER && op1.Ivalue == 0 {
			op1Type = FLOAT
		}

		//si es una variable, se hace comparación de tipos
		//si es entero y su valor es 0 entonces se puede considerar como float
		if !op2.IsBox && op2Type == INTEGER && op2.Ivalue == 0 {
			op2Type = FLOAT
		}
	} //esto va a funcionar ya que el valor por defecto de Fvalue es 0.0 (solo no tomamos esto como error)

	//error tipos no iguales
	if op1Type != op2Type {
		msg += "tipos distintos no se pueden operar."
		return fmt.Errorf(msg)
	}

	return nil
}

func CopyTSValue(original *TSValue) TSValue {
	copia := original
	return *copia
}
