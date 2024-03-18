package tokenizer

import "bytes"

func Tokenize(input string) ([]Token, error) {

	inputBuffer := bytes.NewBufferString(input)

	return []Token{}, nil
}

func TokenizeSingle(input *bytes.Buffer) (Token, error) {
	// ...

	return Token{}, nil
}
