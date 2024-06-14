package expressions

import (
	"cheese-lang/internal/parser"
)

type LiteralExpression struct {
	Literal parser.VariableContainer
}

func (exp *LiteralExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {
	return parser.ExpressionResult{Value: exp.Literal, Return: false, Brake: false}, nil
}
