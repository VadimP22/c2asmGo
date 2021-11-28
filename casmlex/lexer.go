package casmlex


import (
	"../casmutility"
)


func Lex(sourceCode string, logger casmutility.Logger) []casmutility.Token{
	sourceCode = sourceCode + "\n"
	sourcesLen := len(sourceCode)
	i := 0

	for i < sourcesLen {
		sym := sourceCode[i]
		symstr := string(sym)

		if isBracket(sym) {
			logger.Println(symstr)
		}

		if isStringSeparator(sym) {
			logger.Println("string_separator")
		}

		if isOperator(sym) {
			complexOperator := isComplexOperator(symstr + string(sourceCode[i + 1]))

			if complexOperator {
				logger.Println("operator: " + symstr + string(sourceCode[i + 1]))
			} else {
				logger.Println("operator: " + symstr)
			}

			i = i + 1
		}

		if isNumber(sym) {
			j := i
			number := ""

			for isNumber(sourceCode[j]) {
				number = number + string(sourceCode[j])
				j = j + 1
			}

			logger.Println("number: " + number)

			i = j - 1
		}

		if isChar(sym) {

			j := i
			identifier := ""

			for (isChar(sourceCode[j]) || isNumber(sourceCode[j])) {

				identifier = identifier + string(sourceCode[j])
				j = j + 1
			}

			if identifier != "\r" {
				logger.Println("identifier: " + identifier)
			}

			i = j - 1
		}

		i = i + 1

	}


	return nil
}