package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
	"strconv"
)

func (p *Parser) ParseLiteralExpression(global bool) ParserResult {
	result, err := p.ExpectReedNextToken(tokenizer.GorgonzolaLiteral)
	if err == nil {
		floatLiteral, convErr := strconv.ParseFloat(result.TokenVal, 64)
		if convErr != nil {
			panic("Parser Error: wrong float literal expression should have been catch by tokenizer.")
		}
		variable := expressions.GorgonzolaVariable{Value: floatLiteral}
		expression := expressions.LiteralExpression{Literal: &variable}
		return MakeParserResult(true, &expression, nil)
	}
	result, err = p.ExpectReedNextToken(tokenizer.ParmesanLiteral)
	if err == nil {
		intLiteral, convErr := strconv.ParseInt(result.TokenVal, 10, 32)
		if convErr != nil {
			panic("Parser Error: wrong float literal expression should have been catch by tokenizer.")
		}
		variable := expressions.ParmesanVariable{Value: int(intLiteral)}
		expression := expressions.LiteralExpression{Literal: &variable}
		return MakeParserResult(true, &expression, nil)
	}
	result, err = p.ExpectReedNextToken(tokenizer.MozzarellaLiteral)
	if err == nil {
		variable := expressions.MozzarellaVariable{Value: result.TokenVal}
		expression := expressions.LiteralExpression{Literal: &variable}
		return MakeParserResult(true, &expression, nil)
	}
	result, err = p.ExpectReedNextToken(tokenizer.FreshMilk)
	if err == nil {
		variable := expressions.MilkVariable{Value: true}
		expression := expressions.LiteralExpression{Literal: &variable}
		return MakeParserResult(true, &expression, nil)
	}
	result, err = p.ExpectReedNextToken(tokenizer.SpoiledMilk)
	if err == nil {
		variable := expressions.MilkVariable{Value: false}
		expression := expressions.LiteralExpression{Literal: &variable}
		return MakeParserResult(true, &expression, nil)
	}
	return MakeParserResult(false, nil, nil)
}
