package utils

func Filter[T any](vs []T, pred func(t T, ind int) bool) []T {
	var result []T

	for i, val := range vs {
		if pred(val, i) {
			result = append(result, val)
		}
	}

	return result
}
