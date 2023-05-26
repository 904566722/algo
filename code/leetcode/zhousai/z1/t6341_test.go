package z1

import "testing"

func Test_isWinner(t *testing.T) {
	type args struct {
		player1 []int
		player2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				player1: []int{10,2,23},
				player2: []int{3,8,4,5},
			},
			want: 1,
		}, {
			args: args{
				player1: []int{9,8,5,3,7},
				player2: []int{8,7,4,9,0},
			},
			want: 1,
		},{
			args: args{
				player1: []int{0,1,8},
				player2: []int{3,9,1},
			},
			want: 2,
		},{
			args: args{
				player1: []int{2,3},
				player2: []int{4,1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWinner(tt.args.player1, tt.args.player2); got != tt.want {
				t.Errorf("isWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
