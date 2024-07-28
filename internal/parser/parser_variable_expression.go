package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
	"fmt"
)

func (p *Parser) parseVariableDeclaration(global bool) ParserResult {
	tokens := []tokenizer.TokenType{tokenizer.MozzarellaType, tokenizer.GorgonzolaType, tokenizer.MilkType, tokenizer.ParmesanType}

	for _, token := range tokens {
		if p.NextTokenMatch(token) {
			p.ReadNextToken()
			identifier, err := p.ExpectReedNextToken(tokenizer.Identifier)
			if err != nil {
				return p.MakeSuccessfulResult(nil)
			}
			_, err = p.ExpectReedNextToken(tokenizer.AssignOperator)
			if err != nil {
				return p.MakeSuccessfulResult(nil)
			}

			exprResult := p.ParseAnything(global)

			if exprResult.Error != nil {
				return exprResult
			}

			if exprResult.Expression == nil {
				return p.MakeErrorResult(fmt.Errorf("expected expression after assignment operator"))
			}

			var expression expressions.Expression = &expressions.VariableDeclarationExpression{
				Type:     tokenToType(token),
				Name:     identifier.TokenVal,
				ToAssign: exprResult.Expression,
				Global:   global,
			}

			return p.MakeSuccessfulResult(expression)
		}
	}
	return p.MakeUnsuccessfulResult()
}
