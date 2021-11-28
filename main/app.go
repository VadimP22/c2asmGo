package main

import (
	"../casmlex"
	"../casmutility"
)


func main() {
	mainLogger := casmutility.NewConsoleLogger("MAIN")
	lexerLogger := casmutility.NewConsoleLogger("LEXER")
	errorLogger := casmutility.NewConsoleLogger("ERROR")

	fileName := "input.txt"
	mainLogger.Println("input file: " + fileName)

	sourceCode, err := casmutility.GetSourceCode(fileName)
	if err != nil {
		errorLogger.Println("can't open '" + fileName + "':")
		errorLogger.Println(err.Error())
	}

	tokens := casmlex.Lex(sourceCode, lexerLogger)
	_ = tokens
}