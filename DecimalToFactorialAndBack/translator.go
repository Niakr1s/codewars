package d2f

var dictIToRune map[int]rune
var dictRuneToI map[rune]int

func init() {
	initDicts()
}

func translateArrToString(arr []int) string {
	res := ""

	for _, v := range arr {
		res += string(dictIToRune[v])
	}

	return res
}

func initDicts() {
	dictIToRune = make(map[int]rune)
	dictRuneToI = make(map[rune]int)

	i := 0

	for r := '0'; r <= '9'; r++ {
		dictIToRune[i] = r
		dictRuneToI[r] = i
		i++
	}

	for r := 'A'; r <= 'Z'; r++ {
		dictIToRune[i] = r
		dictRuneToI[r] = i
		i++
	}
}
