package math

import (
	"strconv"
	"strings"
)

//根据输入的数字，和位的最大长度，求所有长度内可选数字中，与输入数字做"与"操作，结果为真的数字
func Int2InArray(num int64, length int) []int64 {
	rs := make([]int64, 0, 5)
	max, _ := strconv.ParseInt(strings.Repeat("1", length), 2, 64)
	for i := int64(1); i <= max; i++ {
		if i&num == num {
			rs = append(rs, i)
		}
	}
	return rs
}
