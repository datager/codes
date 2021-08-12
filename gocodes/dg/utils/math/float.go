package math

import "math"

// uint 代表小数位数，格式位 0.000001 如果是几位就指定为几位
func TruncateNaive(f float64, unit float64) float64 {
	return math.Trunc(f/unit) * unit
}
