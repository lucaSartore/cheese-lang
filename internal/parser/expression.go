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

// null expression result is used as a null value when you need to return an error
var NullExpressionResult = ExpressionResult{nil, nil, nil}

// void expression result is used when expressions don't need to return anything
var VoidExpressionResult = ExpressionResult{&NullVariableContainer, nil, nil}
