package parser

import (
	"cheese-lang/internal/parser/expressions"
	"cheese-lang/internal/tokenizer"
	"fmt"
)

func tokenToType(token tokenizer.TokenType) VariableType {
	switch token {
	case tokenizer.MozzarellaType:
		return Mozzarella
	case tokenizer.GorgonzolaType:
		return Gorgonzola
	case tokenizer.MilkType:
		return Milk
	case tokenizer.ParmesanType:
		return Parmesan
	}
	panic(fmt.Sprintf("Unknown token type %v", token))
}

func (p *Parser) parseVariableDeclaration(global bool) ParserResult {
	tokens := []tokenizer.TokenType{tokenizer.MozzarellaType, tokenizer.GorgonzolaType, tokenizer.MilkType, tokenizer.ParmesanType}

	for _, token := range tokens {
		if p.NextTokenMatch(token) {
			p.ReadNextToken()
			identifier, err := p.ExpectReedNextToken(tokenizer.Identifier)
			if err != nil {
				return MakeParserResult(true, nil, err)
			}
			_, err = p.ExpectReedNextToken(tokenizer.AssignOperator)
			if err != nil {
				return MakeParserResult(true, nil, err)
			}

			exprResult := p.ParseAnything(global)

			if exprResult.Error != nil {
				return exprResult
			}

			if !exprResult.progressed {
				return MakeParserResult(false, nil, fmt.Errorf("expected expression after assignment operator"))
			}

			var expression Expression = &expressions.VariableDeclarationExpression{
				Type:     tokenToType(token),
				Name:     identifier.TokenVal,
				ToAssign: exprResult.Expression,
				Global:   global,
			}

			return MakeParserResult(true, expression, nil)
		}
	}
	return MakeParserResult(false, nil, nil)
}
