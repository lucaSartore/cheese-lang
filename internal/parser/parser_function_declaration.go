package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"

	"github.com/go-errors/errors"
)

func (p *Parser) ParseFunctionDeclaration(global bool) ParserResult {
    _, err := p.ExpectReedNextToken(tokenizer.RecipeKeyword)
    if err != nil{
        return p.MakeUnsuccessfulResult()
    }
    
    function_name, err := p.ExpectReedNextToken(tokenizer.Identifier)

    if err != nil {
        return p.MakeErrorResult(err)
    }

    names := make([]string,0)
    types := make([]expressions.VariableType,0)

    _, err = p.ExpectReedNextToken(tokenizer.OpenBracket)

    if err != nil {
        return p.MakeErrorResult(err)
    }

    
    accetted_types := []tokenizer.TokenType{ tokenizer.MilkType, tokenizer.MozzarellaType, tokenizer.ParmesanType, tokenizer.GorgonzolaType}
    for !p.NextTokenMatch(tokenizer.CloseBracket){
        if !p.NextTokenMatchMultiple(accetted_types){
		    return p.MakeErrorResult(errors.Errorf("expected value after operator"))
        }

        arg_type_token, err := p.ReadNextToken()

        if err != nil{
            panic("assertion 1 failed in function declaration parser")
        }

        arg_name, err := p.ExpectReedNextToken(tokenizer.Identifier)
        
        if err != nil{
            return p.MakeErrorResult(err)
        }

        names = append(names, arg_name.TokenVal)

        switch arg_type_token.TokenType{
        case tokenizer.MilkType:
            types = append(types, expressions.Milk)
        case tokenizer.MozzarellaType:
            types = append(types, expressions.Mozzarella)
        case tokenizer.ParmesanType:
            types = append(types, expressions.Parmesan)
        case tokenizer.GorgonzolaType:
            types = append(types, expressions.Gorgonzola)
        default:
            panic("assertion 2 failed in function declaration parser")
        }

    }

    _, err = p.ExpectReedNextToken(tokenizer.CloseBracket)
    
    if err != nil {
        return p.MakeErrorResult(err)
    }

    code := p.parseCodeExpression(false)

    if code.Error != nil {
        return p.MakeErrorResult(code.Error)
    }

    if code.Expression == nil {
        return p.MakeErrorResult(errors.Errorf("unable to parse a valid function code"))
    }

    function := expressions.MakeFunction(function_name.TokenVal,code.Expression, types,names)
    
    expression :=  expressions.FunctionDeclaration{Function: function, Global: global}

    return p.MakeSuccessfulResult(&expression)
}
