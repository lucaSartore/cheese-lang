package parser
import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
)

func (p *Parser) parseVariable(_ bool) ParserResult {
    token, err := p.ExpectReedNextToken(tokenizer.Identifier)

    // this item is not an identifier
    if err != nil {
        return p.MakeParserResult(nil, nil)
    }

    expression := expressions.VariableExpression{Name: token.TokenVal}
    return p.MakeParserResult(&expression, nil)
}
