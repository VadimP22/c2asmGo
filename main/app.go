package main

import (
	"os"

	"../casmlex"
	"../casmparse"
	"../casmutility"
)


func main() {
	argsParserLogger := casmutility.NewConsoleLogger("ARGSPARSER")
	mainLogger := casmutility.NewConsoleLogger("MAIN")
	lexerLogger := casmutility.NewConsoleLogger("LEXER")
	parserLogger := casmutility.NewConsoleLogger("PARSER")
	errorLogger := casmutility.NewConsoleLogger("ERROR")

	mainLogger.Println("Parsing input args")
	fileName, args := casmutility.ParseArgs(os.Args, argsParserLogger)

	sourceCode, err := casmutility.GetSourceCode(fileName, mainLogger)
	if err != nil {
		errorLogger.Println("can't open '" + fileName + "':")
		errorLogger.Println(err.Error())
		os.Exit(1)
	}

	mainLogger.Println("Lexing started")
	tokens := casmlex.Lex(sourceCode, lexerLogger)

	mainLogger.Println("Finding functions definitions")
	functionsDefinitions, err := casmparse.FindFunctionsDefinitions(tokens, parserLogger)
	if err != nil {
		errorLogger.Println(err.Error())
		os.Exit(1)
	}

	err = casmutility.CheckEntryFunction("main", functionsDefinitions)
	if err != nil {
		errorLogger.Println(err.Error())
		os.Exit(1)
	}
	
	root := casmparse.ParseFunctions(functionsDefinitions, tokens, parserLogger)

	_ = root
	_ = fileName
	_ = functionsDefinitions
	_ = args
}