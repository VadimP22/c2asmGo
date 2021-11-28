package casmparse

import (
	"errors"
	"strings"

	"../casmutility"
)


func FindFunctionsDefinitions(tokens []casmutility.Token, logger casmutility.Logger) ([]casmutility.FunctionDefinition, error) {
	
	var result []casmutility.FunctionDefinition

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		prefixString := "    "

		if token.GetType() == "identifier" {
			nextToken := tokens[i + 1]
			prevToken := tokens[i - 1]
			if nextToken.GetType() == "bracket_open" && nextToken.GetValue() == "(" {
				if prevToken.GetType() == "typename" {
					logger.Println("definition found: " + token.GetValue())

					k, argTypes, argNames := findFunctionArgs(tokens, i)
					returnType := prevToken.GetValue()

					logger.Println(prefixString + "argTypes: [" + strings.Join(argTypes, ", ") + "]")
					logger.Println(prefixString + "argNames: [" + strings.Join(argNames, ", ") + "]")
					logger.Println(prefixString + "retnType: [" + returnType + "]")

					result = append(result, casmutility.NewCFunctionDefinition(i, token.GetValue(), argTypes, argNames, returnType))

					i = k
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


func findFunctionArgs(tokens []casmutility.Token, j int) (int, []string, []string) {
	i := j + 2
	var (
		argTypes []string
		argNames []string
	)

	for tokens[i].GetType() != "bracket_close" {
		token := tokens[i]
		if token.GetType() == "typename" {
			argTypes = append(argTypes, token.GetValue())
		}

		if token.GetType() == "identifier" {
			argNames = append(argNames, token.GetValue())
		}

		i = i + 1
	}

	return i, argTypes, argNames
}