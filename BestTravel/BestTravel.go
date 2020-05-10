package besttravel

// ChooseBestSum solves Best Travel 5 kyu cata
func ChooseBestSum(maxDist, iterationsRemained int, ls []int) int {
	// return immediatly
	if len(ls) < iterationsRemained {
		return -1
	}

	// after that we will be are assured that len(ls) > 0
	if iterationsRemained <= 0 {
		return -1
	}

	if iterationsRemained == 1 {
		return bestNotGreaterThanN(maxDist, ls)
	}

	// all "good" combinations (that are <= maxDist) will be stored here
	filteredResults := []int{}

	for i := 0; i != len(ls); i++ {
		for j := i + 1; j != len(ls); j++ {
			best := ChooseBestSum(maxDist-ls[i], iterationsRemained-1, ls[j:])
			if best != -1 {
				filteredResults = append(filteredResults, best+ls[i])
			}
		}
	}

	return bestNotGreaterThanN(maxDist, filteredResults)
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
