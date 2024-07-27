package expressions

type LiteralExpression struct {
	Literal VariableContainer
}

func (exp *LiteralExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
	return ExpressionResult{Value: exp.Literal, Return: false, Brake: false}, nil
}
