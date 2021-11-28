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


func TestLex(t *testing.T) {
	testString := "int main() {\n"

	var expected []casmutility.Token
	expected = append(expected, casmutility.NewToken("typename", "int"))
	expected = append(expected, casmutility.NewToken("identifier", "main"))
	expected = append(expected, casmutility.NewToken("bracket_open", "("))
	expected = append(expected, casmutility.NewToken("bracket_close", ")"))
	expected = append(expected, casmutility.NewToken("bracket_open", "{"))

	tokens := Lex(testString, casmutility.NewVoidLogger())
	
	for i := range tokens {
		if tokens[i] != expected[i] {
			t.Error("for i=",i ,"expected", expected[i], "got", tokens[i])
		}
	}
}