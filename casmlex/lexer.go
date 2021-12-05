package casmlex


import (
	"../casmutility"
)


func Lex(sourceCode string, logger casmutility.Logger) []casmutility.Token{
	var tokens []casmutility.Token

	sourceCode = sourceCode + "\n"
	sourcesLen := len(sourceCode)
	i := 0

	for i < sourcesLen {
		sym := sourceCode[i]
		symstr := string(sym)

		if isBracket(sym) {
			if isOpenBracket(sym) {
				tokens = append(tokens, casmutility.NewToken("bracket_open", string(sym)))
				logger.Println("bracket_open: " + string(sym))
			} else {
				tokens = append(tokens, casmutility.NewToken("bracket_close", string(sym)))
				logger.Println("bracket_close: " + string(sym))
			}
		}

		if isComma(sym) {
			tokens = append(tokens, casmutility.NewToken("comma", ","))
			logger.Println("comma: ,")
		}

		if isStringSeparator(sym) {
			logger.Println("string_separator")
			tokens = append(tokens, casmutility.NewToken("string_separator", ""))
		}

		if isOperator(sym) {
			complexOperator := isComplexOperator(symstr + string(sourceCode[i + 1]))

			if complexOperator {
				operator := symstr + string(sourceCode[i + 1])

				logger.Println("operator: " + operator)
				tokens = append(tokens, casmutility.NewToken("operator", operator))
				i = i + 1
			} else {
				logger.Println("operator: " + symstr)
				tokens = append(tokens, casmutility.NewToken("operator", symstr))
			}

			
		}

		if isNumber(sym) {
			j := i
			number := ""

			for isNumber(sourceCode[j]) {
				number = number + string(sourceCode[j])
				j = j + 1
			}

			logger.Println("number: " + number)
			tokens = append(tokens, casmutility.NewToken("number", number))

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
				if isKeyword(identifier) {
					logger.Println("keyword: " + identifier)
					tokens = append(tokens, casmutility.NewToken("keyword", identifier))
				} else if isTypename(identifier) {
					logger.Println("typename: " + identifier)
					tokens = append(tokens, casmutility.NewToken("typename", identifier))
				} else {
					logger.Println("identifier: " + identifier)
					tokens = append(tokens, casmutility.NewToken("identifier", identifier))
				}
			}

			i = j - 1
		}

		i = i + 1

	}


	return tokens
}