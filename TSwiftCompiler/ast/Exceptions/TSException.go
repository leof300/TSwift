package TSExceptions

// intefaz en Go
type TSException struct {
	message  string
	value    string
	line     int
	position int
}

func NewTSException(message string, line int, position int) *TSException {
	return &TSException{line: line, position: position, message: message}
}

//type TSSemanticE struct {
//	TSException
//}
//
//type TSSintacticE struct {
//	TSException
//}
//
//type TSLexicalE struct {
//	TSException
//}
