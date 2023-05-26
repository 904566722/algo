package simple_and_midium

import "testing"

func Test_largestSumAfterKNegations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{3, 2, 2, 4, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumAfterKNegations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestSumAfterKNegations1(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{-4, -2, -3},
				k:    4,
			},
			want: 5,
		},
		{
			args: args{
				nums: []int{4, 2, 3},
				k:    1,
			},
			want: 5,
		}, {
			args: args{
				nums: []int{3, -1, 0, 2},
				k:    3,
			},
			want: 6,
		}, {
			args: args{
				nums: []int{2, -3, -1, 5, -4},
				k:    2,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumAfterKNegations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestSumAfterKNegations2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{-6, -8, -7, 1},
				k:    4,
			},
			want: 20,
		},
		{
			args: args{
				nums: []int{-4, -2, -3},
				k:    4,
			},
			want: 5,
		},
		{
			args: args{
				nums: []int{4, 2, 3},
				k:    1,
			},
			want: 5,
		}, {
			args: args{
				nums: []int{3, -1, 0, 2},
				k:    3,
			},
			want: 6,
		}, {
			args: args{
				nums: []int{2, -3, -1, 5, -4},
				k:    2,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumAfterKNegations2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("largestSumAfterKNegations2() = %v, want %v", got, tt.want)
			}
		})
	}
}
