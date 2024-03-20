package tokenizer

import "bytes"

func Tokenize(input string) ([]Token, error) {
	// todo: check tat the string is made of only asci characters... utf8 characters may brake this tokenizer
	inputBuffer := bytes.NewBufferString(input)

	return []Token{}, nil
}

func TokenizeSingle(input *bytes.Buffer) (Token, error) {
	inputBytes := input.Bytes()
	return Token{}, nil
}

func AdvanceWitheSpace(input []byte) uint {
	count := uint(0)
	for b := range input {
		if b == ' ' || b == '\t' || b == '\n' {
			count++
		} else {
			break
		}
	}
	return count
}

// try to match with operators that are made of 2 characters
func TryReadDualCharacterOperator(input []byte) (Token, uint) {
	if len(input) < 2 {
		return MakeToken(NullToken), 0
	}
	text := string(input[0:1])
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
	default:
		return MakeToken(NullToken), 0
	}
}

func TryReadSingleCHaracterOperator(input []byte) (Token, uint) {
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
	default:
		return MakeToken(NotOperator), 1
	}
}
