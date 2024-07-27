package operators

import (
	"cheese-lang/internal/expressions"
	"errors"
)

func EqualOperator(v1 expressions.VariableContainer, v2 expressions.VariableContainer) (expressions.VariableContainer, error) {
	value1, ok1 := v1.(*expressions.ParmesanVariable)
	value2, ok2 := v2.(*expressions.ParmesanVariable)
	if ok1 && ok2 {
		return &expressions.MilkVariable{Value: value1.Value == value2.Value}, nil
	}

	value3, ok1 := v1.(*expressions.GorgonzolaVariable)
	value4, ok2 := v2.(*expressions.GorgonzolaVariable)
	if ok1 && ok2 {
		return &expressions.MilkVariable{Value: value3.Value == value4.Value}, nil
	}

	value5, ok1 := v1.(*expressions.MilkVariable)
	value6, ok2 := v2.(*expressions.MilkVariable)
	if ok1 && ok2 {
		return &expressions.MilkVariable{Value: value5.Value == value6.Value}, nil
	}

	value7, ok1 := v1.(*expressions.MozzarellaVariable)
	value8, ok2 := v1.(*expressions.MozzarellaVariable)
	if ok1 && ok2 {
		return &expressions.MilkVariable{Value: value7.Value == value8.Value}, nil
	}

	return nil, errors.New("Unequal operator unsupported for types: " + v1.GetVariableType().String() + ", " + v2.GetVariableType().String())
}
