package kata

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		str  string
		want int
	}{
		{"a", 0},
		{"bcd", 9},
		{"zea", 26},
		{"az", 26},
		{"baz", 26},
		{"aeiou", 0},
		{"uaoczei", 29},
		{"abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes", 143},
		{"codewars", 37},
		{"bup", 16},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := solve(tt.str); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitBySonant(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []string
	}{
		{"empty", "", []string{}},
		{"sonant only", "aeiou", []string{}},
		{"consonant only", "bcdfgh", []string{"bcdfgh"}},
		{"mixed starting by sonant ended by consonant", "abcdfgh", []string{"bcdfgh"}},
		{"mixed starting by sonant ended by sonant", "abcdafgha", []string{"bcd", "fgh"}},
		{"mixed starting by consonant ended by consonant", "bcdafgh", []string{"bcd", "fgh"}},
		{"mixed starting by consonant ended by sonant", "bcdafgha", []string{"bcd", "fgh"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitBySonant(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitBySonant() = %v, want %v", got, tt.want)
			}
		})
	}
}
