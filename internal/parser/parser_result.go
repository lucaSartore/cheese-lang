package parser

import "cheese-lang/internal/expressions"

type ParserResult struct {
	Expression expressions.Expression // return of the expression result
	Error      error                  // error because of malformed input
    Line uint
    Colum uint
}

func (p *Parser) MakeParserResult(expression expressions.Expression, error error) ParserResult {
    index := p.IndexTmp
    if index == len(p.Tokens) {
        index--
    }
    currentToken := p.Tokens[index]
    return ParserResult{Expression: expression, Error: error, Line: currentToken.Line, Colum: currentToken.Colum}
}

func (p *Parser) MakeSuccessfulResult(expression expressions.Expression) ParserResult {
	advancedTokens := p.IndexTmp - p.Index
	if advancedTokens == 0 {
		panic("Trying to make a successful result without advancing the index")
	}
	return p.MakeParserResult(expression, nil)
}

func (p *Parser) MakeUnsuccessfulResult() ParserResult {
	return p.MakeParserResult(nil, nil)
}

func (p *Parser) MakeErrorResult(err error) ParserResult {
	return p.MakeParserResult(nil, err)
}
