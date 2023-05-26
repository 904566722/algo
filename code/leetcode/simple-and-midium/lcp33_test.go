package simple_and_midium

import "testing"

func Test_storeWater(t *testing.T) {
	type args struct {
		bucket []int
		vat    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				bucket: []int{9,0,1},
				vat: []int{0,2,2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storeWater(tt.args.bucket, tt.args.vat); got != tt.want {
				t.Errorf("storeWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
