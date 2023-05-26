package simple_and_midium

import "testing"

func Test_rmDup(t *testing.T) {
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
				s: "aaaaa",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rmDup(tt.args.s); got != tt.want {
				t.Errorf("rmDup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_similarPairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				words: []string{"aba","aabb","abcd","bac","aabc"},
			},
			want: 2,
		},
		{
			args: args{
				words: []string{"aabb","ab","ba"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := similarPairs(tt.args.words); got != tt.want {
				t.Errorf("similarPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}