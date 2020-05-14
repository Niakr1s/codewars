package d2f

import "testing"

func TestDec2FactString(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{2982, "4041000"},
		{36288000, "A0000000000"},
	}

	for _, testCase := range testCases {
		got := Dec2FactString(testCase.input)
		if got != testCase.expected {
			t.Errorf("Input: %v, expected: %v, got %v",
				testCase.input, testCase.expected, got)
		}
	}
}

func TestFactString2Dec(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"341010", 463},
		{"4042100", 2990},
	}

	for _, testCase := range testCases {
		got := FactString2Dec(testCase.input)
		if got != testCase.expected {
			t.Errorf("Input: %v, expected: %v, got %v",
				testCase.input, testCase.expected, got)
		}
	}
}
