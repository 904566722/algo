package simple_and_midium

import "testing"

func Test_fac(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fac(tt.args.n); got != tt.want {
				t.Errorf("fac() = %v, want %v", got, tt.want)
			}
		})
	}
}
