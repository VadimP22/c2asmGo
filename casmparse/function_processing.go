package casmparse

import (
	"errors"
	"strconv"
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
		err := parseFunction(funTokens, funNode, def.GetName(), logger)
		if err != nil {
			logger.Println(err.Error())
		}
	}

	return root
}


func parseFunction(tokens []casmutility.Token, node *Node, name string, logger casmutility.Logger) error {
	//return errors.New("func parseFunction(" + name +  "): WORK IN PROGRESS")

	i := 0
	for i < len(tokens) {
		token := tokens[i]

		//assignment parsing
		if token.GetType() == "operator" {
			if token.GetValue() == "=" {
				if isSimpleAssignment(tokens, i, logger) {
					logger.Println("Parsing assignment: token[" + strconv.Itoa(i) + "]")
					err := parseAssignment(tokens, node, i)
					if err != nil {
						return errors.New("function " + name + ": token[" + string(strconv.Itoa(i)) + "] left identifier expected")
					}
				} else {
					logger.Println("assignment than causes function call parsing is WIP")
				}
			}
		}


		i = i + 1
	}

	return nil
}


//complex assignment means assignment causes function call,
//simple - not causes 
func isSimpleAssignment(tokens []casmutility.Token, i int, logger casmutility.Logger) bool {
	var j int = i + 1
	var current, next casmutility.Token
	prefixString := "    "

	current = tokens[j]

	for tokens[j].GetType() != "string_separator" {
		next = tokens[j]

		if current.GetType() == "identifier" {
			if next.GetType() == "bracket_open" {
				logger.Println("Complex assignment found:")
				logger.Println(prefixString + current.GetValue())
				return false
			}
		}

		j += 1
		current = next
	}

	return true
}


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
	//fmt.Println("parse", tokens)

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
				root.AddChild(left2[0].GetType(), left2[0].GetValue())
			}
		}
	}
}


func separateByToken(tokens []casmutility.Token,separator byte) ([]casmutility.Token, []casmutility.Token, bool) {
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