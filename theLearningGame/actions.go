package learninggameone

func Actions() []func(int) int {
	return []func(int) int{
		func(x int) int { return x + 1 },
		func(x int) int { return 0 },
		func(x int) int { return int(x / 2) },
		func(x int) int { return x * 100 },
		func(x int) int { return x % 2 }}
}
