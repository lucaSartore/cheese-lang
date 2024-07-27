package expressions

import (
	"errors"
	"fmt"
)

// AssignExpression is used every time a variable is assigned a value

type AssignExpression struct {
	VariablesToAssign []string
	ValueToAssign     Expression
}

func (a AssignExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {

	if len(a.VariablesToAssign) == 0 {
		panic("Error in AssignExpression construction, the parser did not enforce the presence of at least one variable at the right of the assign operation")
	}

	result, err := a.ValueToAssign.Evaluate(globalContext, localContext)

	if err != nil {
		return NullExpressionResult, err
	}

	if result.Value == nil {
		return NullExpressionResult, errors.New("no valid assignment value at the right side of the assignment")
	}

	if result.Value.GetVariableType() == Ricotta {
		return NullExpressionResult, fmt.Errorf("trying to assign a value Ricotta (Void) value")
	}

	counterLeft := len(a.VariablesToAssign)
	counterRight := 1

	tupleValue, isTuple := result.Value.(*TupleVariableType)

	if isTuple {
		counterRight = len((*tupleValue).Variables)
	}

	if counterLeft != counterRight {
		return NullExpressionResult, fmt.Errorf("impossible to unpack %v values into %v variables", counterRight, counterLeft)
	}

	if !isTuple {
		assignSingeVariable(a.VariablesToAssign[0], result.Value, globalContext, localContext)
		return VoidExpressionResult, nil
	}

	for i, variableName := range a.VariablesToAssign {
		err := assignSingeVariable(variableName, tupleValue.Variables[i], globalContext, localContext)

		if err != nil {
			return NullExpressionResult, err
		}
	}

	return VoidExpressionResult, nil
}

func assignSingeVariable(variableName string, valueToAssign VariableContainer, globalContext *Context, localContext *Context) error {

	variable, ok := GetVariable(localContext, globalContext, variableName)

	if !ok {
		return errors.New("Variable with name: " + variableName + " not found in current context")
	}

	if variable.Value.GetVariableType() == Tuple {
		panic("The parser should not allow to create a tuple variable")
	}

	if variable.Value.GetVariableType() != valueToAssign.GetVariableType() {
		return fmt.Errorf("unsupported assignment between values with type: %v and %v", variable.Value.GetVariableType(), valueToAssign.GetVariableType())
	}

	SetVariable(localContext, globalContext, variableName, valueToAssign)
	return nil
}
