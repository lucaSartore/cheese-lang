package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
	"fmt"
	// "fmt"
)

func (p *Parser) parseCodeExpression(global bool) ParserResult {

	if !p.NextTokenMatch(tokenizer.OpenBracket) {
		return p.MakeUnsuccessfulResult()
	}

	closeBracketIndex := p.FindMatchingCurlingBrackets(p.IndexTmp)

	if closeBracketIndex == -1 {
		return p.MakeErrorResult(fmt.Errorf("could not find closing bracket of code expression"))
	}

	var expressionsList []expressions.Expression = make([]expressions.Expression, 0)

	for {

		parseResult := p.ParseAnything(global)

		if parseResult.Error != nil {
			return p.MakeErrorResult(parseResult.Error)
		}

		if parseResult.Expression == nil {
			break
		}

		newExpression := parseResult.Expression

		expressionsList = append(expressionsList, newExpression)

		_, isCodeExpression := newExpression.(*expressions.CodeExpression)
		_, isTasteExpression := newExpression.(*expressions.TasteExpression)
		_, isCuddleExpression := newExpression.(*expressions.CuddleExpression)

		if !isCodeExpression && !isTasteExpression && !isCuddleExpression {
			_, err := p.ExpectReedNextToken(tokenizer.SemiColon)
			if err != nil {
				return p.MakeErrorResult(err)
			}
		}

		if p.IndexTmp == closeBracketIndex {
			break
		}
	}

	_, err := p.ExpectReedNextToken(tokenizer.CloseBracket)

	if err != nil {
		panic("assertion 1 fail in parseCodeExpression")
	}

	return p.MakeSuccessfulResult(&expressions.CodeExpression{Expressions: expressionsList})
}
