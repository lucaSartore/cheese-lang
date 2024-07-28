package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
	"fmt"
)

func (p *Parser) parseAssignExpression(global bool) ParserResult {

	var identifiers []string = make([]string, 0)

	for {
		identifier, err := p.ExpectReedNextToken(tokenizer.Identifier)

		if err != nil {
			return p.MakeErrorResult(fmt.Errorf("expected identifier"))
		}

		identifiers = append(identifiers, identifier.TokenVal)

		// assign operator is found, meaning that the parsing of the identifiers is done
		if p.NextTokenMatch(tokenizer.AssignOperator) {
			p.ExpectReedNextToken(tokenizer.AssignOperator)
			break
		}

		// if there is no comma, or assign operator, this is not an assign expression
		if !p.NextTokenMatch(tokenizer.Comma) {
			return p.MakeUnsuccessfulResult()
		}
	}

	rightValueResult := p.ParseAnything(global)

	if rightValueResult.Error != nil {
		return rightValueResult
	}

	if rightValueResult.Expression == nil {
		return p.MakeErrorResult(fmt.Errorf("expected expression after assign operator"))
	}

	return p.MakeSuccessfulResult(&expressions.AssignExpression{VariablesToAssign: identifiers, ValueToAssign: rightValueResult.Expression})
}
