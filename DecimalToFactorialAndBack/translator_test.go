package d2f

import "testing"

func TestNewDict(t *testing.T) {
	testCases := []struct {
		i        int
		expected rune
	}{
		{0, '0'},
		{9, '9'},
		{10, 'A'},
		{35, 'Z'},
	}

	for _, testCase := range testCases {
		got := dictIToRune[testCase.i]
		if got != testCase.expected {
			t.Errorf("i=%d, expected: %v, got %v", testCase.i, testCase.expected, got)
		}
	}
}

func TestTranslate(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected string
	}{
		{[]int{}, ""},
		{[]int{1, 5, 10}, "15A"},
		{[]int{9, 35}, "9Z"},
	}

	for _, testCase := range testCases {
		got := translateArrToString(testCase.arr)
		if got != testCase.expected {
			t.Errorf("input: %v, expected: %v, got %v",
				testCase.arr, testCase.expected, got)
		}
	}
}
