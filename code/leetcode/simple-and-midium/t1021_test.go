package simple_and_midium

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
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
				s: "(()())(())",
			},
			want: "()()()",
		},
		{
			args: args{
				s: "(())((()))",
			},
			want: "()(())",
		},
		{
			args: args{
				s: "(()()())(()(()))",
			},
			want: "()()()()(())",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.args.s); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeOuterParentheses2(t *testing.T) {
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
				s: "(()())(())",
			},
			want: "()()()",
		},
		{
			args: args{
				s: "(())((()))",
			},
			want: "()(())",
		},
		{
			args: args{
				s: "(()()())(()(()))",
			},
			want: "()()()()(())",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses2(tt.args.s); got != tt.want {
				t.Errorf("removeOuterParentheses2() = %v, want %v", got, tt.want)
			}
		})
	}
}