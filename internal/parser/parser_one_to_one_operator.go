package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/expressions/operators"
	"cheese-lang/internal/tokenizer"

	"github.com/go-errors/errors"
)

type OneToOneOperatorTuple struct {
	OperatorToken tokenizer.TokenType
	OperatorFunc  func(expressions.VariableContainer) (expressions.VariableContainer, error)
}

var OneToOneOperators = []OneToOneOperatorTuple{
	{tokenizer.NotOperator, operators.NotOperator},
    {tokenizer.SubOperator, operators.NegateOperator},
}


var oneToOneOperatorTokens []tokenizer.TokenType = Map(OneToOneOperators, func(v OneToOneOperatorTuple) tokenizer.TokenType { return v.OperatorToken })

func (p *Parser) parseOneToOneOperator(global bool) ParserResult {
    
    if !p.NextTokenMatchMultiple(oneToOneOperatorTokens){
        return p.MakeUnsuccessfulResult()
    }
    
    operator, err := p.ReadNextToken()
    
    if err != nil {
        panic("assertion 1 fail in parse one to one operator")
    }
    
    var function func(expressions.VariableContainer) (expressions.VariableContainer, error) = nil

    for _, op := range OneToOneOperators{
        if operator.TokenType == op.OperatorToken{
            function = op.OperatorFunc
            break
        }
    }

    if function == nil {
        return p.MakeUnsuccessfulResult()
    }

    value := p.ParseAnything(global)
    
    if value.Error != nil {
        return value
    }

    if value.Expression == nil {
        return p.MakeErrorResult(errors.Errorf("expected expression after %s operator", operator.TokenType.String()))
    }
    
    return p.MakeSuccessfulResult(&expressions.OneToOneOperator{Operator: function, Value: value.Expression})
}
