package expressions

import (
	"cheese-lang/internal/parser"
	"fmt"
)

type TwoToOneOperatorExpression struct {
	LeftVariable  string
	RightVariable string
	Operator      func(parser.VariableContainer, parser.VariableContainer) (parser.VariableContainer, error)
}

func (exp *TwoToOneOperatorExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {
	rightVariable, ok := parser.GetVariable(localContext, globalContext, exp.RightVariable)
	if !ok {
		return parser.NullExpressionResult, fmt.Errorf(
			"unable to find a variable named %v",
			exp.RightVariable)
	}
	leftVariable, ok := parser.GetVariable(localContext, globalContext, exp.LeftVariable)
	if !ok {
		return parser.NullExpressionResult, fmt.Errorf(
			"unable to find a variable named %v",
			exp.LeftVariable)
	}
	result, err := exp.Operator(rightVariable.Value, leftVariable.Value)
	if err != nil {
		return parser.NullExpressionResult, err
	}
	return parser.ExpressionResult{Value: result, Return: false, Brake: false}, nil
}
