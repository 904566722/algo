package simple_and_midium

import "testing"

func Test_isLongPressedName(t *testing.T) {
	type args struct {
		name  string
		typed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				name: "abc",
				typed: "abccccc",
			},
			want: true,
		},{
			args: args{
				name: "abc",
				typed: "aaabbcc",
			},
			want: true,
		},{
			args: args{
				name: "abc",
				typed: "abbbbbc",
			},
			want: true,
		},{
			args: args{
				name: "abc",
				typed: "aabc",
			},
			want: true,
		},{
			args: args{
				name: "abc",
				typed: "bca",
			},
			want: false,
		},{
			args: args{
				name: "abc",
				typed: "abcd",
			},
			want: false,
		},{
			args: args{
				name: "abc",
				typed: "abddc",
			},
			want: false,
		},{
			args: args{
				name: "abc",
				typed: "abcd",
			},
			want: false,
		},{
			args: args{
				name: "alex",
				typed: "aaleexa",
			},
			want: false,
		},{
			args: args{
				name: "alex",
				typed: "aaleexeex",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.args.name, tt.args.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
