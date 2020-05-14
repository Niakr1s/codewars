package d2f

// F is global cached factorial
var F CachedFactorial

func init() {
	F = NewCachedFactorial()
}

// CachedFactorial computes factorial and holds cache for faster execution
type CachedFactorial struct {
	cache map[int]int64
}

// NewCachedFactorial constructs Factorial
func NewCachedFactorial() CachedFactorial {
	return CachedFactorial{cache: make(map[int]int64)}
}

func (f *CachedFactorial) compute(n int) int64 {
	if res, ok := f.cache[n]; ok {
		return res
	}

	if n == 0 {
		return 0
	}

	var res int64 = 1

	for i := 1; i <= n; i++ {
		if v, ok := f.cache[i]; ok {
			res = v
			continue
		}
		res *= int64(i)
		f.cache[i] = res
	}

	return res
}
