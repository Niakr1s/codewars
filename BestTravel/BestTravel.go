package besttravel

// ChooseBestSum solves Best Travel 5 kyu cata
func ChooseBestSum(maxDist, maxTowns int, ls []int) int {
	permutations := make([][]int, 0)
	getPermutations(maxTowns, ls, &permutations, []int{})

	best := 0
	for i := 0; i != len(permutations); i++ {
		if len(permutations[i]) < maxTowns {
			continue
		}
		if sum := sumArr(permutations[i]); sum > best && sum <= maxDist {
			best = sum
		}
	}

	if best == 0 {
		return -1
	}

	return best
}

func sumArr(arr []int) int {
	sum := 0
	for i := 0; i != len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func getPermutations(iterationsRemained int, arr []int, accum *[][]int, path []int) {
	// shouldn't never happen
	if iterationsRemained <= 0 {
		return
	}

	// just appending array into accum
	if iterationsRemained == 1 {
		for i := 0; i != len(arr); i++ {
			pathCopy := append(path[:0:0], path...)
			*accum = append(*accum, append(pathCopy, arr[i]))
		}
	}

	// iterationsRemained > 1
	if len(arr) == 0 {
		return
	}

	// iterationsRemained > 1
	// len(arr) > 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			pathCopy := append(path[:0:0], path...)
			getPermutations(iterationsRemained-1, arr[j:], accum, append(pathCopy, arr[i]))
		}
	}
}
