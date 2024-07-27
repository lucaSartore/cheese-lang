package expressions

type TwoToOneOperatorExpression struct {
	LeftValue  Expression
	RightValue Expression
	Operator   func(VariableContainer, VariableContainer) (VariableContainer, error)
}

func (exp *TwoToOneOperatorExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
	leftValue, err := exp.LeftValue.Evaluate(globalContext, localContext)
	if err != nil {
		return NullExpressionResult, err
	}
	rightValue, err := exp.RightValue.Evaluate(globalContext, localContext)
	if err != nil {
		return NullExpressionResult, err
	}
	result, err := exp.Operator(leftValue.Value, rightValue.Value)
	if err != nil {
		return NullExpressionResult, err
	}
	return ExpressionResult{Value: result, Return: false, Brake: false}, nil
}
