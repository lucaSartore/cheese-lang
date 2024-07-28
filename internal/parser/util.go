package parser

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
