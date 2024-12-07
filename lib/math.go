package lib

import "golang.org/x/exp/constraints"

func Min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func Sum(values []int) int {
	var acc int
	for _, v := range values {
		acc += v
	}
	return acc
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
