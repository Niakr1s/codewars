package d2f

import (
	"reflect"
	"testing"
)

func BenchmarkFactString2Dec(b *testing.B) {
	for i := 0; i != b.N; i++ {
		FactString2Dec("A0000000000")
	}
}

func BenchmarkDec2FactString(b *testing.B) {
	for i := 0; i != b.N; i++ {
		Dec2FactString(36288000)
	}
}

func BenchmarkFactString2Dec2(b *testing.B) {
	for i := 0; i != b.N; i++ {
		FactString2Dec2("A0000000000")
	}
}

func BenchmarkDec2FactString2(b *testing.B) {
	for i := 0; i != b.N; i++ {
		Dec2FactString2(36288000)
	}
}

func TestDec2FactString(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{463, "341010"},
		{2982, "4041000"},
		{2990, "4042100"},
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
		{"4041000", 2982},
		{"4042100", 2990},
		{"A0000000000", 36288000},
	}

	for _, testCase := range testCases {
		got := FactString2Dec(testCase.input)
		if got != testCase.expected {
			t.Errorf("Input: %v, expected: %v, got %v",
				testCase.input, testCase.expected, got)
		}
	}
}

func TestDec2FactSlice(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
	}{
		{463, []int{3, 4, 1, 0, 1, 0}},
		{2982, []int{4, 0, 4, 1, 0, 0, 0}},
		{2990, []int{4, 0, 4, 2, 1, 0, 0}},
		{36288000, []int{10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}

	for _, testCase := range testCases {
		got := dec2FactSlice(testCase.input)
		if !reflect.DeepEqual(got, testCase.expected) {
			t.Errorf("Input: %v, expected: %v, got %v",
				testCase.input, testCase.expected, got)
		}
	}
}
