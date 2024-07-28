package parser

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/tokenizer"
	"fmt"
)

func contains[T comparable](arr []T, x T) bool {
	for _, n := range arr {
		if n == x {
			return true
		}
	}
	return false
}

func Map[T any, E any](input []T, function func(T) E) []E {
	v := make([]E, 0, len(input))
	for _, t := range input {
		v = append(v, function(t))
	}
	return v
}

func tokenToType(token tokenizer.TokenType) expressions.VariableType {
	switch token {
	case tokenizer.MozzarellaType:
		return expressions.Mozzarella
	case tokenizer.GorgonzolaType:
		return expressions.Gorgonzola
	case tokenizer.MilkType:
		return expressions.Milk
	case tokenizer.ParmesanType:
		return expressions.Parmesan
	}
	panic(fmt.Sprintf("Unknown token type %v", token))
}
