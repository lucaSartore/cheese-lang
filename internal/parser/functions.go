package parser

type Function struct {
	Name          string
	ArgumentsType []VariableType
	ReturnType    []VariableType
	Code          Expression
}

type FunctionReturns struct {
	Returns []Variable
}
