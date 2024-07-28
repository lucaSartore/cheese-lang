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
		p.IndexTmp--
		return t, e
	}
	if t.TokenType != token {
		p.IndexTmp--
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
	if start < p.Index || start >= p.End {
		panic("start index out of range")
	}
	if p.Tokens[p.Index].TokenType != open {
		panic("start index is not an open bracket")
	}

	toClose := 1
	index := start + 1

	for toClose != 0 && index < p.End {
		if p.Tokens[index].TokenType == open {
			toClose++
		}
		if p.Tokens[index].TokenType == close {
			toClose--
		}
		index++
	}

	if toClose != 0 {
		return -1
	}

	return index - 1
}

func (p *Parser) FindMatchingBrackets(start int) int {
	return p.FindNextMatchingCupuleOfTokens(start, tokenizer.OpenBracket, tokenizer.CloseBracket)
}

func (p *Parser) FindMatchingCurlingBrackets(start int) int {
	return p.FindNextMatchingCupuleOfTokens(start, tokenizer.OpenCurlyBracket, tokenizer.CloseBracket)
}
