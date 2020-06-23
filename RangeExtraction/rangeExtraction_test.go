package rangeextraction

import (
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}},
			"-6,-3-1,3-5,7-11,14,15,17-20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.list); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLastRangeIdxAndLen(t *testing.T) {
	type args struct {
		list []int
		idx  int
	}
	tests := []struct {
		name    string
		args    args
		wantIdx int
		wantLen int
	}{
		{"0 idx", args{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}, 0}, 0, 1},
		{"1 idx", args{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}, 1}, 5, 5},
		{"2 idx", args{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}, 2}, 5, 4},
		{"3 idx", args{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}, 3}, 5, 3},
		{"-2 idx", args{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}, 18}, 19, 2},
		{"-1 idx", args{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}, 19}, 19, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findLastRangeIdxAndLen(tt.args.list, tt.args.idx)
			if got != tt.wantIdx {
				t.Errorf("findLastRangeIdx() gotIdx = %v, want %v", got, tt.wantIdx)
			}
			if got1 != tt.wantLen {
				t.Errorf("findLastRangeIdx() gotLen = %v, want %v", got1, tt.wantLen)
			}
		})
	}
}
