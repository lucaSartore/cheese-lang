package expressions

import (
	"cheese-lang/internal/parser"
)

// a code expression is a list of expressions that are evaluated in order

type CodeExpression struct {
	Expressions []parser.Expression
}

func (ce *CodeExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {
	lastResult := parser.VoidExpressionResult

	for _, exp := range ce.Expressions {

		lastResult, err := exp.Evaluate(globalContext, localContext)

		if err != nil {
			return parser.NullExpressionResult, nil
		}

		if lastResult.Return || lastResult.Brake {
			return lastResult, nil
		}
	}
	return lastResult, nil
}
