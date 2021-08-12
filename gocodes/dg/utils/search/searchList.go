package search

import "codes/gocodes/dg/utils/slice"

func reverse(data []interface{}, asc bool) []interface{} {
	if asc {
		for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}
	}
	return data
}

func getNextOffset(reqOffset, lenReversedData int) int {
	return reqOffset + lenReversedData
}

func GetResp(data []interface{}, reqOffset, tryLimit int, ascWithFromPage bool) (int, []interface{}) {
	nextOffset := getNextOffset(reqOffset, len(data))

	_removed := removeTryElem(data, tryLimit)
	reversed := reverse(_removed, ascWithFromPage)

	return nextOffset, reversed
}

func removeTryElem(data []interface{}, tryLimit int) []interface{} {
	if len(data) == tryLimit {
		return slice.RemoveLastElem(data)
	}
	return data
}
