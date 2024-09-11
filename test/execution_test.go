package test

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/expressions/operators"
	"testing"
)

func VerifySingleExpression(expressionToTest expressions.Expression, t *testing.T, callerName string, context *expressions.Context,
	expectedValue expressions.VariableContainer, expectReturn bool, expectBrake bool, expectError bool, expectedError error) {

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

func VerifyValueEquivalence(v1 expressions.VariableContainer, v2 expressions.VariableContainer) bool {
	if v1.GetVariableType() != v2.GetVariableType() {
		return false
	}
	switch v1.GetVariableType() {
	case expressions.Mozzarella:
		return v1.(*expressions.MozzarellaVariable).Value == v2.(*expressions.MozzarellaVariable).Value
	case expressions.Parmesan:
		return v1.(*expressions.ParmesanVariable).Value == v2.(*expressions.ParmesanVariable).Value
	case expressions.Gorgonzola:
		return v1.(*expressions.GorgonzolaVariable).Value == v2.(*expressions.GorgonzolaVariable).Value
	case expressions.Milk:
		return v1.(*expressions.MilkVariable).Value == v2.(*expressions.MilkVariable).Value
	case expressions.Ricotta:
		return true
	case expressions.Tuple:
		t1 := v1.(*expressions.TupleVariableType)
		t2 := v2.(*expressions.TupleVariableType)
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

	context := expressions.MakeContext()
	context.AddVariable(expressions.MakeVariable("m1", &expressions.MozzarellaVariable{Value: "ABC"}))
	context.AddVariable(expressions.MakeVariable("m2", &expressions.MozzarellaVariable{Value: "ABC"}))
	context.AddVariable(expressions.MakeVariable("m3", &expressions.MozzarellaVariable{Value: "123"}))

	m1 := expressions.VariableExpression{Name: "m1"}
	m2 := expressions.VariableExpression{Name: "m2"}
	m3 := expressions.VariableExpression{Name: "m3"}

	exp := expressions.TwoToOneOperatorExpression{
		Operator:   operators.EqualOperator,
		RightValue: &m1,
		LeftValue:  &m2,
	}

	VerifySingleExpression(&exp, t, "TestEqualOperator", &context, &expressions.MilkVariable{Value: true}, false, false, false, nil)

	exp = expressions.TwoToOneOperatorExpression{
		Operator:   operators.EqualOperator,
		RightValue: &m1,
		LeftValue:  &m3,
	}

	VerifySingleExpression(&exp, t, "TestEqualOperator", &context, &expressions.MilkVariable{Value: false}, false, false, false, nil)
}

func TestAssignment(t *testing.T) {

	context := expressions.MakeContext()

	context.AddVariable(expressions.MakeVariable("m1", &expressions.MozzarellaVariable{Value: ""}))
	context.AddVariable(expressions.MakeVariable("m2", &expressions.MozzarellaVariable{Value: ""}))

	value := expressions.AssignExpression{
		VariablesToAssign: []string{"m1"},
		ValueToAssign: &expressions.LiteralExpression{
			Literal: &expressions.MozzarellaVariable{Value: "ABC"},
		},
	}

	VerifySingleExpression(&value, t, "testAssignment", &context, &expressions.RicottaVariable{}, false, false, false, nil)

	newM1, _ := context.GetVariable("m1")
	if newM1.Value.(*expressions.MozzarellaVariable).Value != "ABC" {
		t.Errorf("testAssignment: Expected m1 to be ABC, but got: %v", newM1.Value.(*expressions.MozzarellaVariable).Value)
	}

	value2 := expressions.AssignExpression{
		VariablesToAssign: []string{"m1", "m2"},
		ValueToAssign: &expressions.LiteralExpression{
			Literal: &expressions.TupleVariableType{
				Variables: []expressions.VariableContainer{
					&expressions.MozzarellaVariable{Value: "DEF"},
					&expressions.MozzarellaVariable{Value: "GHI"},
				},
			},
		},
	}

	VerifySingleExpression(&value2, t, "testAssignment", &context, &expressions.RicottaVariable{}, false, false, false, nil)

	newM1, _ = context.GetVariable("m1")
	if newM1.Value.(*expressions.MozzarellaVariable).Value != "DEF" {
		t.Errorf("testAssignment: Expected m1 to be DEF, but got: %v", newM1.Value.(*expressions.MozzarellaVariable).Value)
	}

	newM2, _ := context.GetVariable("m2")
	if newM2.Value.(*expressions.MozzarellaVariable).Value != "GHI" {
		t.Errorf("testAssignment: Expected m2 to be GHI, but got: %v", newM2.Value.(*expressions.MozzarellaVariable).Value)
	}
}

func TestCuddle(t *testing.T) {

	context := expressions.MakeContext()

	context.AddVariable(expressions.MakeVariable("i", &expressions.ParmesanVariable{Value: 0}))

	cuddle := expressions.CuddleExpression{
		CodeInside: &expressions.CodeExpression{
			Expressions: []expressions.Expression{
				&expressions.AssignExpression{
					VariablesToAssign: []string{"i"},
					ValueToAssign: &expressions.TwoToOneOperatorExpression{
						Operator:   operators.AddOperator,
						LeftValue:  &expressions.VariableExpression{Name: "i"},
						RightValue: &expressions.LiteralExpression{Literal: &expressions.ParmesanVariable{Value: 1}},
					},
				},
				&expressions.TasteExpression{
					Condition: &expressions.TwoToOneOperatorExpression{
						Operator:   operators.EqualOperator,
						LeftValue:  &expressions.VariableExpression{Name: "i"},
						RightValue: &expressions.LiteralExpression{Literal: &expressions.ParmesanVariable{Value: 10}},
					},
					Code: &expressions.BrakeExpression{},
				},
			},
		},
	}

	VerifySingleExpression(&cuddle, t, "TestCuddle", &context, &expressions.RicottaVariable{}, false, false, false, nil)

	i, _ := context.GetVariable("i")
	if i.Value.(*expressions.ParmesanVariable).Value != 10 {
		t.Errorf("TestCuddle: Expected i to be 10, but got: %v", i.Value.(*expressions.ParmesanVariable).Value)
	}
}

func TestFunction(t *testing.T) {
	globalContext := expressions.MakeContext()
	localContext := expressions.MakeContext()
	globalContext.AddFunction(expressions.MakeFunction(
		"pow",
		&expressions.CodeExpression{
			Expressions: []expressions.Expression{
				&expressions.TasteExpression{
					Condition: &expressions.TwoToOneOperatorExpression{
						Operator:   operators.EqualOperator,
						LeftValue:  &expressions.VariableExpression{Name: "exponent"},
						RightValue: &expressions.LiteralExpression{Literal: &expressions.ParmesanVariable{Value: 0}},
					},
					Code: &expressions.ReturnExpression{
                        Expressions: []expressions.Expression{&expressions.LiteralExpression{Literal: &expressions.ParmesanVariable{Value: 1}}},
					},
				},
				&expressions.ReturnExpression{
                    Expressions: []expressions.Expression{
                        &expressions.FunctionCallExpression{
                            FunctionToCall: "pow",
                            Args: []expressions.Expression{
                                &expressions.VariableExpression{Name: "base"},
                                &expressions.TwoToOneOperatorExpression{
                                    Operator:   operators.SubOperator,
                                    LeftValue:  &expressions.VariableExpression{Name: "exponent"},
                                    RightValue: &expressions.LiteralExpression{Literal: &expressions.ParmesanVariable{Value: 1}},
                                },
                            },
                        },
                    },
				},
			},
		},
		[]expressions.VariableType{expressions.Parmesan, expressions.Parmesan},
		[]string{"base", "exponent"},
	))

	code := expressions.FunctionCallExpression{
		FunctionToCall: "pow",
		Args: []expressions.Expression{
			&expressions.LiteralExpression{Literal: &expressions.ParmesanVariable{Value: 5}},
			&expressions.LiteralExpression{Literal: &expressions.ParmesanVariable{Value: 3}},
		},
	}

	result, err := code.Evaluate(&globalContext, &localContext)
	if err != nil {
		t.Errorf("unexpected error in function evaluation: %v", err)
	}
	if result.Return == true {
		t.Errorf("unexpected return in function evaluation")
	}
	if result.Brake == true {
		t.Errorf("unexpected brake in function evaluation")
	}
	VerifyValueEquivalence(result.Value, &expressions.ParmesanVariable{Value: 5 * 5 * 5})
}



