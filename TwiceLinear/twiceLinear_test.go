package twicelinear

import (
	"testing"
)

func TestDblLinear(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{10, 22},
		{50, 175},
		{100, 447},
		{500, 3355},
		{1000, 8488},
		{100000, 289579},
	}

	for _, v := range testCases {
		got := DblLinear(v.input)
		if got != v.expected {
			t.Errorf("Expected: %v, got: %v", v.expected, got)
		}
	}
}
