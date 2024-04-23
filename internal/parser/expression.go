package parser

type Expression interface {
	// Evaluate the expression
	Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error)
}

// when an expression is evaluated it euther:
// - returns a value (including a ricotta (aka void))
// - be a return statement from a function
// - be a brake statement
type ExpressionResult struct {
	Value  *VariableContainer
	Return *FunctionReturns
	Brake  *bool
}

var NullExpressionResult = ExpressionResult{nil, nil, nil}
var VoidExpressionResult = ExpressionResult{&NullVariable, nil, nil}
