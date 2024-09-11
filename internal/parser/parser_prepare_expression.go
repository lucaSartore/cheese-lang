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

    returns_expressions := make([]expressions.Expression,0)

    for {
        expression := p.ParseAnything(global)

        if expression.Error != nil {
            return p.MakeErrorResult(expression.Error)
        }

        if expression.Expression == nil {
            break
        }
        
        _, _ = p.ExpectReedNextToken(tokenizer.Comma)
    }

    var return_expression expressions.Expression

    if len(returns_expressions) == 0 {
        return_expression = &expressions.LiteralExpression{Literal: &expressions.RicottaVariable{}}
    }else if len(returns_expressions) == 1 {
        return_expression = returns_expressions[1]
    }else {
        return_expression = &expressions.TupleVariableType
    }

	if global {
		return p.MakeErrorResult(errors.Errorf("drain expressions are not allowed in global scope"))
	}

    returns := expressions.ReturnExpression{}
}
