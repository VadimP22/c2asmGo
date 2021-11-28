package main

import (
	"../casmlex"
)

func main() {
	sourcCode := casmlex.GetSourceCode("input.txt")
	casmlex.Lex(sourcCode)
	
}