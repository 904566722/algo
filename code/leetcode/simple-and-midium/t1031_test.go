package simple_and_midium

import "testing"

func Test_maxSumTwoNoOverlap(t *testing.T) {
	type args struct {
		nums      []int
		firstLen  int
		secondLen int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:      []int{0, 6, 5, 2, 2, 5, 1, 9, 4},
				firstLen:  1,
				secondLen: 2,
			},
			want: 20,
		}, {
			args: args{
				nums:      []int{4, 5, 14, 16, 16, 20, 7, 13, 8, 15},
				firstLen:  3,
				secondLen: 5,
			},
			want: 109,
		}, {
			args: args{
				nums:      []int{1, 2},
				firstLen:  1,
				secondLen: 1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumTwoNoOverlap(tt.args.nums, tt.args.firstLen, tt.args.secondLen); got != tt.want {
				t.Errorf("maxSumTwoNoOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
