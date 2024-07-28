package parser

import "cheese-lang/internal/expressions"

type ParserResult struct {
	Expression expressions.Expression // return of the expression result
	Error      error                  // error because of malformed input
}

func MakeParserResult(expression expressions.Expression, error error) ParserResult {
	return ParserResult{Expression: expression, Error: error}
}

func (p *Parser) MakeSuccessfulResult(expression expressions.Expression) ParserResult {
	advancedTokens := p.IndexTmp - p.Index
	if advancedTokens == 0 {
		panic("Trying to make a successful result without advancing the index")
	}
	return MakeParserResult(expression, nil)
}

func (p *Parser) MakeUnsuccessfulResult() ParserResult {
	return MakeParserResult(nil, nil)
}

func (p *Parser) MakeErrorResult(err error) ParserResult {
	return MakeParserResult(nil, err)
}
