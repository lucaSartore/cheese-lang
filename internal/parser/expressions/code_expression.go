package expressions

import (
	"cheese-lang/internal/parser"
)

// a code expression is a list of expressions that are evaluated in order

type CodeExpression struct {
	Expressions []Expression
}

func (ce *CodeExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (ExpressionResult, error) {
	lastResult := VoidExpressionResult

	for _, exp := range ce.Expressions {

		lastResult, err := exp.Evaluate(globalContext, localContext)

		if err != nil {
			return NullExpressionResult, nil
		}

		if lastResult.Return != nil || lastResult.Brake != nil {
			return lastResult, nil
		}
	}
	return lastResult, nil
}
