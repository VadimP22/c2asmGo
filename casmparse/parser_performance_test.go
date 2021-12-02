package casmparse


import (
	"testing"
	"../casmutility"
	"../casmlex"
)


func BenchmarkParser(b *testing.B) {
	testString := ""
	testSubstring := "int main()\n {\n int a = 33*(44 + 2);\n }\n"
	countOfSubstrings := 5000
	logger := casmutility.NewVoidLogger()

	for i := 0; i < countOfSubstrings; i++ {
		testString = testString + testSubstring
	}


	tokens := casmlex.Lex(testString, logger)

	
	b.ResetTimer()
	functionsDefinitions, _ := FindFunctionsDefinitions(tokens, logger)
	root := ParseFunctions(functionsDefinitions, tokens, logger)
	b.StopTimer()

	_ = root

}