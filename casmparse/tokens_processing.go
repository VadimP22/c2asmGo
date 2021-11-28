package casmparse


import (
	"errors"

	"../casmutility"
)


func FindFunctionsDefinitions(tokens []casmutility.Token, logger casmutility.Logger) ([]casmutility.FunctionDefinition, error) {
	
	var result []casmutility.FunctionDefinition

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token.GetType() == "identifier" {
			nextToken := tokens[i + 1]
			prevToken := tokens[i - 1]
			if nextToken.GetType() == "bracket_open" && nextToken.GetValue() == "(" {
				if prevToken.GetType() == "typename" {
					logger.Println("definition found: " + token.GetValue())
				}
			}
		}

	}

	lenght := len(result)

	if lenght == 0 {
		return nil, errors.New("there's no functions in your file")
	}

	return result, nil
}