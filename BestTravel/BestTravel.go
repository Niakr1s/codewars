package besttravel

// ChooseBestSum solves Best Travel 5 kyu cata
func ChooseBestSum(maxDist, iterationsRemained int, ls []int) int {
	// return immediatly, bad imput
	if iterationsRemained == 0 {
		return -1
	}

	// return immediatly, bad input
	if len(ls) < iterationsRemained {
		return -1
	}

	best := -1
	for i, d := range ls {
		if iterationsRemained > 1 {
			innerbest := ChooseBestSum(maxDist-d, iterationsRemained-1, ls[i+1:])
			if innerbest < 0 {
				continue
			}
			d += innerbest
		}
		if d > best && d <= maxDist {
			best = d
		}
	}
	return best
}

func bestNotGreaterThanN(n int, arr []int) int {
	best := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > best && arr[i] <= n {
			best = arr[i]
		}
	}
	return best
}
