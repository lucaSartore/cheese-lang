package expressions

// CuddleExpression is the basic loop control flow of the language

type CuddleExpression struct {
	CodeInside Expression
}

func (ce *CuddleExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {

	for {
		result, err := ce.CodeInside.Evaluate(globalContext, localContext)

		if err != nil {
			return NullExpressionResult, err
		}

		if result.Return {
			return result, nil
		}

		if result.Brake {
			// to avoid thai i endlessly brake around loops
			result.Brake = false
			return result, nil
		}
	}

}
