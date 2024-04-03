package main

import (
	"fmt"

	"cheese-lang/internal/tokenizer"
)

func main() {
	s := "\\\\this is a comment\n == = != && ^{    }\t	\n\n \t  -="
	t, err := tokenizer.Tokenize(s)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(t)
}
