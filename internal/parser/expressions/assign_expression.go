package expressions

import (
	"cheese-lang/internal/parser"
	"errors"
)

// AssignExpression is used every time a variable is assigned a value

// todo: change variable to assign to []string to support tuple

type AssignExpression struct {
	variableToAssign string
	valueToAssign    parser.Expression
}

func (a AssignExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {

	toAssign, ok := parser.GetVariable(localContext, globalContext, a.variableToAssign)

	if !ok {
		return parser.NullExpressionResult, errors.New("Variable with name: " + a.variableToAssign + " not found in context")
	}

	result, err := a.valueToAssign.Evaluate(globalContext, localContext)

	if err != nil {
		return parser.NullExpressionResult, err
	}

	if result.Return == nil {
		return parser.NullExpressionResult, errors.New("no valid assignment value at the right side of the assignment")
	}

	if toAssign.Value.GetVariableType() != (*result.Value).GetVariableType() {
		return parser.NullExpressionResult, errors.New("Type mismatch in assignment, expected: " + toAssign.Value.GetVariableType().String() + " got: " + result.Value.Value.GetVariableType().String())
	}

	toAssign.Value = *result.Value

	return parser.VoidExpressionResult, nil
}
