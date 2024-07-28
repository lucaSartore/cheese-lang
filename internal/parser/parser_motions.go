package parser

import (
	"cheese-lang/internal/tokenizer"
	"errors"
	"fmt"
)

func (p *Parser) NextTokenMatch(token tokenizer.TokenType) bool {
	if p.IndexTmp == p.End {
		return false
	}
	return p.Tokens[p.IndexTmp].TokenType == token
}

func (p *Parser) NextTokenMatchMultiple(token []tokenizer.TokenType) bool {
	for _, t := range token {
		if p.NextTokenMatch(t) {
			return true
		}
	}
	return false
}

func (p *Parser) ReadNextToken() (tokenizer.Token, error) {
	if p.IndexTmp == p.End {
		return tokenizer.Token{}, errors.New("trying to read a token, but reached end of file")
	}
	p.IndexTmp++
	return p.Tokens[p.IndexTmp-1], nil
}

func (p *Parser) ExpectReedNextToken(token tokenizer.TokenType) (tokenizer.Token, error) {
	t, e := p.ReadNextToken()
	if e != nil {
		return t, e
	}
	if t.TokenType != token {
		return t, fmt.Errorf("expected token %s, got %s", token.String(), t.TokenType.String())
	}
	return t, nil
}

// FindNextMatch finds the next token that matches the tokenType
func (p *Parser) FindNextMatch(tokenType tokenizer.TokenType) int {
	return p.findNextMatchStartingFrom(tokenType, p.IndexTmp)
}

func (p *Parser) findNextMatchStartingFrom(tokenType tokenizer.TokenType, start int) int {
	for i := start; i < p.End; i++ {
		if p.Tokens[i].TokenType == tokenType {
			return i
		}
	}
	return -1
}

func (p *Parser) FindNextMatchMultiple(tokens []tokenizer.TokenType) int {
	for i := p.IndexTmp; i < p.End; i++ {
		for _, t := range tokens {
			if p.Tokens[i].TokenType == t {
				return i
			}
		}
	}
	return -1
}

func (p *Parser) FindNextMatchingCupuleOfTokens(start int, open tokenizer.TokenType, close tokenizer.TokenType) int {
	if start < p.IndexTmp || start >= p.End {
		panic("start index out of range")
	}

	index_open := p.findNextMatchStartingFrom(open, start)

	if index_open == -1 {
		return -1
	}

	index_close := p.findNextMatchStartingFrom(close, index_open+1)

	return index_close
}

func (p *Parser) FindMatchingBrackets(start int) int {
	return p.FindNextMatchingCupuleOfTokens(start, tokenizer.OpenBracket, tokenizer.CloseBracket)
}

func (p *Parser) FindMatchingCurlingBrackets(start int) int {
	return p.FindNextMatchingCupuleOfTokens(start, tokenizer.OpenCurlyBracket, tokenizer.CloseBracket)
}
