package casmutility


type FunctionDefinition interface {
	GetTokenNumber() 	int
	GetName()			string
	GetArgsTypes()		[]string
	GetArgsNames()		[]string
	GetReturns()		string
}