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

func PowInt(base, exp int) int {
	result := 1
	for {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}

	return result
}
