package casmlex

import "fmt"

func Lex(sourceCode string) {
	sourceCode = sourceCode + "\n"
	sourcesLen := len(sourceCode)
	i := 0

	for i < sourcesLen {
		sym := sourceCode[i]
		symstr := string(sym)

		if isBracket(sym) {
			fmt.Println(symstr)
		}

		if isStringSeparator(sym) {
			fmt.Println("STR_SEP")
		}

		if isOperator(sym) {
			complexOperator := isComplexOperator(symstr + string(sourceCode[i + 1]))

			if complexOperator {
				fmt.Println("complex:", symstr + string(sourceCode[i + 1]))
			} else {
				fmt.Println("simple:", symstr)
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

			fmt.Println(number)

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
				fmt.Println(identifier)
			}

			i = j - 1
		}

		i = i + 1

	}
}