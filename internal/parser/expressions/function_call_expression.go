package expressions

import (
	"cheese-lang/internal/parser"
	"errors"
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
		return parser.NullExpressionResult, errors.New(fmt.Sprintf("unable to find the function: %s", fc.functionToCall))
	}

}
