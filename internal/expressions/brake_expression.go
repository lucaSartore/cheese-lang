package expressions

type BrakeExpression struct {
	CodeInside Expression
}

func (ce *BrakeExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
	v := VoidExpressionResult
	v.Brake = true
	return v, nil
}
