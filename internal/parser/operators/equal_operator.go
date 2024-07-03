package operators

import (
	"cheese-lang/internal/parser"
	"errors"
)

func EqualOperator(v1 parser.VariableContainer, v2 parser.VariableContainer) (parser.VariableContainer, error) {
	value1, ok1 := v1.(*parser.ParmesanVariable)
	value2, ok2 := v2.(*parser.ParmesanVariable)
	if ok1 && ok2 {
		return &parser.MilkVariable{Value: value1.Value == value2.Value}, nil
	}

	value3, ok1 := v1.(*parser.GorgonzolaVariable)
	value4, ok2 := v2.(*parser.GorgonzolaVariable)
	if ok1 && ok2 {
		return &parser.MilkVariable{Value: value3.Value == value4.Value}, nil
	}

	value5, ok1 := v1.(*parser.MilkVariable)
	value6, ok2 := v2.(*parser.MilkVariable)
	if ok1 && ok2 {
		return &parser.MilkVariable{Value: value5.Value == value6.Value}, nil
	}

	value7, ok1 := v1.(*parser.MozzarellaVariable)
	value8, ok2 := v1.(*parser.MozzarellaVariable)
	if ok1 && ok2 {
		return &parser.MilkVariable{Value: value7.Value == value8.Value}, nil
	}

	return nil, errors.New("Unequal operator unsupported for types: " + v1.GetVariableType().String() + ", " + v2.GetVariableType().String())
}
