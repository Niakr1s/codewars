package findthesmallest

import (
	"reflect"
	"testing"
)

func TestSmallest(t *testing.T) {
	testCases := []struct {
		input    int64
		expected []int64
	}{
		{261235, []int64{126235, 2, 0}},
		{209917, []int64{29917, 0, 1}},
		{199819884756, []int64{119989884756, 4, 0}},
		{5190930055291, []int64{519093055291, 6, 0}},
		{60496893958362, []int64{4696893958362, 0, 2}},
	}

	for _, c := range testCases {
		got := Smallest(c.input)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("TestSmallest: input: %v, expected: %v, got: %v", c.input, c.expected, got)
		}
	}
}

func TestExchange(t *testing.T) {
	testCases := []struct {
		arr      []int
		from     int
		to       int
		expected []int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 3, 1, []int{0, 3, 1, 2, 4, 5}},
		{[]int{0, 1, 2, 3, 4, 5}, 1, 3, []int{0, 2, 3, 1, 4, 5}},
	}

	for _, testCase := range testCases {
		got := move(testCase.arr, testCase.from, testCase.to)
		if !reflect.DeepEqual(got, testCase.expected) {
			t.Errorf("TestExchange: got %v (%d => %d), expected: %v, got: %v",
				testCase.arr, testCase.from, testCase.to,
				testCase.expected, got)
		}
	}
}

func TestSliceToInt64(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int64
	}{
		{[]int{1, 2, 1, 1, 1}, 12111},
		{[]int{1, 2, 3, 4, 5, 6}, 123456},
		{[]int{6, 5, 4, 3, 2, 1}, 654321},
		{[]int{6, 5, 4, 7, 5, 4, 8}, 6547548},
	}

	for _, testCase := range testCases {
		got := sliceToInt64(testCase.input)
		if got != testCase.expected {
			t.Errorf("TestSliceToInt64: expected: %v, got: %v", testCase.expected, got)
		}
	}
}

func BenchmarkSmallest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Smallest(5275027536092356050)
	}
}

func BenchmarkSmallest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Smallest2(5275027536092356050)
	}
}
