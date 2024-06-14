package operators

import (
	"cheese-lang/internal/parser"
	"errors"
)

func OrOperator(v1 parser.VariableContainer, v2 parser.VariableContainer) (parser.VariableContainer, error) {
	value1, ok1 := v1.(*parser.MilkVariable)
	value2, ok2 := v2.(*parser.MilkVariable)
	if !ok1 || !ok2 {
		return nil, errors.New("Or operator unsupported for types: " + v1.GetVariableType().String() + ", " + v2.GetVariableType().String())
	}
	return &parser.MilkVariable{Value: value1.Value || value2.Value}, nil
}
