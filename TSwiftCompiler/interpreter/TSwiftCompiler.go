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
	ctx := TSStructs.NewContext()

	lexer := TSVisitor.NewTSParser_rulesLexer(input)

	/******************************************************************/

	//lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&LexicalErrorListener{ctx: ctx})
	//
	/******************************************************************/
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := TSVisitor.NewTSParser_rulesParser(tokens)

	/******************************************************************/
	//parser.RemoveErrorListeners()
	parser.AddErrorListener(&SyntaxErrorListener{ctx: ctx})
	/******************************************************************/

	tree := parser.Start_()

	tsvisitor := NewTSwiftVisitor()

	root := tsvisitor.Visit(tree).(TSStructs.TSExpressioner)

	result := root.Interpret(ctx)

	if result == nil {
		return nil
	}

	fmt.Println("***************** Console:    **************************")
	for _, msg := range ctx.Console {
		fmt.Println(msg)
	}
	fmt.Println("***************** Exceptions: **************************")

	for _, msg := range ctx.Exceptions {
		fmt.Println(msg.ToString())
	}

	fmt.Println("********************************************************")

	ctx.FillTDS()

	return ctx
}

type LexicalErrorListener struct {
	antlr.ErrorListener
	ctx *TSStructs.TSContext
}

func (l *LexicalErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.ctx.AddLexicalException(msg, line, column)

}

type SyntaxErrorListener struct {
	antlr.ErrorListener
	ctx *TSStructs.TSContext
}

func (l *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.ctx.AddSyntaxException(msg, line, column)

}

//type CustomErrorStrategy struct {
//	antlr.DefaultErrorStrategy // Puedes extender DefaultErrorStrategy
//}
//
//func (e *CustomErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
//	// Personalizar la recuperación de errores aquí
//	// Por ejemplo, puedes implementar una lógica de recuperación personalizada.
//
//
//}
