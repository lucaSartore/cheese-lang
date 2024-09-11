package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
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

    args := make([]expressions.Expression,0)

    for !p.NextTokenMatch(tokenizer.CloseBracket){
        arg := p.ParseAnything(global)
        if arg.Error != nil {
            return p.MakeErrorResult(arg.Error)
        }
        args = append(args, arg.Expression)
        
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

    function_call := expressions.FunctionCallExpression{
        FunctionToCall: function_name.TokenVal,
        Args: args,
    }

    return p.MakeSuccessfulResult(&function_call)
}
