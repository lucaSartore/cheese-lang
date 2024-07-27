package parser

import (
	"cheese-lang/internal/expressions"
)

type ParserCache struct {
	// parsing has various stages, (one for each type of expression)
	// when an expression is parsed, the parser try to parse each kind of expression one by one.
	// we store the stage here to avoid trying to parse the same expression multiple times
	IndexToParserStage map[int]int
	// this cache the result of the parsing of each expression
	IndexToExpression map[int]expressions.Expression
}

func MakeParserCache() ParserCache {
	return ParserCache{IndexToParserStage: make(map[int]int), IndexToExpression: make(map[int]expressions.Expression)}
}

func (p *ParserCache) GetExpression(index int) (expressions.Expression, bool) {
	expression, ok := p.IndexToExpression[index]
	return expression, ok
}

func (p *ParserCache) SetExpression(index int, expression expressions.Expression) {
	p.IndexToExpression[index] = expression
}

func (p *ParserCache) GetParserStage(index int) (int, bool) {
	stage, ok := p.IndexToParserStage[index]
	return stage, ok
}

func (p *ParserCache) SetParserStage(index int, stage int) {
	p.IndexToParserStage[index] = stage
}
