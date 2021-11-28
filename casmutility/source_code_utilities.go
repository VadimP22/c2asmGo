package casmutility


import (
	"errors"
	"io/ioutil"
)


func GetSourceCode(filepath string) (string, error) {
	sourceCode, err := ioutil.ReadFile(filepath)

	if err != nil {
		return "", errors.New("at func casmutility.GetSourceCode: " + err.Error())
	}

	return string(sourceCode), nil
}