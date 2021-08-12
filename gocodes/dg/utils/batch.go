package utils

func SplitSliceByOnceNum(slice []string, num int) (result [][]string) {
	var once []string
	for index, elem := range slice {
		once = append(once, elem)
		if index >= num && index%num == 0 {
			result = append(result, once)
			once = []string{}
		}
	}
	result = append(result, once)
	return
}
