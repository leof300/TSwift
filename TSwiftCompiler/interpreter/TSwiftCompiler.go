package interpreter

import (
	"TSwiftCompiler/ast/TSStructs"
	TSVisitor "TSwiftCompiler/visitor"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
)

func ProcessInputText(input string) *TSStructs.TSContext {

	stream := antlr.NewInputStream(input)
	return ExecuteInterpreter(stream)
}

func ProcessFileText(fileName string) *TSStructs.TSContext {
	stream, err := antlr.NewFileStream(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return ExecuteInterpreter(stream)
}

func ExecuteInterpreter(input antlr.CharStream) *TSStructs.TSContext {
	lexer := TSVisitor.NewTSParser_rulesLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := TSVisitor.NewTSParser_rulesParser(tokens)
	tree := parser.Start_()

	tsvisitor := NewTSwiftVisitor()
	root := tsvisitor.Visit(tree).(TSStructs.TSExpressioner)

	ctx := TSStructs.NewContext()

	result := root.Interpret(ctx)

	fmt.Println("Console:")
	for _, msg := range ctx.Console {
		fmt.Println(msg)
	}
	fmt.Println("Exceptions: ")

	for _, msg := range ctx.Exceptions {
		fmt.Println(msg.ToString())
	}

	fmt.Println(result)

	ctx.FillTDS()

	return ctx
}
