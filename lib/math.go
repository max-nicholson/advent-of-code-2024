package lib

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
