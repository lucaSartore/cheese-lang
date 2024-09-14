package operators

import (
	"cheese-lang/internal/expressions"
	"errors"
)

func NegateOperator(v expressions.VariableContainer) (expressions.VariableContainer, error) {
	valueP, ok := v.(*expressions.ParmesanVariable)
	if ok {
	    return &expressions.ParmesanVariable{Value: -valueP.Value}, nil
	}
	valueG, ok := v.(*expressions.GorgonzolaVariable)
	if ok {
	    return &expressions.GorgonzolaVariable{Value: -valueG.Value}, nil
	}
	return nil, errors.New("Negate operator unsupported for type: " + v.GetVariableType().String())
}
