package operators

import (
	"cheese-lang/internal/parser"
	"errors"
)

func NotOperator(v parser.Variable) (parser.Variable, error) {
	value, ok := v.Value.(*parser.MilkVariable)
	if !ok {
		return parser.NullVariable, errors.New("Not operator unsupported for type: " + v.Value.GetVariableType().String())
	}
	return parser.Variable{Value: &parser.MilkVariable{Value: !value.Value}}, nil
}
