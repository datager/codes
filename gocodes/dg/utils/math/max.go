package math

func MaxInt(head int, tail ...int) int {
	max := head
	for _, v := range tail {
		if v > max {
			max = v
		}
	}
	return max
}
