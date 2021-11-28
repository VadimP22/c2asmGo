package main


import (
	"../casmlex"
	"../casmutility"
)


func main() {
	sourcCode := casmlex.GetSourceCode("input.txt")
	casmlex.Lex(sourcCode, casmutility.NewConsoleLogger("LEXER"))
	
}