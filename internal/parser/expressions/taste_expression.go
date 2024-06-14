package expressions

import (
	"cheese-lang/internal/parser"
	"fmt"
)

type TasteExpression struct {
	Condition parser.Expression
	Code      parser.Expression
}

func (exp *TasteExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {
	result, err := exp.Condition.Evaluate(globalContext, localContext)
	if err != nil {
		return parser.NullExpressionResult, err
	}
	milk, ok := result.Value.(*parser.MilkVariable)
	if !ok {
		return parser.NullExpressionResult, fmt.Errorf(
			"expected a Milk value inside Taste condition but got %s",
			result.Value.GetVariableType().String())
	}
	if milk.Value {
		return exp.Code.Evaluate(globalContext, localContext)
	}
	return parser.VoidExpressionResult, nil
}
