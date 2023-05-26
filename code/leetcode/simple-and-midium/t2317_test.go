package simple_and_midium

import "testing"

func Test_maximumXOR(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{3,2,4,6},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumXOR(tt.args.nums); got != tt.want {
				t.Errorf("maximumXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
