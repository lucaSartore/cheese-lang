package operators

import (
	"cheese-lang/internal/expressions"
	"errors"
)

func OrOperator(v1 expressions.VariableContainer, v2 expressions.VariableContainer) (expressions.VariableContainer, error) {
	value1, ok1 := v1.(*expressions.MilkVariable)
	value2, ok2 := v2.(*expressions.MilkVariable)
	if !ok1 || !ok2 {
		return nil, errors.New("Or operator unsupported for types: " + v1.GetVariableType().String() + ", " + v2.GetVariableType().String())
	}
	return &expressions.MilkVariable{Value: value1.Value || value2.Value}, nil
}
