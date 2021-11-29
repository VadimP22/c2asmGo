package casmparse

import (
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
		parseFunction(funTokens, funNode)
	}

	return root
}


func parseFunction(tokens []casmutility.Token, node *Node) {
	i := 0
	for i < len(tokens) {
		token := tokens[i]

		if token.GetType() == "operator" {
			if token.GetValue() == "=" {
				parseAssignment(tokens, node, i)
			}
		}


		i = i + 1
	}
}


func parseAssignment(tokens []casmutility.Token, node *Node, i int) {
	left := tokens[i - 1]
	j := i + 1
	for {
		if tokens[j].GetType() == "string_separator" {
			break
		}
		j = j + 1
	}
	right := tokens [i + 1 : j]

	fmt.Println(left, " = ", right)

	set := node.AddChild("set", "set")
	set.AddChild("identifier", left.GetValue())
	mathNode := set.AddChild("expression", "")

	parseMathExpression(right, mathNode)

	
}


//TODO
func parseMathExpression(expression []casmutility.Token, node *Node) {
	for i := 0; i < len(expression); i++ {
		exp := expression[i]
		if exp.GetType() == "bracket_open" {
			fmt.Println("returnBrackets")
			parseMathExpression(returnBrackets(expression, node, i), node)
		}
	}
}


//TODO
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