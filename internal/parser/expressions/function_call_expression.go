package expressions

import (
	"cheese-lang/internal/parser"
	"fmt"
)

// FunctionCallExpression is used every time a function is called

type FunctionCallExpression struct {
	FunctionToCall string
	Args           []parser.Expression
}

func (fc *FunctionCallExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {

	function, ok := parser.GetFunction(localContext, globalContext, fc.FunctionToCall)

	if !ok {
		return parser.NullExpressionResult, fmt.Errorf("unable to find the function: %s", fc.FunctionToCall)
	}

	// create the new context for the function call
	var newLocalContext = parser.MakeContext()

	if len(fc.Args) != int(len(function.ArgumentsType)) {
		return parser.NullExpressionResult, fmt.Errorf(
			"the function %s take in input %d parameters, but %d were given instead",
			fc.FunctionToCall,
			len(function.ArgumentsType),
			len(fc.Args))
	}

	for i := range len(fc.Args) {
		value, err := fc.Args[i].Evaluate(globalContext, localContext)
		if err != nil {
			return parser.NullExpressionResult, err
		}

		value_type := value.Value.GetVariableType()
		expected_type := function.ArgumentsType[i]

		arg_name := function.ArgumentsNames[i]

		if value_type != expected_type {
			return parser.NullExpressionResult, fmt.Errorf(
				"expected type %s for argument %d got type %s instead",
				expected_type.String(),
				i,
				value_type.String(),
			)
		}

		newLocalContext.Variables[arg_name] = parser.MakeVariable(arg_name, value.Value)
	}

	returnValue, error := function.Code.Evaluate(globalContext, &newLocalContext)
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
