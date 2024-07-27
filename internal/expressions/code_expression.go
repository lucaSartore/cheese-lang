package expressions

// a code expression is a list of expressions that are evaluated in order

type CodeExpression struct {
	Expressions []Expression
}

func (ce *CodeExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
	lastResult := VoidExpressionResult

	for _, exp := range ce.Expressions {

		lastResult, err := exp.Evaluate(globalContext, localContext)

		if err != nil {
			return NullExpressionResult, err
		}

		if lastResult.Return {
			return lastResult, nil
		}

		if lastResult.Brake {
			return lastResult, nil
		}
	}
	return lastResult, nil
}
