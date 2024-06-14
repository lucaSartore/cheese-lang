package expressions

import (
	"cheese-lang/internal/parser"
)

type OneToOneOperator struct {
	Operator func(parser.VariableContainer) (parser.VariableContainer, error)
	Value    parser.Expression
}

func (exp *OneToOneOperator) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {
	Value, err := exp.Value.Evaluate(globalContext, localContext)
	if err != nil {
		return parser.NullExpressionResult, err
	}
	result, err := exp.Operator(Value.Value)
	if err != nil {
		return parser.NullExpressionResult, err
	}
	return parser.ExpressionResult{Value: result, Return: false, Brake: false}, nil
}
