package expressions

type Function struct {
	Name           string
	ArgumentsType  []VariableType
	ArgumentsNames []string
	Code           Expression
}

type FunctionReturns struct {
	Returns []Variable
}

func MakeFunction(name string, code Expression, argumentsType []VariableType, argumentsNames []string) Function {
	if len(argumentsType) != len(argumentsNames) {
		panic("Arguments type and names must have the same length")
	}
	return Function{
		Name:           name,
		Code:           code,
		ArgumentsType:  argumentsType,
		ArgumentsNames: argumentsNames,
	}
}
