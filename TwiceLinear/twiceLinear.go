package twicelinear

import (
	"sort"
)

func DblLinear(n int) int {
	return dblLinearInner(n, newArr([]int{1}, map[int]bool{1: true}))
}

func dblLinearInner(n int, a ArrWIthCache) int {
	if n == 0 {
		return a.arr[0]
	}

	y0 := y(a.arr[0])
	z0 := z(a.arr[0])

	// if y0 if in out of array - we got result!
	if len(a.arr) > n+1 && y0 > a.arr[n] {
		return a.arr[n]
	}

	a.AppendInSorted(y0)
	a.AppendInSorted(z0)

	return dblLinearInner(n-1, newArr(a.arr[1:], a.appended))
}

func y(x int) int {
	return 2*x + 1
}

func z(x int) int {
	return 3*x + 1
}

type Arr []int

type ArrWIthCache struct {
	arr      sort.IntSlice
	appended map[int]bool
}

func newArr(a []int, app map[int]bool) ArrWIthCache {
	return ArrWIthCache{
		arr:      a,
		appended: app,
	}
}

func (a *ArrWIthCache) AppendInSorted(n int) bool {
	if _, ok := a.appended[n]; ok {
		return false
	}
	a.appended[n] = true

	idx := sort.SearchInts(a.arr, n)

	a.arr = append(a.arr, 0)
	copy(a.arr[idx+1:], a.arr[idx:])

	a.arr[idx] = n

	return true
}
