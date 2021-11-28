package casmlex


import "io/ioutil"


func GetSourceCode(filepath string) string {
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


func isByteInSlice(input byte, slice []byte) bool {
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


func isChar(input byte) bool {
	numbers := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	delimiters := []byte{'\t','\n',',',';','(',')','{','}','[',']', ' '}
	operators := []byte{'+', '-', '*', '/', '=', '<', '>', '&'}
	complexOperatorsStart := []byte{'!'}
	//complexOperators := []string{"<=", ">=", "==", "!="}

	if isByteInSlice(input, numbers) {
		return false
	}

	if isByteInSlice(input, delimiters) {
		return false
	}

	if isByteInSlice(input, operators) {
		return false
	}

	if isByteInSlice(input, complexOperatorsStart) {
		return false
	}

	return true
}


func isNumber(input byte) bool {
	numbers := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	return isByteInSlice(input, numbers)
}


func isBracket(input byte) bool {
	brackets := []byte{'{', '}', '[', ']', '(', ')'}

	return isByteInSlice(input, brackets)
}


func isOperator(input byte) bool {
	operators := []byte{'+', '-', '*', '/', '=', '<', '>', '&', '!'}

	return isByteInSlice(input, operators)
}


func isComplexOperator(input string) bool{
	complexOperators := []string{">=", "<=", "!=", "=="}

	return isInSlice(input, complexOperators)
}


func isStringSeparator(input byte) bool {
	return input == ';'
}