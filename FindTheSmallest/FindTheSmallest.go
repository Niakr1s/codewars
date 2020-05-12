package findthesmallest

import (
	"math"
)

// Smallest my implementation
func Smallest(n int64) []int64 {
	slicedN := int64ToSlice(n)

	var res0 int64 = math.MaxInt64
	res1 := math.MaxInt32
	res2 := math.MaxInt32

	for from := range slicedN {
		for to := range slicedN {
			// making new copy
			arr := make([]int, len(slicedN))
			copy(arr, slicedN)

			arr = move(arr, from, to)

			got := sliceToInt64(arr)

			if got < res0 {
				// should update all in res
				res0 = got
				res1 = from
				res2 = to
			} else if got == res0 {
				// updating only values at 1 and 2 indices
				if from < res1 {
					res1 = from
					res2 = to
				} else if from == res1 {
					// updating value at 2 indice with smallest possible
					if to < res2 {
						res2 = to
					}
				}
			}
		}
	}
	return []int64{res0, int64(res1), int64(res2)}
}

// make sure n < m
func move(arr []int, from, to int) []int {
	if from == to {
		return arr
	}

	// saving value
	fromVal := arr[from]

	if from > to {
		// moving from right to left...
		for i := from; i > to; i-- {
			arr[i] = arr[i-1]
		}
	} else {
		// moving from left to right...
		for i := from; i < to; i++ {
			arr[i] = arr[i+1]
		}
	}

	// assigning saved value
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
