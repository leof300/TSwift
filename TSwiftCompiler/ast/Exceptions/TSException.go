package TSExceptions

import "fmt"

// intefaz en Go
type TSException struct {
	Message  string
	value    string
	Line     int
	position int
}

func NewTSException(message string, line int, position int) *TSException {
	return &TSException{Line: line, position: position, Message: message}
}

func (T TSException) ToString() string {
	return fmt.Sprintf("{Message: \"%s\", Line: %d, Position: %d}", T.Message, T.Line, T.position)
}
