package rangeextraction

import (
	"strconv"
	"strings"
)

func Solution(list []int) string {
	res := []string{}
	i := 0
	for i < len(list) {
		lastIdx, l := findLastRangeIdxAndLen(list, i)
		if l <= 2 {
			res = append(res, strconv.Itoa(list[i]))
			i++
			continue
		}
		res = append(res, generateRangeString(list[i], list[lastIdx]))
		i += l
	}
	return strings.Join(res, ",")
}

func generateRangeString(start, end int) string {
	return strconv.Itoa(start) + "-" + strconv.Itoa(end)
}

// findLastRangeIdxAndLen returns index of last element in range and length of this range
func findLastRangeIdxAndLen(list []int, idx int) (int, int) {
	i := idx
	expected := list[idx]
	for ; i < len(list); i++ {
		if list[i] != expected {
			break
		}
		expected++
	}
	return i - 1, i - idx
}
