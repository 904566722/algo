package simple_and_midium

import "testing"

func Test_lastSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "aabaaaab",
			},
			want: "bab",
		}, {
			args: args{
				s: "abab",
			},
			want: "bab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastSubstring(tt.args.s); got != tt.want {
				t.Errorf("lastSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lastSubstring3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "aabaaaab",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastSubstring3(tt.args.s); got != tt.want {
				t.Errorf("lastSubstring3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structPractice(t *testing.T) {
	structPractice2()
}
