package kata

import (
	"strings"
)

func toWeirdCase(str string) string {
	res := []string{}
	for _, subStr := range strings.Split(str, " ") {
		weirdSubStr := ""
		for i, r := range subStr {
			if i%2 == 1 {
				weirdSubStr += strings.ToLower(string(r))
			} else {
				weirdSubStr += strings.ToUpper(string(r))
			}
		}
		res = append(res, weirdSubStr)
	}
	return strings.Join(res, " ")
}
