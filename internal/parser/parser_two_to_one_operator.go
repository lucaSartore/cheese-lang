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

var operatorTokens []tokenizer.TokenType = Map(Operators, func(v OperatorTuple) tokenizer.TokenType { return v.OperatorToken })

func (p *Parser) parseTwoToOneOperator(global bool) ParserResult {

	leftValueResult := p.ParseAnything(global)

	if leftValueResult.Error != nil {
		return leftValueResult
	}

	if leftValueResult.Expression == nil {
		return p.MakeUnsuccessfulResult()
	}

	leftValue := leftValueResult.Expression

	// in this case this is not a two to one operator, however we can return the left value
	if p.NextTokenMatchMultiple(operatorTokens) == false {
		return leftValueResult
	}

	token, err := p.ReadNextToken()

	if err != nil {
		panic("assertion 1 fail in parseTwoToOneOperator")
	}

	var operator *OperatorTuple = nil

	for _, op := range Operators {
		if op.OperatorToken == token.TokenType {
			operator = &op
			break
		}
	}

	if operator == nil {
		panic("assertion 2 fail in parseTwoToOneOperator")
	}

	rightValueResult := p.ParseAnything(global)

	if rightValueResult.Error != nil {
		return rightValueResult
	}

	if rightValueResult.Expression == nil {
		return p.MakeErrorResult(fmt.Errorf("expected value after operator %s", token.TokenType.String()))
	}

	rightValue := rightValueResult.Expression

	return p.MakeSuccessfulResult(&expressions.TwoToOneOperatorExpression{LeftValue: leftValue, RightValue: rightValue, Operator: operator.OperatorFunc})
}
