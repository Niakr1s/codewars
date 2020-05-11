package twicelinear

import (
	"sort"
)

// DblLinear is my Twice linear solution
func DblLinear(n int) int {
	return dblLinearInner(n, []int{1})
}

func dblLinearInner(n int, arr sort.IntSlice) int {
	if n == 0 {
		return arr[0]
	}

	y0 := y(arr[0])
	z0 := z(arr[0])

	// if y0 if in out of array - we got result!
	if len(arr) > n+1 && y0 > arr[n] {
		return arr[n]
	}

	arr = appendIntoSorted(arr, y0)
	arr = appendIntoSorted(arr, z0)

	return dblLinearInner(n-1, arr[1:])
}

func y(x int) int {
	return 2*x + 1
}

func z(x int) int {
	return 3*x + 1
}

func appendIntoSorted(arr sort.IntSlice, n int) []int {
	idx := sort.SearchInts(arr, n)

	// if n is already in arr
	if idx != len(arr) && arr[idx] == n {
		return arr
	}

	arr = append(arr, 0) // make sure arr doesn't overflow while copying
	copy(arr[idx+1:], arr[idx:])

	arr[idx] = n

	return arr
}
