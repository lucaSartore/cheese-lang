package test

import (
	"cheese-lang/internal/parser"
	"cheese-lang/internal/parser/expressions"
	"cheese-lang/internal/parser/operators"

	// "fmt"
	"testing"
)

func VerifySingleExpression(expressionToTest parser.Expression, t *testing.T, callerName string, context *parser.Context,
	expectedValue parser.VariableContainer, expectReturn bool, expectBrake bool, expectError bool, expectedError error) {

	result, err := expressionToTest.Evaluate(context, context)

	// the function dose not return error as it should
	if expectError && err == nil {
		t.Errorf("%v: Expected error, but got none", callerName)
		return
	}

	// the function returns error but it should not
	if !expectError && err != nil {
		t.Errorf("%v: Unexpected error: %v", callerName, err)
		return
	}

	// the returned error is different from the expected one
	if expectError && err != nil {
		if err.Error() != expectedError.Error() {
			t.Errorf("%v: Expected error: %v, but got: %v", callerName, expectedError, err)
		}
		return
	}

	// check for brake and return
	if result.Return != expectReturn {
		t.Errorf("%v: Expected return: %v, but got: %v", callerName, expectReturn, result.Return)
		return
	}

	if result.Brake != expectBrake {
		t.Errorf("%v: Expected brake: %v, but got: %v", callerName, expectBrake, result.Brake)
		return
	}

}

func VerifyValueEquivalence(v1 parser.VariableContainer, v2 parser.VariableContainer) bool {
	if v1.GetVariableType() != v2.GetVariableType() {
		return false
	}
	switch v1.GetVariableType() {
	case parser.Mozzarella:
		return v1.(*parser.MozzarellaVariable).Value == v2.(*parser.MozzarellaVariable).Value
	case parser.Parmesan:
		return v1.(*parser.ParmesanVariable).Value == v2.(*parser.ParmesanVariable).Value
	case parser.Gorgonzola:
		return v1.(*parser.GorgonzolaVariable).Value == v2.(*parser.GorgonzolaVariable).Value
	case parser.Milk:
		return v1.(*parser.MilkVariable).Value == v2.(*parser.MilkVariable).Value
	case parser.Ricotta:
		return true
	case parser.Tuple:
		t1 := v1.(*parser.TupleVariableType)
		t2 := v2.(*parser.TupleVariableType)
		if len(t1.Variables) != len(t2.Variables) {
			return false
		}
		for i := 0; i < len(t1.Variables); i++ {
			if !VerifyValueEquivalence(t1.Variables[i], t2.Variables[i]) {
				return false
			}
		}
		return true
	default:
		panic("Unsupported variable type")
	}
}

func TestEqualOperator(t *testing.T) {

	context := parser.MakeContext()
	context.AddVariable(parser.MakeVariable("m1", &parser.MozzarellaVariable{Value: "ABC"}))
	context.AddVariable(parser.MakeVariable("m2", &parser.MozzarellaVariable{Value: "ABC"}))
	context.AddVariable(parser.MakeVariable("m3", &parser.MozzarellaVariable{Value: "123"}))

	m1 := expressions.VariableExpression{Name: "m1"}
	m2 := expressions.VariableExpression{Name: "m2"}
	m3 := expressions.VariableExpression{Name: "m3"}

	exp := expressions.TwoToOneOperatorExpression{
		Operator:   operators.EqualOperator,
		RightValue: &m1,
		LeftValue:  &m2,
	}

	VerifySingleExpression(&exp, t, "TestEqualOperator", &context, &parser.MilkVariable{Value: true}, false, false, false, nil)

	exp = expressions.TwoToOneOperatorExpression{
		Operator:   operators.EqualOperator,
		RightValue: &m1,
		LeftValue:  &m3,
	}

	VerifySingleExpression(&exp, t, "TestEqualOperator", &context, &parser.MilkVariable{Value: false}, false, false, false, nil)
}

func TestAssignment(t *testing.T) {

	context := parser.MakeContext()

	context.AddVariable(parser.MakeVariable("m1", &parser.MozzarellaVariable{Value: ""}))
	context.AddVariable(parser.MakeVariable("m2", &parser.MozzarellaVariable{Value: ""}))

	value := expressions.AssignExpression{
		VariablesToAssign: []string{"m1"},
		ValueToAssign: &expressions.LiteralExpression{
			Literal: &parser.MozzarellaVariable{Value: "ABC"},
		},
	}

	VerifySingleExpression(&value, t, "testAssignment", &context, &parser.RicottaVariable{}, false, false, false, nil)

	newM1, _ := context.GetVariable("m1")
	if newM1.Value.(*parser.MozzarellaVariable).Value != "ABC" {
		t.Errorf("testAssignment: Expected m1 to be ABC, but got: %v", newM1.Value.(*parser.MozzarellaVariable).Value)
	}

	value2 := expressions.AssignExpression{
		VariablesToAssign: []string{"m1", "m2"},
		ValueToAssign: &expressions.LiteralExpression{
			Literal: &parser.TupleVariableType{
				Variables: []parser.VariableContainer{
					&parser.MozzarellaVariable{Value: "DEF"},
					&parser.MozzarellaVariable{Value: "GHI"},
				},
			},
		},
	}

	VerifySingleExpression(&value2, t, "testAssignment", &context, &parser.RicottaVariable{}, false, false, false, nil)

	newM1, _ = context.GetVariable("m1")
	if newM1.Value.(*parser.MozzarellaVariable).Value != "DEF" {
		t.Errorf("testAssignment: Expected m1 to be DEF, but got: %v", newM1.Value.(*parser.MozzarellaVariable).Value)
	}

	newM2, _ := context.GetVariable("m2")
	if newM2.Value.(*parser.MozzarellaVariable).Value != "GHI" {
		t.Errorf("testAssignment: Expected m2 to be GHI, but got: %v", newM2.Value.(*parser.MozzarellaVariable).Value)
	}
}
