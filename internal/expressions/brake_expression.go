package expressions

type BrakeExpression struct {
}

func (ce *BrakeExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
	v := VoidExpressionResult
	v.Brake = true
	return v, nil
}
