package parser

import (
	"cheese-lang/internal/tokenizer"
	"errors"
	"fmt"
)

type Parser struct {
	Tokens        []tokenizer.Token
	Index         int // point to the first token that has't been parsed yet
	End           int // point to the first token that is not part of the expression that has been parsed
	GlobalContext *Context
}

func MakeParser(tokens []tokenizer.Token, globalContext *Context) Parser {
	return Parser{Tokens: tokens, Index: 0, GlobalContext: globalContext}
}

func (p *Parser) NextTokenMatch(token tokenizer.TokenType) bool {
	if p.Index == p.End {
		return false
	}
	return p.Tokens[p.Index].TokenType == token
}

func (p *Parser) ReadNextToken() (tokenizer.Token, error) {
	if p.Index == p.End {
		return tokenizer.Token{}, errors.New("trying to read a token, but reached end of file")
	}
	p.Index++
	return p.Tokens[p.Index-1], nil
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
	for i := p.Index; i < p.End; i++ {
		if p.Tokens[i].TokenType == tokenType {
			return i
		}
	}
	return -1
}

// FindNextMatchInMany finds the next token that matches any of the tokenTypes

func (p *Parser) FindMatchingParenthesis(start int) int {
	if start < p.Index || start >= p.End {
		panic("start index out of range")
	}
	if p.Tokens[p.Index].TokenType != tokenizer.OpenBracket {
		panic("start index is not an open bracket")
	}

	toClose := 1
	index := start + 1

	for toClose != 0 && index < p.End {
		if p.Tokens[index].TokenType == tokenizer.OpenBracket {
			toClose++
		}
		if p.Tokens[index].TokenType == tokenizer.CloseBracket {
			toClose--
		}
		index++
	}

	if toClose != 0 {
		return -1
	}

	return index - 1
}

func (p *Parser) NewSplicedParser(newIndex int, newEnd int) Parser {
	newParser := MakeParser(p.Tokens, p.GlobalContext)
	newParser.Index = newIndex
	return newParser
}

func (p *Parser) GetNextParserRegion() (Parser, bool) {
	index := p.FindNextMatch(tokenizer.SemiColon)
	if index == -1 {
		return Parser{}, false
	}
	to_return := p.NewSplicedParser(p.Index, index)
	p.Index = index + 1
	return to_return, true
}

// func ParseAnything(tokens []tokenizer.Token, index int, isLocalContext bool, globalContext *Context) ParserResult {
// 	return UnsuccessfulParserResult
// }
