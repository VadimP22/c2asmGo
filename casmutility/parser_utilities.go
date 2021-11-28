package casmutility


type FunctionDefinition interface {
	GetTokenNumber() 	int
	GetName()			string
	GetArgTypes()		[]string
	GetArgNames()		[]string
	GetReturnType()		string
}


type CFunctionDefinition struct {
	tokenNumber int
	name 		string
	argsTypes	[]string
	argsNames 	[]string
	retnType	string
}


func (cfd CFunctionDefinition) GetTokenNumber() int {
	return cfd.tokenNumber
}


func (cfd CFunctionDefinition) GetName() string {
	return cfd.name
}


func (cfd CFunctionDefinition) GetArgTypes() []string {
	return cfd.argsTypes
}


func (cfd CFunctionDefinition) GetArgNames() []string {
	return cfd.argsNames
}


func (cfd CFunctionDefinition) GetReturnType() string {
	return cfd.retnType
}


func NewCFunctionDefinition(tokenNumber int, name string, argTypes, argNames []string, retType string) FunctionDefinition {
	var cfd CFunctionDefinition
	cfd.argsNames = argNames
	cfd.argsTypes = argTypes
	cfd.name = name
	cfd.tokenNumber = tokenNumber
	cfd.retnType = retType
	return cfd
}