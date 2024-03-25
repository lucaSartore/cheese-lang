package tokenizer

import (
	"bytes"
	"errors"
	"strings"
)

func Tokenize(input string) ([]Token, error) {
	// todo: check tat the string is made of only asci characters... utf8 characters may brake this tokenizer
	inputBuffer := bytes.NewBuffer([]byte(input))
	tokens := []Token{}
	for {
		token, err := TokenizeSingle(inputBuffer)

		if token.TokenType != NullToken {
			tokens = append(tokens, token)
		}

		if inputBuffer.Len() == 0 {
			break
		}

		if err != nil {
			return tokens, err
		}

	}
	return tokens, nil
}

func TokenizeSingle(input *bytes.Buffer) (Token, error) {
	if input.Len() == 0 {
		return MakeToken(NullToken), nil
	}
	inputBytes := input.Bytes()

	functionToTry := []func([]byte) (Token, uint){AdvanceWitheSpace, TryReadComment, TryReadDualCharacterToken, TryReadSingleCharacterToken, TryReadKeyword, TryReadMozzarellaLiteral}
	for _, function := range functionToTry {
		token, len := function(inputBytes)
		if len > 0 {
			buffer := make([]byte, len)
			input.Read(buffer)
			return token, nil
		}
	}
	return MakeToken(NullToken), errors.New("unable to find token that match with: " + string(inputBytes))
}

func AdvanceWitheSpace(input []byte) (Token, uint) {
	count := uint(0)
	for _, b := range input {
		if b == ' ' || b == '\t' || b == '\n' {
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

	if text != "\\\\" {
		return MakeToken(NullToken), 0
	}
	count := uint(2)

	for _, c := range input[2:] {
		count++
		if c == '\n' {
			break
		}
	}
	return MakeTokenWithMessage(Comment, string(input[2:count-1])), count
}

func TryReadKeyword(input []byte) (Token, uint) {
	str := string(input)

	type KeyWordOption struct {
		keyword string
		token   TokenType
	}

	options := []KeyWordOption{
		{"mozzarella", MozzarellaType},
		{"parmesan", ParmesanType},
		{"gorgonzola", GorgonzolaType},
		{"ricotta", RicottaType},
		{"taste", TasteKeyword},
		{"recipe", RecipeKeyword},
		{"prepare", PrepareKeyword},
		{"curdle", CurdleKeyword},
		{"drain", DrainKeyword},
	}

	for _, option := range options {
		if strings.HasPrefix(str, option.keyword) {
			return MakeToken(option.token), uint(len(option.keyword))
		}
	}
	return MakeToken(NullToken), 0
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
			return MakeToken(NullToken), 0
		}
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
