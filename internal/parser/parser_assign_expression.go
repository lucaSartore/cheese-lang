package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"

	"github.com/go-errors/errors"
)

func (p *Parser) parseAssignExpression(global bool) ParserResult {

	var identifiers []string = make([]string, 0)

	if !p.NextTokenMatch(tokenizer.Identifier) {
		return p.MakeUnsuccessfulResult()
	}

	for {
		identifier, err := p.ExpectReedNextToken(tokenizer.Identifier)

		if err != nil {
			return p.MakeErrorResult(errors.Errorf("expected identifier"))
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

        _, err = p.ExpectReedNextToken(tokenizer.Comma)
        if err != nil {
            panic("assertion 1 fail in parser assign expression")
        }
	}

	rightValueResult := p.ParseAnything(global)

	if rightValueResult.Error != nil {
		return rightValueResult
	}

	if rightValueResult.Expression == nil {
		return p.MakeErrorResult(errors.Errorf("expected expression after assign operator"))
	}

	return p.MakeSuccessfulResult(&expressions.AssignExpression{VariablesToAssign: identifiers, ValueToAssign: rightValueResult.Expression})
}
