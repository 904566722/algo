package simple_and_midium

import "testing"

func Test_mostCommonWord(t *testing.T) {
	type args struct {
		paragraph string
		banned    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
				banned: []string{"hit"},
			},
			want: "ball",
		},{
			args: args{
				paragraph: "a, a, a, a, b,b,b,c, c",
				banned: []string{"a"},
			},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCommonWord(tt.args.paragraph, tt.args.banned); got != tt.want {
				t.Errorf("mostCommonWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
