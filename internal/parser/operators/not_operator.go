package operators

import (
	"cheese-lang/internal/parser"
	"errors"
)

func NotOperator(v parser.VariableContainer) (parser.VariableContainer, error) {
	value, ok := v.(*parser.MilkVariable)
	if !ok {
		return nil, errors.New("Not operator unsupported for type: " + v.GetVariableType().String())
	}
	return &parser.MilkVariable{Value: !value.Value}, nil
}
