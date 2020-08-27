package snail

import (
	"reflect"
	"testing"
)

func TestSnail(t *testing.T) {
	tests := []struct {
		name     string
		snailMap [][]int
		want     []int
	}{
		{"0x0", [][]int{{}}, []int{}},
		{"3x3 #1", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"3x3 #2", [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"4x4", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Snail(tt.snailMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Snail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getOuterPath(t *testing.T) {
	tests := []struct {
		name     string
		snailMap [][]int
		want     []int
	}{
		{"0x0", [][]int{}, []int{}},
		{"1x1", [][]int{{1}}, []int{1}},
		{"2x2", [][]int{{1, 2}, {3, 4}}, []int{1, 2, 4, 3}},
		{"3x3", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOuterPath(tt.snailMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOuterPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strip(t *testing.T) {
	tests := []struct {
		name     string
		snailMap [][]int
		want     [][]int
	}{
		{"0x0", [][]int{}, [][]int{}},
		{"1x1", [][]int{{1}}, [][]int{}},
		{"2x2", [][]int{{1, 2}, {3, 4}}, [][]int{}},
		{"3x3", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{5}}},
		{"4x4", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, [][]int{{6, 7}, {10, 11}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strip(tt.snailMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strip() = %v, want %v", got, tt.want)
			}
		})
	}
}
