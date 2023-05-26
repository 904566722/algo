package simple_and_midium

import "testing"

func Test_rotateString(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "mndqvdqktyxknfdtklxapbkuxuzwubpiwmqgickuwqishkrgzu",
				goal: "rgzumndqvdqktyxknfdtklxapbkuxuzwubpiwmqgickuwqishk",
			},
			want: true,
		},{
			args: args{
				s: "abcde",
				goal: "cdeab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
