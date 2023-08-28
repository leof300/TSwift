package ast

type TSContext struct {
	console    string
	exceptions []string
}

func NewContext() *TSContext {
	c := &TSContext{
		console:    "",
		exceptions: make([]string, 0),
	}
	return c
}
