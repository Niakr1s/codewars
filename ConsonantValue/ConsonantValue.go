package kata

import (
	"strings"
)

func solve(str string) int {
	return getMaxValue(str)
}

func getMaxValue(str string) int {
	splitted := splitBySonant(str)
	max := 0
	for _, str := range splitted {
		value := getStrValue(str)
		if value > max {
			max = value
		}
	}
	return max
}

func splitBySonant(str string) []string {
	res := []string{""}
	for _, r := range str {
		if !isSonant(r) {
			res[len(res)-1] = res[len(res)-1] + string(r)
			continue
		}
		if res[len(res)-1] != "" {
			res = append(res, "")
		}
	}
	if res[len(res)-1] == "" {
		res = res[:len(res)-1]
	}
	return res
}

func getStrValue(str string) int {
	res := 0
	for _, r := range str {
		res += getRuneValue(r)
	}
	return res
}

func getRuneValue(r rune) int {
	return int(r - 'a' + 1)
}

func isSonant(r rune) bool {
	const sonant = "aeiou"
	return strings.IndexRune(sonant, r) != -1
}
