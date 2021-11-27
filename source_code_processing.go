package main


import "io/ioutil"


func getSourceCode(filepath string) string {
	sourceCode, err := ioutil.ReadFile(filepath)

	if err != nil {
		panic("can't read souce code from file " + filepath)
	}

	return string(sourceCode)
}


func isInSlice(input string, slice []string) bool {
	for _, str := range slice {
		if str == input {
			return true
		}
	}
	return false
}


func isKeyword(input string) bool {
	keywords := []string{"else", "if", "while", "return", "void", "for", "continue", "break", "const", "struct", "struct", "sizeof"}
	return isInSlice(input, keywords)
}


func isTypename(input string) bool {
	types := []string{"int", "void", "char"}
	return isInSlice(input, types)
}