package parser

type Function struct {
	Name           string
	ArgumentsType  []VariableType
	ArgumentsNames []string
	ReturnType     []VariableType
	Code           Expression
}

type FunctionReturns struct {
	Returns []Variable
}
