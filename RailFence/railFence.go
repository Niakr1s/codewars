package railfence

import "strings"

func Encode(s string, n int) string {
	if s == "" {
		return ""
	}
	encodedSlice := encodeToSlice(s, n)
	res := strings.Join(encodedSlice, "")
	return res
}

func encodeToSlice(s string, n int) []string {
	res := make([]string, n)
	i := 0
	diff := 1
	for _, r := range s {
		res[i] += string(r)
		i += diff
		if i == n-1 || i == 0 {
			diff *= -1
		}
	}
	return res
}

func decodeToSlice(s string, n int) []string {
	decodedSlice := encodeToSlice(s, n)
	counter := 0
	for i, line := range decodedSlice {
		decodedSlice[i] = s[counter : counter+len(line)]
		counter += len(line)
	}
	return decodedSlice
}

func Decode(s string, n int) string {
	if s == "" {
		return ""
	}
	decodedSlice := decodeToSlice(s, n)
	indexes := make([]int, len(decodedSlice))
	i := 0
	diff := 1
	res := ""
	for {
		if indexes[i] >= len(decodedSlice[i]) {
			return res
		}
		res += decodedSlice[i][indexes[i] : indexes[i]+1]
		indexes[i]++
		i += diff
		if i == n-1 || i == 0 {
			diff *= -1
		}
	}
}
