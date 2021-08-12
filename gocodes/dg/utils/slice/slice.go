package slice

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

// HasElem only suport slice type
func HasElem(s interface{}, elem interface{}) bool {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

func IsInSlice(v string, slice []string) bool {
	for _, vv := range slice {
		if vv == v {
			return true
		}
	}
	return false
}

func IsInIntSlice(v int64, slice []int64) bool {
	for _, vv := range slice {
		if vv == v {
			return true
		}
	}
	return false
}

func GetIntersectionStrings(slice1, slice2 []string) (diffSlice []string) {
	diffSlice = make([]string, 0)
	for _, v := range slice1 {
		if IsInSlice(v, slice2) {
			diffSlice = append(diffSlice, v)
		}
	}
	return
}

func GetUnionStrings(slice1, slice2 []string) (unionSlice []string) {
	slice := make([]string, 0)
	slice = append(slice, slice1...)
	slice = append(slice, slice2...)
	return RemoveDuplicateElement(slice)
}

func GetIntersectRemoveDupElems(slice1, slice2 []string) []string {
	return GetIntersectionStrings(RemoveDuplicateElement(slice1), RemoveDuplicateElement(slice2))
}

func RemoveDuplicateElement(slice []string) []string {
	result := make([]string, 0, len(slice))
	temp := map[string]struct{}{}
	for _, item := range slice {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//--------------------------------------------------------
// item 是 slice 的子集
func IsItemInSlice(item string, slice []string) bool {
	for _, s := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func IsIntItemInSlice(item int, slice []int) bool {
	for _, s := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// s1 是 s2 的子集
func IsSubSet(s1, s2 []string) bool {
	if len(s2) == 0 {
		return false
	}

	for _, s := range s1 {
		if !IsItemInSlice(s, s2) {
			return false
		}
	}
	return true
}

// s1 与 s2 相同(不考虑顺序)
// a: 12345
// b: 34251
// IsEqualWithoutOrder(a, b) = true
func IsEqualWithoutOrder(s1, s2 []string) bool {
	return IsSubSet(s1, s2) && IsSubSet(s2, s1)
}

//--------------------------------------------------------
// a 相对 b 的差集: 即属于 a, 但是不属于 b 的
// a: 12345
// b: 34567
// differentSet(a, b): 12
// differentSet(b, a): 67
// intersection(a, b): 345
// union(a, b): 1234567
// isSubset(a, b): false
// isSubset(b, a): false
// removeDuplicated(11223445): 12345
func GetDiffStrSet(sa, sb []string) (differentSetAThanB []string) {
	if len(sb) == 0 {
		return sa
	}

	differentSetAThanB = make([]string, 0)
	for _, a := range sa {
		if !IsItemInSlice(a, sb) {
			differentSetAThanB = append(differentSetAThanB, a)
		}
	}
	return
}

func GetDiffIntSet(sa, sb []int64) (differentSetAThanB []int64) {
	if len(sb) == 0 {
		return sa
	}

	differentSetAThanB = make([]int64, 0)
	for _, a := range sa {
		if !HasElem(sb, a) {
			differentSetAThanB = append(differentSetAThanB, a)
		}
	}
	return
}

// old = [1,2,3,4,5,6,7,8,9,10], batchSize = 3
// new = [[1,2,3], [4,5,6], [7,8,9], [10]]
func SplitSlices(old []interface{}, batchSize int) (newBatches [][]interface{}, err error) {
	if len(old) == 0 {
		return nil, errors.New("old slice is nil")
	}

	batchNum := len(old) / batchSize
	//remainedNum := len(old) % batchSize

	newBatches = make([][]interface{}, 0)
	for i := 0; i < batchNum; i++ {
		newOneBatch := old[(i * batchSize):((i + 1) * batchSize)]
		newBatches = append(newBatches, newOneBatch)
	}

	remainedBatch := old[(batchNum * batchSize):]
	newBatches = append(newBatches, remainedBatch)

	return
}

func SplitSlicesStr(old []string, batchSize int) (newBatches [][]string, err error) {
	oldEface := make([]interface{}, 0)
	for _, o := range old {
		oldEface = append(oldEface, o)
	}

	newBatchesEface, err := SplitSlices(oldEface, batchSize)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	newBatches = make([][]string, 0)
	for _, newOneBatchEface := range newBatchesEface {
		newOneBatch := reflect.ValueOf(newOneBatchEface).Interface().([]string)
		newBatches = append(newBatches, newOneBatch)
	}

	return
}

func ParseIDsByRemoveComma(before string) []string {
	return strings.Split(before, ",")
}

func FormatIDsByAddQuote(before []string) []string {
	ret := make([]string, 0)
	for _, b := range before {
		ret = append(ret, fmt.Sprintf("'%v'", b))
	}
	return ret
}

func FormatIDsByAddComma(before []string) string {
	return strings.Join(before, ",")
}

func RemoveLastElem(before []interface{}) []interface{} {
	if len(before) == 0 {
		return before
	}
	return before[:len(before)-1]
}
