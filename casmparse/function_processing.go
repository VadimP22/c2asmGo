package casmparse

import (
	"errors"
	"fmt"

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
	return errors.New("func parseFunction(" + name +  "): WORK IN PROGRESS")

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
	//typename := tokens[i - 2]
	left := tokens[i - 1]

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

	fmt.Println(left, " = ", right)

	set := node.AddChild("=", "")
	set.AddChild("left", left.GetValue())
	mathNode := set.AddChild("right", "")

	parseMathExpression(right, mathNode)

	return nil
}


//NOT WORKING
func parseMathExpression(expression []casmutility.Token, node *Node) {
	for i := 0; i < len(expression); i++ {
		exp := expression[i]
		if exp.GetType() == "bracket_open" {
			fmt.Println("returnBrackets")
			parseMathExpression(returnBrackets(expression, node, i), node)
		}
	}
}


//NOT WORKING
func returnBrackets(expression []casmutility.Token, node *Node, i int) []casmutility.Token {
	j := i
	open, close := 0, 0
	var retTokens []casmutility.Token

	for {
		exp := expression[j]
		if exp.GetType() == "bracket_open" {
			open = open + 1
		}
		if exp.GetType() == "bracket_close" {
			close = close + 1
		}
		if open == close {
			break
		}
		j = j + 1
	}
	retTokens = expression[i + 1 : j]
	fmt.Println(retTokens)

	return retTokens
}