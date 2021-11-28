package casmlex


import (
	"testing"
	"../casmutility"
)

func BenchmarkLexer(b *testing.B) {
	testString := ""
	testSubstring := "int main()\n { int a = 33;\n return sum(a, 66);\n }\n"
	countOfSubstrings := 10000
	logger := casmutility.NewVoidLogger()

	for i := 0; i < countOfSubstrings; i++ {
		testString = testString + testSubstring
	}


	b.ResetTimer()
	Lex(testString, logger)
	b.StopTimer()

}