package casmparse

import (
	"errors"
	"../casmutility"
)


func getFunctionTokens(name string, funcs []casmutility.FunctionDefinition, tokens []casmutility.Token) []casmutility.Token{
	var fun casmutility.FunctionDefinition
	var newTokens []casmutility.Token 

	for _, f := range funcs {
		if f.GetName() == name {
			fun = f
		}
	}

	i := fun.GetTokenNumber()
	i = findFirstBracket(i, tokens)
	isRunning := true
	openBrackets := 0
	closeBrackets := 0
	for isRunning {
		if tokens[i].GetType() == "bracket_open" {
			openBrackets = openBrackets + 1
		}

		if tokens[i].GetType() == "bracket_close" {
			closeBrackets = closeBrackets + 1
		}

		if openBrackets == closeBrackets {
			isRunning = false
		}

		newTokens = append(newTokens, tokens[i])

		i = i + 1
	}

	return newTokens
}


func findFirstBracket(i int, tokens []casmutility.Token) int{
	j := i
	for {
		if tokens[j].GetValue() == "{" {
			break
		}
		
		j = j + 1
	}
	
	return j
}


func ParseFunctions(funcs []casmutility.FunctionDefinition, tokens []casmutility.Token, logger casmutility.Logger) *Node{
	root := NewNode("root", "root")

	for _, def := range funcs {
		funTokens := getFunctionTokens(def.GetName(), funcs, tokens)
		funNode := root.AddChild("function", def.GetName())
		err := parseFunction(funTokens, funNode, def.GetName())
		if err != nil {
			logger.Println(err.Error())
		}
	}

	return root
}


func parseFunction(tokens []casmutility.Token, node *Node, name string) error {
	//eturn errors.New("func parseFunction(" + name +  "): WORK IN PROGRESS")

	i := 0
	for i < len(tokens) {
		token := tokens[i]

		if token.GetType() == "operator" {
			if token.GetValue() == "=" {
				err := parseAssignment(tokens, node, i)
				if err != nil {
					return errors.New("function " + name + ": token[" + string(i) + "] left identifier expected")
				}
			}
		}


		i = i + 1
	}

	return nil
}


//NOT WORKING
func parseAssignment(tokens []casmutility.Token, node *Node, i int) error {
	left := tokens[i - 1]
	typename := tokens[i - 2]

	if left.GetType() != "identifier" {
		return errors.New("")
	}

	j := i + 1
	for {
		if tokens[j].GetType() == "string_separator" {
			break
		}
		j = j + 1
	}
	right := tokens [i + 1 : j]

	na := node.AddChild("=", "")
	na.AddChild("identifier", left.GetValue())
	na.AddChild("+", "nil")

	_ = right
	_ = typename
	return nil
}
