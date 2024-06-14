package expressions

import (
	"cheese-lang/internal/parser"
	"fmt"
)

// FunctionCallExpression is used every time a function is called

type FunctionCallExpression struct {
	functionToCall string
	args           []string
}

func (fc *FunctionCallExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {

	function, ok := parser.GetFunction(localContext, globalContext, fc.functionToCall)

	if !ok {
		return parser.NullExpressionResult, fmt.Errorf("unable to find the function: %s", fc.functionToCall)
	}

	// create the new context for the function call
	var newLocalContext = parser.MakeContext()

	if len(fc.args) != int(len(function.ArgumentsType)) {
		return parser.NullExpressionResult, fmt.Errorf(
			"the function %s take in input %d parameters, but %d were given instead",
			fc.functionToCall,
			len(function.ArgumentsType),
			len(fc.args))
	}

	for i := range len(fc.args) {
		local_variable_name := fc.args[i]
		local_variable, ok := parser.GetVariable(localContext, globalContext, local_variable_name)
		if !ok {
			return parser.NullExpressionResult, fmt.Errorf("unable to find the variable: %s", local_variable_name)
		}

		local_variable_type := local_variable.Value.GetVariableType()
		expected_variable_type := function.ArgumentsType[i]

		arg_name := function.ArgumentsNames[i]

		if local_variable_type != expected_variable_type {
			return parser.NullExpressionResult, fmt.Errorf(
				"expected type %s for argument %s got type %s instead",
				expected_variable_type.String(),
				local_variable_type.String(),
				arg_name)
		}

		newLocalContext.Variables[arg_name] = parser.MakeVariable(local_variable_name, local_variable.Value)
	}

	returnValue, error := function.Code.Evaluate(&newLocalContext, globalContext)
	if error != nil {
		return parser.NullExpressionResult, error
	}
	// the function failed to return a value, therefore there was an error in the code parsing phase
	if returnValue.Brake {
		panic("parser error 1 in function call expression")
	}
	if returnValue.Return {
		panic("parser error 2 in function call expression")
	}
	if returnValue.Value == nil {
		panic("parser error 3 in function call expression")
	}
	// return the entire variable, an assignment expression will put the return items
	// in the respective variables
	return returnValue, nil
}
