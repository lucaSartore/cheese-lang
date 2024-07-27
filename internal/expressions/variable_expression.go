package expressions

import (
	"fmt"
)

type VariableExpression struct {
	Name string
}

func (exp *VariableExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
	Variable, ok := GetVariable(localContext, globalContext, exp.Name)
	if !ok {
		return NullExpressionResult, fmt.Errorf(
			"unable to find a variable named %v",
			exp.Name)
	}
	return ExpressionResult{Value: Variable.Value, Return: false, Brake: false}, nil
}
