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


//working, but WIP
//TODO
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
	expNode := na.AddChild("expression", "")
	parseMathExpression(right, expNode)

	_ = right
	_ = typename
	return nil
}


func parseMathExpression(tokens []casmutility.Token, root *Node) {
	if tokens[0].GetType() == "bracket_open" && tokens[len(tokens) - 1].GetType() == "bracket_close" {
		tokens = tokens[1:len(tokens) - 1]
	}
	fmt.Println("parse", tokens)

	left, right, isTwo := separateByToken(tokens, '+')
	if isTwo {
		newNode := root.AddChild("+", "")
		parseMathExpression(left, newNode)
		parseMathExpression(right, newNode)
	} else {
		left1, right1, isTwo1 := separateByToken(tokens, '-')
		if isTwo1 {
			newNode1 := root.AddChild("-", "")
			parseMathExpression(left1, newNode1)
			parseMathExpression(right1, newNode1)
		} else {
			left2, right2, isTwo2 := separateByToken(tokens, '*')
			if isTwo2 {
				newNode2 := root.AddChild("*", "")
				parseMathExpression(left2, newNode2)
				parseMathExpression(right2, newNode2)
			} else {
				root.AddChild(left2[0].GetValue(), "")
			}
		}
	}
}


func  separateByToken(tokens []casmutility.Token,separator byte) ([]casmutility.Token, []casmutility.Token, bool) {
	sepChar := string(separator)
	open, close := 0, 0
	for i := 0; i < len(tokens); i++ {
		curr := tokens[i]

		if curr.GetType() == "bracket_open" {
			open += 1
		}

		if curr.GetType() == "bracket_close" {
			close += 1
		}

		if curr.GetType() == "operator" && curr.GetValue() == sepChar {
			if open == close {
				return tokens[:i], tokens[i + 1:], true
			}
		}
	}
	return tokens, nil, false
}