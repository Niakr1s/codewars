package d2f

import "testing"

func TestFactorial(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	f := NewCachedFactorial()

	for _, testCase := range testCases {
		got := f.compute(testCase.input)
		if got != testCase.expected {
			t.Errorf("Input: %d, expected: %d, got %d",
				testCase.input, testCase.expected, got)
		}
	}
}
