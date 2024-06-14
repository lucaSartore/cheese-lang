package expressions

import (
	"cheese-lang/internal/parser"
)

type ReturnExpression struct {
	Expression parser.Expression
}

func (exp *ReturnExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {
	result, err := exp.Expression.Evaluate(globalContext, localContext)
	if err != nil {
		return result, err
	}
	result.Return = true
	return result, nil
}
