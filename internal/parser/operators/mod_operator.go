package operators

import (
	"cheese-lang/internal/parser"
	"errors"
)

func ModOperator(v1 parser.VariableContainer, v2 parser.VariableContainer) (parser.VariableContainer, error) {
	value1, ok1 := v1.(*parser.ParmesanVariable)
	value2, ok2 := v2.(*parser.ParmesanVariable)
	if ok1 && ok2 {
		return &parser.ParmesanVariable{Value: value1.Value % value2.Value}, nil
	}

	return nil, errors.New("Modulo unsupported for types: " + v1.GetVariableType().String() + ", " + v2.GetVariableType().String())
}
