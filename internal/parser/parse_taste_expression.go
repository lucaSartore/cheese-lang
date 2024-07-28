package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
	"fmt"
)

func (p *Parser) parseTasteExpression(global bool) ParserResult {

	_, err := p.ExpectReedNextToken(tokenizer.TasteKeyword)

	if err != nil {
		return p.MakeUnsuccessfulResult()
	}

	if global {
		return p.MakeErrorResult(fmt.Errorf("taste expressions are not allowed in global scope"))
	}

	conditionResult := p.ParseAnything(global)

	if conditionResult.Error != nil {
		return conditionResult
	}

	if conditionResult.Expression == nil {
		return p.MakeErrorResult(fmt.Errorf("expected expression after taste keyword"))
	}

	blockResult := p.parseCodeExpression(global)

	if blockResult.Error != nil {
		return blockResult
	}

	if blockResult.Expression == nil {
		return p.MakeErrorResult(fmt.Errorf("expected block after taste expression"))
	}

	return p.MakeSuccessfulResult(&expressions.TasteExpression{Condition: conditionResult.Expression, Code: blockResult.Expression})
}
