package util

import "strconv"

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func ToInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
