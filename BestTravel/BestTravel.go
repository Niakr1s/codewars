package besttravel

// ChooseBestSum solves Best Travel 5 kyu cata
func ChooseBestSum(maxDist, townsLeft int, ls []int) int {
	return chooseBestSumWithAccum(maxDist, townsLeft, ls, 0)
}

func chooseBestSumWithAccum(maxDist, townsLeft int, ls []int, accum int) int {
	if townsLeft <= 0 {
		return accum
	}
	if accum > maxDist {
		return -1
	}

	bestSum := 0

	for i := 0; i != len(ls); i++ {
		lsWithTownRemoved := append(ls[:0:0], ls...)
		lsWithTownRemoved = deleteItem(i, lsWithTownRemoved[:])

		best := chooseBestSumWithAccum(maxDist, townsLeft-1, lsWithTownRemoved, accum+ls[i])

		if best > maxDist {
			continue
		}
		if best > bestSum {
			bestSum = best
		}
	}

	if bestSum == 0 {
		bestSum = -1
	}

	return bestSum
}

func deleteItem(n int, arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	if n < 0 || n >= len(arr) {
		return arr
	}
	if len(arr) <= 1 {
		return arr[:0]
	}
	if n == 0 {
		return arr[1:]
	} else if n == len(arr)-1 {
		return arr[:len(arr)-1]
	} else {
		return append(arr[:n], arr[n+1:]...)
	}
}
