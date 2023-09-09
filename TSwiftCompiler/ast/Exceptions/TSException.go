package TSExceptions

import "fmt"

// intefaz en Go
type TSException struct {
	Message  string    `json:"Message"`
	EType    TSEXTYPES `json:"Type"`
	Line     int       `json:"Line"`
	Position int       `json:"Position"`
}

func NewTSException(message string, line int, position int) *TSException {
	return &TSException{Line: line, Position: position, Message: message, EType: SEMANTIC}
}

func (T TSException) ToString() string {
	return fmt.Sprintf("{Tipo: %s Message: \"%s\", Line: %d, Position: %d}", T.EType, T.Message, T.Line, T.Position)
}

type TSEXTYPES uint

const (
	LEXICAL  TSEXTYPES = 0
	SYNTAX   TSEXTYPES = 1
	SEMANTIC TSEXTYPES = 2
)

func (t TSEXTYPES) String() string {
	switch t {
	case LEXICAL:
		return "Lexical"
	case SYNTAX:
		return "Syntax"
	case SEMANTIC:
		return "Semantic"
	}
	return "Unknown"
}
