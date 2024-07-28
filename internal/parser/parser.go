package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
)

type Parser struct {
	Tokens        []tokenizer.Token
	Index         int // point to the first token that has't been parsed yet
	IndexTmp      int // temporary index used to store where the parser is at, in the case the succuss of a parsing is not sure yet
	End           int // point to the first token that is not part of the expression that has been parsed
	GlobalContext *expressions.Context
}

type ParserResult struct {
	ParsedTokens int                    // number of tokens that have been parsed
	Expression   expressions.Expression // return of the expression result
	Error        error                  // error because of malformed input
}

func MakeParserResult(parsedTokens int, expression expressions.Expression, error error) ParserResult {
	return ParserResult{ParsedTokens: parsedTokens, Expression: expression, Error: error}
}

func MakeParser(tokens []tokenizer.Token, globalContext *expressions.Context) Parser {
	return Parser{Tokens: tokens, Index: 0, IndexTmp: 0, GlobalContext: globalContext}
}

func (p *Parser) NewSplicedParser(newIndex int, newEnd int) Parser {
	newParser := MakeParser(p.Tokens, p.GlobalContext)
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
)

var AllParsingStages = []ParsingStageType{TwoToOneOperatorStage, VariableDeclarationStage, LiteralExpressionStage}

func (p *Parser) ExecuteParsingStage(stage ParsingStageType, global bool) ParserResult {
	switch stage {
	case TwoToOneOperatorStage:
		return p.parseTwoToOneOperator(global)
	case VariableDeclarationStage:
		return p.parseVariableDeclaration(global)
	case LiteralExpressionStage:
		return p.ParseLiteralExpression(global)
	default:
		panic("Unknown parsing stage")
	}
}

func ParseBySkippingStages(p *Parser, global bool, stagesToSkip []ParsingStageType) ParserResult {
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

		// reset the index before trying the next stage
		p.IndexTmp = indexTmpPre
	}
	return p.MakeUnsuccessfulResult()
}

func (p *Parser) ParseAnything(global bool) ParserResult {
	return ParseBySkippingStages(p, global, []ParsingStageType{})
}

func (p *Parser) MakeSuccessfulResult(expression expressions.Expression) ParserResult {
	advancedTokens := p.IndexTmp - p.Index
	if advancedTokens == 0 {
		panic("Trying to make a successful result without advancing the index")
	}
	return MakeParserResult(advancedTokens, expression, nil)
}

func (p *Parser) MakeUnsuccessfulResult() ParserResult {
	return MakeParserResult(0, nil, nil)
}

func (p *Parser) MakeErrorResult(err error) ParserResult {
	return MakeParserResult(0, nil, err)
}
