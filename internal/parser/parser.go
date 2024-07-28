package parser

import (
	"cheese-lang/internal/tokenizer"
)

type Parser struct {
	Tokens   []tokenizer.Token
	Index    int // point to the first token that has't been parsed yet
	IndexTmp int // temporary index used to store where the parser is at, in the case the succuss of a parsing is not sure yet
	End      int // point to the first token that is not part of the expression that has been parsed
}

func MakeParser(tokens []tokenizer.Token) Parser {
	return Parser{Tokens: tokens, Index: 0, IndexTmp: 0, End: len(tokens)}
}

func (p *Parser) NewSplicedParser(newIndex int, newEnd int) Parser {
	newParser := MakeParser(p.Tokens)
	newParser.Index = newIndex
	newParser.IndexTmp = newIndex
	newParser.End = newEnd
	return newParser
}

func (p *Parser) GetNextParserRegion() (Parser, bool) {
	index := p.FindNextMatch(tokenizer.SemiColon)
	if index == -1 {
		return Parser{}, false
	}
	to_return := p.NewSplicedParser(p.Index, index)
	p.Index = index + 1
	p.IndexTmp = p.Index
	return to_return, true
}

type ParsingStageType int

const (
	TwoToOneOperatorStage    ParsingStageType = 0
	VariableDeclarationStage ParsingStageType = 1
	LiteralExpressionStage   ParsingStageType = 2
	BracketsParsingStage     ParsingStageType = 3
)

var AllParsingStages = []ParsingStageType{TwoToOneOperatorStage, VariableDeclarationStage, LiteralExpressionStage, BracketsParsingStage}

func (p *Parser) ExecuteParsingStage(stage ParsingStageType, global bool) ParserResult {
	switch stage {
	case TwoToOneOperatorStage:
		return p.parseTwoToOneOperator(global)
	case VariableDeclarationStage:
		return p.parseVariableDeclaration(global)
	case LiteralExpressionStage:
		return p.ParseLiteralExpression(global)
	case BracketsParsingStage:
		return p.parseBracketExpression(global)
	default:
		panic("Unknown parsing stage")
	}
}

func (p *Parser) ParseBySkippingStages(global bool, stagesToSkip []ParsingStageType) ParserResult {
	for _, stage := range AllParsingStages {

		if contains(stagesToSkip, stage) {
			continue
		}

		indexTmpPre := p.IndexTmp

		result := p.ExecuteParsingStage(stage, global)

		if result.Error != nil {
			return result
		}

		if result.Expression != nil {
			return result
		}

		p.IndexTmp = indexTmpPre
	}
	return p.MakeUnsuccessfulResult()
}

func (p *Parser) ParseAnything(global bool) ParserResult {
	return p.ParseBySkippingStages(global, []ParsingStageType{})
}
