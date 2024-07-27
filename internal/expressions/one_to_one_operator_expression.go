package expressions

type OneToOneOperator struct {
	Operator func(VariableContainer) (VariableContainer, error)
	Value    Expression
}

func (exp *OneToOneOperator) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
	Value, err := exp.Value.Evaluate(globalContext, localContext)
	if err != nil {
		return NullExpressionResult, err
	}
	result, err := exp.Operator(Value.Value)
	if err != nil {
		return NullExpressionResult, err
	}
	return ExpressionResult{Value: result, Return: false, Brake: false}, nil
}
