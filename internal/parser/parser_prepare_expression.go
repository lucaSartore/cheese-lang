package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
	"github.com/go-errors/errors"
)

// reminder: the brake expression is called drain in the language
func (p *Parser) parseReturnExpression (global bool) ParserResult {

	_, err := p.ExpectReedNextToken(tokenizer.PrepareKeyword)

	if err != nil {
		return p.MakeUnsuccessfulResult()
	}

    if global{
        return p.MakeErrorResult(errors.New("impossible to prepare in global context"))
    }

    returns_expressions := make([]expressions.Expression,0)

    for {
        expression := p.ParseAnything(global)

        if expression.Error != nil {
            return p.MakeErrorResult(expression.Error)
        }

        if expression.Expression == nil {
            break
        }

        returns_expressions = append(returns_expressions, expression.Expression)
        
        _, _ = p.ExpectReedNextToken(tokenizer.Comma)
    }

    return p.MakeSuccessfulResult(&expressions.ReturnExpression{Expressions: returns_expressions})
    
}
