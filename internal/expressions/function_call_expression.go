package expressions

import (
	"fmt"
)

// FunctionCallExpression is used every time a function is called

type FunctionCallExpression struct {
	FunctionToCall string
	Args           []Expression
}

func (fc *FunctionCallExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {

	function, ok := GetFunction(localContext, globalContext, fc.FunctionToCall)

	if !ok {
		return NullExpressionResult, fmt.Errorf("unable to find the function: %s", fc.FunctionToCall)
	}

	// create the new context for the function call
	var newLocalContext = MakeContext()

	if len(fc.Args) != int(len(function.ArgumentsType)) {
		return NullExpressionResult, fmt.Errorf(
			"the function %s take in input %d parameters, but %d were given instead",
			fc.FunctionToCall,
			len(function.ArgumentsType),
			len(fc.Args))
	}

	for i := range len(fc.Args) {
		value, err := fc.Args[i].Evaluate(globalContext, localContext)
		if err != nil {
			return NullExpressionResult, err
		}

		value_type := value.Value.GetVariableType()
		expected_type := function.ArgumentsType[i]

		arg_name := function.ArgumentsNames[i]

		if value_type != expected_type {
			return NullExpressionResult, fmt.Errorf(
				"expected type %s for argument %d got type %s instead",
				expected_type.String(),
				i,
				value_type.String(),
			)
		}

		newLocalContext.Variables[arg_name] = MakeVariable(arg_name, value.Value)
	}

	returnValue, error := function.Code.Evaluate(globalContext, &newLocalContext)

	if error != nil {
		return NullExpressionResult, error
	}

	if returnValue.Brake {
		panic("parser error 1 in function call expression")
	}

	// the function failed to return a value, therefore there was an error in the code parsing phase
	if !returnValue.Return {
		return NullExpressionResult, fmt.Errorf("function %s did not return any value", fc.FunctionToCall)
	}

	returnValue.Return = false
	if returnValue.Value == nil {
		panic("parser error 3 in function call expression")
	}
	// return the entire variable, an assignment expression will put the return items
	// in the respective variables
	return returnValue, nil
}
