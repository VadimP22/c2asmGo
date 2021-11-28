package casmutility

import (
	"errors"
	"io/ioutil"
	"strconv"
)


func GetSourceCode(filepath string, logger Logger) (string, error) {
	logger.Println("file: '" + filepath + "'")
	sourceCode, err := ioutil.ReadFile(filepath)
	
	if err != nil {
		return "", errors.New("at func casmutility.GetSourceCode: " + err.Error())
	}

	size := len(sourceCode)
	sizeStr := strconv.Itoa(size)

	logger.Println("size: " + sizeStr + " bytes")

	return string(sourceCode), nil
}


func ParseArgs(args []string, logger Logger) (string, []string) {
	logger.Println("parsing args...")
	if len(args) < 2 {
		logger.Println("fileName: input.txt")
		logger.Println("No args")
		return "input.txt", nil
	} else {
		var pureArgs[]string

		filename := "input.txt"
		
		for i, arg := range args {
			if i == 0 {
				logger.Println("path: " + arg)
			} else if i == 1 {
				filename = arg
				logger.Println("fileName: " + filename)
			} else if arg[0] == '-' {
				pureArgs = append(args, arg)
				logger.Println("argName: " + arg)
			} else {
				pureArgs = append(args, arg)
				logger.Println("argValue: " + arg)
			}
		}

		return filename, pureArgs
	}
}


func CheckEntryFunction(name string, defs []FunctionDefinition) error {
	for _, function := range defs {
		if function.GetName() == name {
			if function.GetReturnType() == "int" {
				return nil
			}
		}
	}

	return errors.New("there's no entry function [int " + name + "()] in your file")
}


//TODO
func getArgValue(arg string) (string, error) {
	return "", nil
}