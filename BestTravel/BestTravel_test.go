package besttravel

import (
	"testing"
)

func TestChooseBestSum(t *testing.T) {
	dotest := func(tt, k int, ls []int, exp int) {
		var ans = ChooseBestSum(tt, k, ls)

		if ans != exp {
			t.Errorf("(%v, %v, %v): Expected: %v, got %v", tt, k, ls, exp, ans)
		}
	}

	var ts = []int{0, 1, 2, 3}
	dotest(20, 0, ts, -1)
	dotest(20, 1, ts, 3)
	dotest(20, 2, ts, 5)
	dotest(20, 3, ts, 6)
	ts = []int{50, 55, 56, 57, 58}
	dotest(163, 3, ts, 163)
	ts = []int{50}
	dotest(163, 3, ts, -1)
	ts = []int{91, 74, 73, 85, 73, 81, 87}
	dotest(230, 3, ts, 228)
	dotest(331, 2, ts, 178)
	dotest(331, 4, ts, 331)
	dotest(331, 5, ts, -1)
	ts = []int{100, 76, 56, 44, 89, 73, 68, 56, 64, 123, 2333, 144, 50, 132, 123, 34, 89}
	dotest(880, 8, ts, 876)
}
