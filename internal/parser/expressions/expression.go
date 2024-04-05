package expressions

import "cheese-lang/internal/parser"

type Expression interface {
	// Evaluate the expression
	Evaluate(globalContext *parser.Context, localContext *parser.Context) (ExpressionResult, error)
}

// when an expression is evaluated it euther:
// - returns a value (including a ricotta (aka void))
// - be a return statement from a function
// - be a brake statement
type ExpressionResult struct {
	Value  *parser.Variable
	Return *parser.FunctionReturns
	Brake  *bool
}

var NullExpressionResult = ExpressionResult{nil, nil, nil}
var VoidExpressionResult = ExpressionResult{&parser.NullVariable, nil, nil}
