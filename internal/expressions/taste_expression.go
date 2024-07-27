package expressions

import (
	"fmt"
)

type TasteExpression struct {
	Condition Expression
	Code      Expression
}

func (exp *TasteExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
	result, err := exp.Condition.Evaluate(globalContext, localContext)
	if err != nil {
		return NullExpressionResult, err
	}
	milk, ok := result.Value.(*MilkVariable)
	if !ok {
		return NullExpressionResult, fmt.Errorf(
			"expected a Milk value inside Taste condition but got %s",
			result.Value.GetVariableType().String())
	}
	if milk.Value {
		return exp.Code.Evaluate(globalContext, localContext)
	}
	return VoidExpressionResult, nil
}
