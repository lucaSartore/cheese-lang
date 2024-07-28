package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"

	"github.com/go-errors/errors"
)

func (p *Parser) parseTasteExpression(global bool) ParserResult {

	_, err := p.ExpectReedNextToken(tokenizer.TasteKeyword)

	if err != nil {
		return p.MakeUnsuccessfulResult()
	}

	if global {
		return p.MakeErrorResult(errors.Errorf("taste expressions are not allowed in global scope"))
	}

	conditionResult := p.ParseAnything(global)

	if conditionResult.Error != nil {
		return conditionResult
	}

	if conditionResult.Expression == nil {
		return p.MakeErrorResult(errors.Errorf("expected expression after taste keyword"))
	}

	blockResult := p.parseCodeExpression(global)

	if blockResult.Error != nil {
		return blockResult
	}

	if blockResult.Expression == nil {
		return p.MakeErrorResult(errors.Errorf("expected block after taste expression"))
	}

	return p.MakeSuccessfulResult(&expressions.TasteExpression{Condition: conditionResult.Expression, Code: blockResult.Expression})
}
