package expressions

import (
	"cheese-lang/internal/parser"
	"fmt"
)

type VariableExpression struct {
	Name string
}

func (exp *VariableExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {
	Variable, ok := parser.GetVariable(localContext, globalContext, exp.Name)
	if !ok {
		return parser.NullExpressionResult, fmt.Errorf(
			"unable to find a variable named %v",
			exp.Name)
	}
	return parser.ExpressionResult{Value: Variable.Value, Return: false, Brake: false}, nil
}
