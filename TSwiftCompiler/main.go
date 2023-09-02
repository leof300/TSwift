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
	fs, err := antlr.NewFileStream("./grammar/input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	executeInterpreter(fs)
	//lexer := TSVisitor.NewTSParser_rulesLexer(fs)
	//tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	//parser := TSVisitor.NewTSParser_rulesParser(tokens)
	//tree := parser.Start_()
	//
	//visitor := interpreter.NewTSwiftVisitor()
	//
	//root := visitor.Visit(tree).(TSStructs.TSExpressioner)
	////root := visitor.Visit(tree)
	//a := "ds"
	//fmt.Print("Ejemplo" + a)
	//fmt.Print(root)
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

	fmt.Print(ctx.Console)
	fmt.Print(result)

	return true
}
