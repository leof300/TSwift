package main

import (
	"TSwiftCompiler/ast/TSStructs"
	"TSwiftCompiler/interpreter"
	TSVisitor "TSwiftCompiler/visitor"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
)

func main() {
	fs, err := antlr.NewFileStream("./grammar/input5.swift")
	if err != nil {
		log.Fatal(err)
	}
	executeInterpreter(fs)
}

func executeInterpreter(input antlr.CharStream) bool {
	lexer := TSVisitor.NewTSParser_rulesLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := TSVisitor.NewTSParser_rulesParser(tokens)
	tree := parser.Start_()

	tsvisitor := interpreter.NewTSwiftVisitor()
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

	ctx = nil

	return true
}
