package snail

import "fmt"

func Snail(snailMap [][]int) []int {
	fmt.Print(snailMap)
	res := []int{}
	for dimension(snailMap) > 0 {
		res = append(res, getOuterPath(snailMap)...)
		snailMap = strip(snailMap)
	}
	return res
}

func getOuterPath(snailMap [][]int) []int {
	res := []int{}
	dim := dimension(snailMap)
	switch dim {
	case 0:

	case 1:
		res = append(res, snailMap[0][0])

	default:
		// right
		for col := 0; col < dim-1; col++ {
			res = append(res, snailMap[0][col])
		}
		// down
		for row := 0; row < dim-1; row++ {
			res = append(res, snailMap[row][dim-1])
		}
		// left
		for col := dim - 1; col > 0; col-- {
			res = append(res, snailMap[dim-1][col])
		}
		// up
		for row := dim - 1; row > 0; row-- {
			res = append(res, snailMap[row][0])
		}
	}
	return res
}

func strip(snailMap [][]int) [][]int {
	res := [][]int{}
	dim := dimension(snailMap)
	switch dim {
	case 0, 1, 2:

	default:
		for row := 1; row < dim-1; row++ {
			toAppend := []int{}
			for col := 1; col < dim-1; col++ {
				toAppend = append(toAppend, snailMap[row][col])
			}
			res = append(res, toAppend)
		}
	}
	return res
}

func dimension(snailMap [][]int) int {
	res := len(snailMap)
	if res == 1 {
		res = len(snailMap[0])
	}
	return res
}
