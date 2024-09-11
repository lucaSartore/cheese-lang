package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
	"github.com/go-errors/errors"
)

// reminder: the brake expression is called drain in the language
func (p *Parser) parserFunctionCallExpression (global bool) ParserResult {

	function_name, err := p.ExpectReedNextToken(tokenizer.Identifier)
    
    if err != nil{
        return p.MakeUnsuccessfulResult()
    }
    
    _, err = p.ExpectReedNextToken(tokenizer.OpenBracket)


    if err != nil{
        return p.MakeUnsuccessfulResult()
    }

    expressoins := make([]expressions.Expression,0)

    for !p.NextTokenMatch(tokenizer.CloseBracket){
        arg := p.ParseAnything(global)
        if arg.Error != nil {
            return p.MakeErrorResult(arg.Error)
        }
        expressoins = append(expressoins, arg.Expression)
        
        if p.NextTokenMatch(tokenizer.CloseBracket){
            continue
        }

        _, err := p.ExpectReedNextToken(tokenizer.Comma)

        if err != nil {
            return p.MakeErrorResult(err)
        }
    }

    _, err = p.ExpectReedNextToken(tokenizer.CloseBracket)

    if err != nil {
        panic("Assertion 1 failed in parser function call")
    }

    

}
