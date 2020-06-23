package railfence

import (
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"WEAREDISCOVEREDFLEEATONCE", 3}, "WECRLTEERDSOEEFEAOCAIVDEN"},
		{"2", args{"Hello, World!", 3}, "Hoo!el,Wrdl l"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"WECRLTEERDSOEEFEAOCAIVDEN", 3}, "WEAREDISCOVEREDFLEEATONCE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
