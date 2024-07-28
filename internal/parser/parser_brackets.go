package parser

import (
	"cheese-lang/internal/tokenizer"
)

func (p *Parser) parseBracketExpression(global bool) ParserResult {
	_, err := p.ExpectReedNextToken(tokenizer.OpenBracket)

	// there is no an open bracket, therefore this is not a bracket expression
	if err != nil {
		return p.MakeUnsuccessfulResult()
	}

	expressionInside := p.ParseAnything(global)

	_, err = p.ExpectReedNextToken(tokenizer.CloseBracket)

	if err != nil {
		return p.MakeErrorResult(err)
	}

	return expressionInside
}
