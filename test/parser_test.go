package test

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/parser"
	"cheese-lang/internal/tokenizer"
	"testing"
)

func TestMozzarellaVariableParser(t *testing.T) {
	testStr := "Mozzarella x = \"hello\" + \" \"  + \"world\""
	tokens, err := tokenizer.Tokenize(testStr)

	if err != nil {
		t.Errorf("Error while tokenizing: %v", err)
		return
	}

	context := expressions.MakeContext()

	parser := parser.MakeParser(tokens)

	returnValue := parser.ParseAnything(true)

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

	variable, ok := context.GetVariable("x")
	if !ok {
		t.Errorf("Variable not found")
		return
	}

	if variable.Value.GetVariableType() != expressions.Mozzarella {
		t.Errorf("Expected mozzarella variable")
		return
	}

	value := variable.Value.(*expressions.MozzarellaVariable).Value

	if value != "hello world" {
		t.Errorf("Expected value: hello world, got: %v", value)
		return
	}
}
