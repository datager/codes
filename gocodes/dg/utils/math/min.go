package math

func MinInt(head int, tail ...int) int {
	min := head
	for _, v := range tail {
		if v < min {
			min = v
		}
	}
	return min
}

func MinInt64(head int64, tail ...int64) int64 {
	min := head
	for _, v := range tail {
		if v < min {
			min = v
		}
	}
	return min
}
