package operators

import (
	"cheese-lang/internal/parser"
	"errors"
)

func OrOperator(v1, v2 parser.Variable) (parser.Variable, error) {
	value1, ok1 := v1.Value.(*parser.MilkVariable)
	value2, ok2 := v2.Value.(*parser.MilkVariable)
	if !ok1 || !ok2 {
		return parser.NullVariable, errors.New("Or operator unsupported for types: " + v1.Value.GetVariableType().String() + ", " + v2.Value.GetVariableType().String())
	}
	return parser.Variable{Value: &parser.MilkVariable{Value: value1.Value || value2.Value}}, nil
}
