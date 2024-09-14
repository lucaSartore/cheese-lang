package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
	"github.com/go-errors/errors"
)

// reminder: the brake expression is called drain in the language
func (p *Parser) parseBrakeExpression (global bool) ParserResult {

	_, err := p.ExpectReedNextToken(tokenizer.DrainKeyword)

	if err != nil {
		return p.MakeUnsuccessfulResult()
	}

	if global {
		return p.MakeErrorResult(errors.Errorf("drain expressions are not allowed in global scope"))
	}

    return p.MakeParserResult(&expressions.BrakeExpression{}, nil) 
}
