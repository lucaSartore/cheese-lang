package expressions

import (
	"cheese-lang/internal/parser"
	"errors"
	"fmt"
)

// AssignExpression is used every time a variable is assigned a value

type AssignExpression struct {
	variablesToAssign []string
	valueToAssign     parser.Expression
}

func (a AssignExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {

	if len(a.variablesToAssign) == 0 {
		panic("Error in AssignExpression construction, the parser did not enforce the presence of at least one variable at the right of the assign operation")
	}

	result, err := a.valueToAssign.Evaluate(globalContext, localContext)

	if err != nil {
		return parser.NullExpressionResult, err
	}

	if result.Value == nil {
		return parser.NullExpressionResult, errors.New("no valid assignment value at the right side of the assignment")
	}

	if (*result.Value).GetVariableType() == parser.Ricotta {
		return parser.NullExpressionResult, errors.New(fmt.Sprintf("Trying to assign a value Ricotta (Void) value"))
	}

	counterLeft := len(a.variablesToAssign)
	counterRight := 1

	tupleValue, isTuple := (*result.Value).(*parser.TupleVariableType)

	if isTuple {
		counterRight = len((*tupleValue).Variables)
	}

	if counterLeft != counterRight {
		return parser.NullExpressionResult, errors.New(fmt.Sprintf("Impossible to unpack %v values into %v variables", counterRight, counterLeft))
	}

	for i, variableName := range a.variablesToAssign {
		err := assignSingeVariable(variableName, tupleValue.Variables[i], globalContext, localContext)

		if err != nil {
			return parser.NullExpressionResult, err
		}
	}

	return parser.VoidExpressionResult, nil
}

func assignSingeVariable(variableName string, valueToAssign parser.VariableContainer, globalContext *parser.Context, localContext *parser.Context) error {

	variable, ok := parser.GetVariable(localContext, globalContext, variableName)

	if !ok {
		return errors.New("Variable with name: " + variableName + " not found in current context")
	}

	if variable.Value.GetVariableType() == parser.Tuple {
		panic("The parser should not allow to create a tuple variable")
	}

	if variable.Value.GetVariableType() != valueToAssign.GetVariableType() {
		return errors.New(fmt.Sprintf("Unsupported assignment between values with type: %v and %v", variable.Value.GetVariableType(), valueToAssign.GetVariableType()))
	}

	variable.Value = valueToAssign

	return nil
}
