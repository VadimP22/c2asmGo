package main

import "fmt"

type token struct {
	tokenType  string
	tokenValue string
}

func (t *token) getType() string {
	return t.tokenType
}

func (t *token) getValue() string {
	return t.tokenValue
}

func newToken(tokenType, tokenValue string) token {
	var tok token
	tok.tokenType = tokenType
	tok.tokenValue = tokenValue
	return tok
}

func lex(sourceCode string) {
	sourceCode = sourceCode + "\n"
	sourcesLen := len(sourceCode)
	i := 0

	for i < sourcesLen {
		sym := sourceCode[i]

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

			i = j
		}


		i = i + 1

	}
}