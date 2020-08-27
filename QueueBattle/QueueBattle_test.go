package kata

import (
	"reflect"
	"testing"
)

func TestQueueBattle(t *testing.T) {
	type args struct {
		dist   uint32
		armies [][]uint32
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []uint32
	}{
		{"1", args{100, [][]uint32{{25, 38, 55, 46, 82}, {64, 90, 37, 25, 58}}}, 1, []uint32{3, 2}},
		{"2", args{200, [][]uint32{{61, 83, 37, 55, 92, 35, 68, 72}, {90, 81, 36, 114, 67, 25, 31, 84}}}, 0, []uint32{6, 7}},
		{"3", args{300, [][]uint32{{98, 112, 121, 95, 63}, {120, 94, 90, 88, 30}, {116, 144, 45, 200, 32}}}, 0, []uint32{2}},
		{"4", args{400, [][]uint32{{186, 78, 56, 67, 78, 127, 78, 192}, {78, 67, 208, 45, 134, 212, 82, 99}, {327, 160, 49, 246, 109, 98, 44, 57}}}, 2, []uint32{0, 2, 5}},
		{"5", args{500, [][]uint32{{345, 168, 122, 269, 151}, {56, 189, 404, 129, 101}, {364, 129, 209, 163, 379}, {520, 224, 154, 74, 420}}}, -1, []uint32{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := QueueBattle(tt.args.dist, tt.args.armies...)
			if got != tt.want {
				t.Errorf("QueueBattle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("QueueBattle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
