package d2f

import (
	"strconv"
	"strings"
)

// FactString2Dec2 is solution from internet
func FactString2Dec2(str string) int {
	response := 0
	step := len(str)
	iter := 0

	for step > 0 {
		rest, _ := strconv.ParseInt(string(str[iter]), 16, 10)
		response = response*step + int(rest)
		step--
		iter++
	}

	return response
}

// Dec2FactString2 is solution from internet
func Dec2FactString2(nb int) string {
	var response string

	step := 1
	remainder := 0
	for nb > 0 {
		nb, remainder = nb/step, nb%step
		response = strconv.FormatInt(int64(remainder), 16) + response
		step++
	}

	return strings.ToUpper(response)
}

// Dec2FactString from Decimal to Factorial and Back cata
func Dec2FactString(nb int) string {
	return translateArrToString(dec2FactSlice(nb))
}

// FactString2Dec from Decimal to Factorial and Back cata
func FactString2Dec(str string) int {
	res := 0

	l := len(str)

	for i, r := range str {
		res += int(F.compute(l-i-1)) * dictRuneToI[r]
	}

	return res
}

func dec2FactSlice(nb int) []int {
	res := make([]int, 0)

	if nb == 0 {
		return res
	}

	nb64 := int64(nb)

	// now maxFactorial always > 0
	var maxFactorial int = 0

	for i := 0; ; i++ {
		if f := F.compute(i); f > nb64 {
			maxFactorial = i
			break
		}
	}

exit:
	for i := maxFactorial - 1; i >= 0; i-- {
		f := F.compute(i)
		for n := i; n >= 0; n-- {
			if nf := int64(n) * f; nb64 >= nf {
				res = append(res, n)
				nb64 -= nf
				continue exit
			}
		}
	}

	return res
}
