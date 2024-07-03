package expressions

import (
	"cheese-lang/internal/parser"
)

type TwoToOneOperatorExpression struct {
	LeftValue  parser.Expression
	RightValue parser.Expression
	Operator   func(parser.VariableContainer, parser.VariableContainer) (parser.VariableContainer, error)
}

func (exp *TwoToOneOperatorExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {
	leftValue, err := exp.LeftValue.Evaluate(globalContext, localContext)
	if err != nil {
		return parser.NullExpressionResult, err
	}
	rightValue, err := exp.RightValue.Evaluate(globalContext, localContext)
	if err != nil {
		return parser.NullExpressionResult, err
	}
	result, err := exp.Operator(rightValue.Value, leftValue.Value)
	if err != nil {
		return parser.NullExpressionResult, err
	}
	return parser.ExpressionResult{Value: result, Return: false, Brake: false}, nil
}
