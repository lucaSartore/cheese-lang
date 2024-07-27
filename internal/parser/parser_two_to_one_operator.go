package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/expressions/operators"
	"cheese-lang/internal/tokenizer"
	"fmt"
)

type OperatorTuple struct {
	OperatorToken tokenizer.TokenType
	OperatorFunc  func(expressions.VariableContainer, expressions.VariableContainer) (expressions.VariableContainer, error)
}

var Operators = []OperatorTuple{
	{tokenizer.AddOperator, operators.AddOperator},
	{tokenizer.SubOperator, operators.SubOperator},
	{tokenizer.DivOperator, operators.DivOperator},
	{tokenizer.MulOperator, operators.MulOperator},
	{tokenizer.EqualOperator, operators.EqualOperator},
	{tokenizer.UnEqualOperator, operators.UnEqualOperator},
	{tokenizer.OrOperator, operators.OrOperator},
	{tokenizer.AndOperator, operators.AndOperator},
	{tokenizer.ExorOperator, operators.ExorOperator},
}

func (p *Parser) parseTwoToOneOperator(global bool) ParserResult {

	leftValueResult := p.ParseAnything(global)

	if leftValueResult.Error != nil {
		return leftValueResult
	}

	if !leftValueResult.progressed {
		return leftValueResult
	}

	leftValue := leftValueResult.Expression

	token, err := p.ReadNextToken()

	if err != nil {
		return MakeParserResult(true, nil, err)
	}

	var operator *OperatorTuple = nil

	for _, op := range Operators {
		if op.OperatorToken == token.TokenType {
			operator = &op
			break
		}
	}

	if operator == nil {
		return MakeParserResult(false, nil, nil)
	}

	rightValueResult := p.ParseAnything(global)

	if rightValueResult.Error != nil {
		return rightValueResult
	}

	if !rightValueResult.progressed {
		return MakeParserResult(true, nil, fmt.Errorf("expected value after operator %s", token.TokenType.String()))
	}

	rightValue := rightValueResult.Expression

	return MakeParserResult(true, &expressions.TwoToOneOperatorExpression{LeftValue: leftValue, RightValue: rightValue, Operator: operator.OperatorFunc}, nil)
}
