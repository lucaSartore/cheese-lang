package tokenizer

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
)


func Tokenize(input string, exclude_comments bool) ([]Token, error) {
	// todo: check tat the string is made of only asci characters... utf8 characters may brake this tokenizer
	inputBuffer := bytes.NewBuffer([]byte(input))
	tokens := []Token{}

    bw := BufferWrapper {
        input: inputBuffer,
        bytesRead: 0,
        currentLine: 0,
        lastLineBegin: 0,
    }

	for {
		token, err := bw.TokenizeSingle()


		if token.TokenType != NullToken && token.TokenType != NewLineToken && (token.TokenType != Comment || !exclude_comments) {
            token.Line = bw.currentLine
            token.Colum = bw.currentColum()
			tokens = append(tokens, token)
		}

        if token.TokenType == NewLineToken{
            bw.AdvanceOneLine()
        }

		if inputBuffer.Len() == 0 {
			break
		}

		if err != nil {
            return tokens, fmt.Errorf("tokenizer failed at line %v:%v because of error %v", bw.currentLine, bw.currentColum(), err)
		}
	}


	return tokens, nil
}


type BufferWrapper struct{
    input *bytes.Buffer;
    bytesRead uint;
    currentLine uint;
    lastLineBegin uint;
}

func (b *BufferWrapper) AdvanceOneLine() {
    b.currentLine += 1
    b.lastLineBegin = b.bytesRead
}

func (b *BufferWrapper) currentColum() uint{
    return b.bytesRead - b.lastLineBegin
}

func (b *BufferWrapper) TokenizeSingle() (Token, error) {
	if b.input.Len() == 0 {
		return MakeToken(NullToken), nil
	}
	inputBytes := b.input.Bytes()

	functionToTry := []func([]byte) (Token, uint){
		AdvanceWitheSpace,
        AdvanceNewLine,
		TryReadComment,
		TryReadDualCharacterToken,
		TryReadSingleCharacterToken,
		TryReadMozzarellaLiteral,
		TryReadGorgonzolaOrParmesanLiteral,
		TryReadKeyword,
	}

	for _, function := range functionToTry {
		token, len := function(inputBytes)
		if len > 0 {
            b.bytesRead += len
			buffer := make([]byte, len)
            b.input.Read(buffer)
			return token, nil
		}
	}
	return MakeToken(NullToken), errors.New("unable to find token that match with: " + string(inputBytes))
}


func AdvanceNewLine(input []byte) (Token, uint) {
    var c1 byte = 0
    var c2 byte = 0

    if len(input) != 0{
        c1 = input[0]

        if len(input) >= 2{
            c2 = input[1]
        }
    }

    if c1 == '\n' {
	    return MakeToken(NewLineToken), 1
    }

    if c1 == '\r' && c2 == '\n'{
        return MakeToken(NewLineToken), 2
    }

	return MakeToken(NullToken), 0
}

func AdvanceWitheSpace(input []byte) (Token, uint) {
	count := uint(0)
	for _, b := range input {
		if b == ' ' || b == '\t'{
			count++
		} else {
			break
		}
	}
	return MakeToken(NullToken), count
}

// try to match with operators that are made of 2 characters
func TryReadDualCharacterToken(input []byte) (Token, uint) {
	if len(input) < 2 {
		return MakeToken(NullToken), 0
	}
	text := string(input[0:2])
	switch text {
	case ">=":
		return MakeToken(GreaterEqualOperator), 2
	case "<=":
		return MakeToken(LessEqualOperator), 2
	case "==":
		return MakeToken(EqualOperator), 2
	case "!=":
		return MakeToken(UnEqualOperator), 2
	case "&&":
		return MakeToken(AndOperator), 2
	case "||":
		return MakeToken(OrOperator), 2
	case "->":
		return MakeToken(Arrow), 2
	default:
		return MakeToken(NullToken), 0
	}
}

func TryReadSingleCharacterToken(input []byte) (Token, uint) {
	if len(input) == 0 {
		return MakeToken(NullToken), 0
	}
	text := string(input[0])

	switch text {
	case "=":
		return MakeToken(AssignOperator), 1
	case "+":
		return MakeToken(AddOperator), 1
	case "-":
		return MakeToken(SubOperator), 1
	case "*":
		return MakeToken(MulOperator), 1
	case "/":
		return MakeToken(DivOperator), 1
	case "<":
		return MakeToken(LessOperator), 1
	case ">":
		return MakeToken(GreaterOperator), 1
	case "!":
		return MakeToken(NotOperator), 1
	case "^":
		return MakeToken(ExorOperator), 1
	case "(":
		return MakeToken(OpenBracket), 1
	case ")":
		return MakeToken(CloseBracket), 1
	case "{":
		return MakeToken(OpenCurlyBracket), 1
	case "}":
		return MakeToken(CloseCurlyBracket), 1
	case ",":
		return MakeToken(Comma), 1
	case ";":
		return MakeToken(SemiColon), 1
	default:
		return MakeToken(NullToken), 0
	}
}

func TryReadComment(input []byte) (Token, uint) {

	if len(input) < 2 {
		return MakeToken(NullToken), 0
	}

	text := string(input[0:2])

	if text != "//" {
		return MakeToken(NullToken), 0
	}
	count := uint(2)

	for _, c := range input[2:] {
		count++
		if c == '\n' || c == '\r' {
			break
		}
	}
	return MakeTokenWithMessage(Comment, string(input[2:count-1])), count
}

func TryReadMozzarellaLiteral(input []byte) (Token, uint) {
	if len(input) < 2 {
		return MakeToken(NullToken), 0
	}

	if input[0] != '"' {
		return MakeToken(NullToken), 0
	}

	character_counter := uint(1)

	result := []byte{}

	terminated_correctly := false
	skip_one := false

mozzarella_literal_loop:
	for i, c := range input[1:] {

		character_counter++

		if skip_one {
			skip_one = false
			continue
		}

		if c == '"' {
			terminated_correctly = true
			break
		}

		if c == '\n' {
			break
		}

		if c == '\\' {
			if len(input) <= i+2 {
				break
			}
			next_char := input[i+2]
			switch next_char {
			case '\\':
				result = append(result, '\\')
			case 't':
				result = append(result, '\t')
			case 'n':
				result = append(result, '\n')
			case '"':
				result = append(result, '"')
			default:
				break mozzarella_literal_loop
			}

			skip_one = true
		} else {
			result = append(result, c)
		}
	}

	if !terminated_correctly {
		return MakeToken(NullToken), 0
	}

	return MakeTokenWithMessage(MozzarellaLiteral, string(result)), character_counter
}

func TryReadGorgonzolaOrParmesanLiteral(input []byte) (Token, uint) {

	chars := []byte{}

	found_separator := false

	for i, b := range input {

		if b >= '0' && b <= '9' {
			chars = append(chars, b)
		} else if i != 0 && b == '.' && !found_separator {
			// the decimal separator can't be the first character
			// and there can be at most one decimal separator in each literal
			chars = append(chars, b)
			found_separator = true
		} else {
			break
		}
	}

	if len(chars) == 0 {
		return MakeToken(NullToken), 0
	}

	// make sore the last pushed value is sot the decimal separator
	if chars[len(chars)-1] == '.' {
		return MakeToken(NullToken), 0
	}

	message := string(chars)
	len := uint(len(chars))

	if found_separator {
		return MakeTokenWithMessage(GorgonzolaLiteral, message), len
	} else {
		return MakeTokenWithMessage(ParmesanLiteral, message), len
	}
}

func IsValidKeywordStartChar(c byte) bool {
	return regexp.MustCompile("^[a-zA-Z_]*$").Match([]byte{c})
}

func IsValidKeywordChar(c byte) bool {
	return regexp.MustCompile("^[a-zA-Z0-9_]*$").Match([]byte{c})
}

func TryReadKeyword(input []byte) (Token, uint) {

	if len(input) == 0 {
		return MakeToken(NullToken), 0
	}
	if !IsValidKeywordStartChar(input[0]) {
		return MakeToken(NullToken), 0
	}

	keyword_length := uint(0)

	for _, c := range input {
		if !IsValidKeywordChar(c) {
			break
		}
		keyword_length++
	}

	keyword := string(input[0:keyword_length])

	switch keyword {
	case "Mozzarella":
		return MakeToken(MozzarellaType), keyword_length
	case "Parmesan":
		return MakeToken(ParmesanType), keyword_length
	case "Gorgonzola":
		return MakeToken(GorgonzolaType), keyword_length
	case "Milk":
		return MakeToken(MilkType), keyword_length
	case "Ricotta":
		return MakeToken(RicottaType), keyword_length
	case "taste":
		return MakeToken(TasteKeyword), keyword_length
	case "recipe":
		return MakeToken(RecipeKeyword), keyword_length
	case "prepare":
		return MakeToken(PrepareKeyword), keyword_length
	case "curdle":
		return MakeToken(CurdleKeyword), keyword_length
	case "drain":
		return MakeToken(DrainKeyword), keyword_length
	case "spoiled":
		return MakeToken(SpoiledMilk), keyword_length
	case "fresh":
		return MakeToken(FreshMilk), keyword_length
	default:
		return MakeTokenWithMessage(Identifier, keyword), keyword_length
	}
}
