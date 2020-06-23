package validator

func ValidateSolution(m [][]int) bool {
	if !validateRows(m) {
		return false
	}
	if !validateColumns(m) {
		return false
	}
	if !validateBlocks(m) {
		return false
	}
	return true
}

func validateRows(m [][]int) bool {
	for _, row := range m {
		if !validate(row) {
			return false
		}
	}
	return true
}

func validateColumns(m [][]int) bool {
	for colIdx := range m[0] {
		col := make([]int, 0, len(m))
		for rowIdx := range m {
			col = append(col, m[rowIdx][colIdx])
		}
		if !validate(col) {
			return false
		}
	}
	return true
}

func validateBlocks(m [][]int) bool {
	for blockCol := 0; blockCol < 3; blockCol++ {
		for blockRow := 0; blockRow < 3; blockRow++ {
			block := make([]int, 0, 9)
			for col := blockCol * 3; col < blockCol*3+3; col++ {
				for row := blockRow * 3; row < blockRow*3+3; row++ {
					block = append(block, m[row][col])
				}
			}
			if !validate(block) {
				return false
			}
		}
	}
	return true
}

func validate(m []int) bool {
	visited := map[int]struct{}{}
	for _, n := range m {
		if n < 1 || n > 9 {
			return false
		}
		if _, ok := visited[n]; ok {
			return false
		}
		visited[n] = struct{}{}
	}
	return true
}
