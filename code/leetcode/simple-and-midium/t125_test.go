package simple_and_midium

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "0P",
			},
			want: false,
		},
		{
			args: args{
				s: "race a car",
			},
			want: false,
		}, {
			args: args{
				s: "",
			},
			want: true,
		}, {
			args: args{
				s: "  ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
