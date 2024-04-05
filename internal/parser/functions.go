package parser

type Function struct {
	Name          string
	ArgumentsType []VariableType
	ReturnType    []VariableType
}

type FunctionArg struct {
	Name  string
	Value Variable
}

type FunctionReturns struct {
	Returns []Variable
}
