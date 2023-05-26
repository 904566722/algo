package simple_and_midium

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				gas: []int{0,0,0,0,0,0,0,0,0,0,0,0,2},
				cost: []int{0,0,0,0,0,0,0,0,0,0,0,1,0},
			},
		},{
			args: args{
				gas: []int{5,8,2,8},
				cost: []int{6,5,6,6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
