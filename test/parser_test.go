package test

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/parser"
	"cheese-lang/internal/tokenizer"
	"testing"
)

func DoTestOnString(code string, variableToTest string, expectedValue expressions.VariableContainer, t *testing.T) {

	tokens, err := tokenizer.Tokenize(code)

	if err != nil {
		t.Errorf("Error while tokenizing: %v", err)
		return
	}

	context := expressions.MakeContext()

	parser := parser.MakeParser(tokens)

	returnValue := parser.ParseAnything(false)

	if returnValue.Error != nil {
		t.Errorf("Error while parsing: %v", returnValue.Error)
		return
	}

	if returnValue.Expression == nil {
		t.Errorf("Expect to parse an expression, but got nil")
		return
	}

	_, err = returnValue.Expression.Evaluate(&context, &context)

	if err != nil {
		t.Errorf("Error while evaluating: %v", err)
		return
	}

	variable, ok := context.GetVariable(variableToTest)
	if !ok {
		t.Errorf("Variable not found")
		return
	}

	if !VerifyValueEquivalence(variable.Value, expectedValue) {
		t.Errorf("Expected value: %v, got: %v", expectedValue, variable.Value)
		return
	}
}

func TestMozzarellaVariableParser(t *testing.T) {
	code := "Mozzarella x = \"hello\" + \" \"  + \"world\""
	DoTestOnString(code, "x", &expressions.MozzarellaVariable{Value: "hello world"}, t)
}

func TestParmesanVariableParser(t *testing.T) {
	code := "Parmesan x = (100-((5 + 3) * 2)) * (11 - 1)"
	DoTestOnString(code, "x", &expressions.ParmesanVariable{Value: 840}, t)
}

func TestCodeBlock(t *testing.T) {
	code := `
	{
		Mozzarella x = "hello";
		Mozzarella y = "world";
		Mozzarella z = "undefined";
		z = x + " " + y;
	}
	`
	DoTestOnString(code, "z", &expressions.MozzarellaVariable{Value: "hello world"}, t)
}
