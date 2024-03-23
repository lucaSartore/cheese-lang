package tokenizer

import (
	"bytes"
	"errors"
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

	functionToTry := []func([]byte) (Token, uint){AdvanceWitheSpace, TryReadComment, TryReadDualCharacterToken, TryReadSingleCharacterToken}
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

	text := string(input[0:1])

	if text != "\\" {
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
