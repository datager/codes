package math

import "math"

func PowInt64(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}