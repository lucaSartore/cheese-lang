package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
	"github.com/go-errors/errors"
)

// reminder: the brake expression is called drain in the language
func (p *Parser) parserCuddleExpression (global bool) ParserResult {

	_, err := p.ExpectReedNextToken(tokenizer.CurdleKeyword)

	if err != nil {
		return p.MakeUnsuccessfulResult()
	}

	if global {
		return p.MakeErrorResult(errors.Errorf("cuddle expressions are not allowed in global scope"))
	}

    code := p.parseCodeExpression(global)

	if code.Error != nil {
        return code
	}
    
    return MakeParserResult(&expressions.CuddleExpression{CodeInside: code.Expression}, nil) 
}
