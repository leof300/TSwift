package TSExceptions

import "TSwiftCompiler/ast"

// intefaz en Go
type TSException struct {
	ast.TSExpression
	message  string
	value    string
	line     int
	position int
}

type TSSemanticE struct {
	TSException
}

func NewTSSemanticE(message string, line int, position int) *TSSemanticE {
	return &TSSemanticE{TSException{line: line, position: position, message: message}}
}

type TSSintacticE struct {
	TSException
}

type TSLexicalE struct {
	TSException
}
