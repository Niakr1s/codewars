package d2f

// CachedFactorial computes factorial and holds cache for faster execution
type CachedFactorial struct{
	cache map[int]int
}

// NewCachedFactorial constructs Factorial
func NewCachedFactorial() CachedFactorial{
	return CachedFactorial{cache: make(map[int]int)}
}

func (f *CachedFactorial) compute(n int) int {
	if res, ok := f.cache[n]; ok {
		return res
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	res := n * f.compute(n-1)
	f.cache[n] = res

	return res
}