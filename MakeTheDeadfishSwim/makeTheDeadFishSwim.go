package kata

func Parse(data string) []int {
	res := []int{}

	value := 0
	innerFunc := func(op rune) int {
		switch op {
		case 'i':
			value++
		case 'd':
			value--
		case 's':
			value *= value
		case 'o':
			res = append(res, value)
		}
		return value
	}

	for _, op := range data {
		innerFunc(op)
	}

	return res
}
