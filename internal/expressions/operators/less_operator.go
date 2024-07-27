package operators

import (
	"cheese-lang/internal/expressions"
	"errors"
)

func LessOperator(v1 expressions.VariableContainer, v2 expressions.VariableContainer) (expressions.VariableContainer, error) {
	value1, ok1 := v1.(*expressions.ParmesanVariable)
	value2, ok2 := v2.(*expressions.ParmesanVariable)
	if ok1 && ok2 {
		return &expressions.MilkVariable{Value: value1.Value < value2.Value}, nil
	}

	value3, ok1 := v1.(*expressions.GorgonzolaVariable)
	value4, ok2 := v2.(*expressions.GorgonzolaVariable)
	if ok1 && ok2 {
		return &expressions.MilkVariable{Value: value3.Value < value4.Value}, nil
	}

	return nil, errors.New("Less operator unsupported for types: " + v1.GetVariableType().String() + ", " + v2.GetVariableType().String())
}
