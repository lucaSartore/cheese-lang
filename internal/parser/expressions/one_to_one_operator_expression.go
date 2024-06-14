package expressions

import (
	"cheese-lang/internal/parser"
	"fmt"
)

type OneToOneOperator struct {
	Operator     func(parser.VariableContainer) (parser.VariableContainer, error)
	VariableName string
}

func (exp *OneToOneOperator) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {
	variable, ok := parser.GetVariable(localContext, globalContext, exp.VariableName)
	if !ok {
		return parser.NullExpressionResult, fmt.Errorf("unable to find the variable: %s", exp.VariableName)
	}
	result, err := exp.Operator(variable.Value)
	if err != nil {
		return parser.NullExpressionResult, err
	}
	return parser.ExpressionResult{Value: result, Return: false, Brake: false}, nil
}
