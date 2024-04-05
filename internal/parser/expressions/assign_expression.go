package expressions

import (
	"cheese-lang/internal/parser"
	"errors"
)

// AssignExpression is used every time a variable is assigned a value

type AssignExpression struct {
	variableToAssign string
	valueToAssign    Expression
}

func (a AssignExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (ExpressionResult, error) {

	toAssign, ok := parser.GetVariable(localContext, globalContext, a.variableToAssign)

	if !ok {
		return NullExpressionResult, errors.New("Variable with name: " + a.variableToAssign + " not found in context")
	}

	result, err := a.valueToAssign.Evaluate(globalContext, localContext)

	if err != nil {
		return NullExpressionResult, err
	}

	if result.Return == nil {
		return NullExpressionResult, errors.New("no valid assignment value at the right side of the assignment")
	}

	if toAssign.Value.GetVariableType() != result.Value.Value.GetVariableType() {
		return NullExpressionResult, errors.New("Type mismatch in assignment, expected: " + toAssign.Value.GetVariableType().String() + " got: " + result.Value.Value.GetVariableType().String())
	}

	toAssign.Value = result.Value.Value

	return VoidExpressionResult, nil
}
