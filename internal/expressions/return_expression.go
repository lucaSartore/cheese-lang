package expressions

type ReturnExpression struct {
	Expression Expression
}

func (exp *ReturnExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
	result, err := exp.Expression.Evaluate(globalContext, localContext)
	if err != nil {
		return result, err
	}
	result.Return = true
	return result, nil
}
