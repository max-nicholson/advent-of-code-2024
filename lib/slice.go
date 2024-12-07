package lib

func Pairs[T any](elements []T) [][]T {
	totalElements := len(elements)
	totalPairs := (totalElements * (totalElements - 1)) / 2
	pairs := make([][]T, totalPairs)
	var pair int
	for i, first := range elements {
		for j := i + 1; j < totalElements; j++ {
			pairs[pair] = []T{first, elements[j]}
			pair += 1
		}
	}
	return pairs
}
