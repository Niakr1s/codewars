package squareintosquares

import "sort"

func Decompose(n int64) []int64 {
	return sortAscending(decomposeInner(n*n, n-1))
}

func decomposeInner(n int64, startFrom int64) []int64 {
	if startFrom == 0 || n == 0 {
		return []int64{}
	}
	for i := startFrom; i > 0; i-- {
		res := []int64{i}
		rem := n - i*i
		if rem < 0 {
			continue
		}
		if rem == 0 {
			return res
		}
		res = append(res, decomposeInner(rem, i-1)...)
		if checkResult(n, res) {
			return res
		}
	}
	return []int64{}
}

func checkResult(expected int64, decomposed []int64) bool {
	var sumOfSquares int64 = 0
	for _, e := range decomposed {
		sumOfSquares += e * e
	}
	return sumOfSquares == expected
}

func sortAscending(slice []int64) []int64 {
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	return slice
}
