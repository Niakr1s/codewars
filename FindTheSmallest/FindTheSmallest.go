package findthesmallest

import (
	"math"
	"strconv"
)

// Smallest my implementation
func Smallest(n int64) []int64 {
	slicedN := int64ToSlice(n)

	var resValue int64 = math.MaxInt64
	resFrom := math.MaxInt32
	resTo := math.MaxInt32

	for from := range slicedN {
		for to := range slicedN {
			// making new copy
			arr := make([]int, len(slicedN))
			copy(arr, slicedN)

			arr = move(arr, from, to)

			got := sliceToInt64(arr)

			if got < resValue {
				resValue = got
				resFrom = from
				resTo = to
			}
			// if got == resValue is unneeded
			// because we will get more bad result
		}
	}
	return []int64{resValue, int64(resFrom), int64(resTo)}
}

// Smallest2 if from internet
// My variant is 2 times faster than this xD
func Smallest2(n int64) []int64 {
	tmp := n
	s := strconv.FormatInt(n, 10)
	var res = []int64{-1, 0, 0}
	for i, _ := range s {
		c := string(s[i])
		str1 := s[:i] + s[i+1:]
		for j, _ := range s {
			str2 := str1[:j] + c + str1[j:]
			nbStr2, _ := strconv.ParseInt(str2, 10, 64)
			if nbStr2 < tmp {
				tmp = nbStr2
				res = []int64{nbStr2, int64(i), int64(j)}
			}
		}
	}
	if res[0] == -1 {
		res = []int64{n, 0, 0}
	}
	return res
}

// make sure n < m
func move(arr []int, from, to int) []int {
	if from == to {
		return arr
	}

	fromVal := arr[from]

	// ejecting from
	copy(arr[from:], arr[from+1:])

	// making space for to
	copy(arr[to+1:], arr[to:])

	arr[to] = fromVal

	return arr
}

func sliceToInt64(slice []int) int64 {
	var multiplier int64 = 1
	var result int64 = 0

	for i := len(slice) - 1; i >= 0; i-- {
		result += int64(slice[i]) * multiplier
		multiplier *= 10
	}

	return result
}

func int64ToSlice(n int64) []int {
	arr := []int{}

	for n != 0 {
		arr = append(arr, int(n%10))
		n = n / 10
	}

	for i := 0; i != len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	return arr
}
