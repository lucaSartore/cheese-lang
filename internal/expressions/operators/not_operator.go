package operators

import (
	"cheese-lang/internal/expressions"
	"errors"
)

func NotOperator(v expressions.VariableContainer) (expressions.VariableContainer, error) {
	value, ok := v.(*expressions.MilkVariable)
	if !ok {
		return nil, errors.New("Not operator unsupported for type: " + v.GetVariableType().String())
	}
	return &expressions.MilkVariable{Value: !value.Value}, nil
}
