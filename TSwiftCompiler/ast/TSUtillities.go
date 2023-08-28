package ast

import (
	"TSwiftCompiler/ast/TExpression"
	"reflect"
)

/**
  @staticmethod
   def desempacar(tcaja):
       '''Obtener el valor de una posible caja'''
       if not isinstance(tcaja, TCaja):
           return tcaja  # es un TValor

       contenido = tcaja.tvalor
       if contenido is None:
           raise SemanticException("Esta variable que se desea acceder no fue declarada anteriormente.")

       return contenido  # si no es una caja es un
*/

func unbox(tvalue *TExpression.TValue) *TExpression.TValue {

	if reflect.TypeOf(tvalue).String() != "TExpression.TBox" {
		return tvalue
	}

	return tvalue
}
