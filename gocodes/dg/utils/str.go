package utils

import (
	"encoding/binary"
	"strconv"
)

func SubstrLength(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func Float64ToBytes(f float64) []byte {
	str := strconv.FormatFloat(f, 'E', -1, 64) //float64
	return []byte(str)
}

func Float64ToStr(f float64) string {
	return string(Float64ToBytes(f))
}
