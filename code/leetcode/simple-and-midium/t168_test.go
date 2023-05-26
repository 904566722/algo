package simple_and_midium

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				columnNumber: 78,
			},
			want: "BZ",
		},
		{
			args: args{
				columnNumber: 26,
			},
			want: "Z",
		},
		{
			args: args{
				columnNumber: 27,
			},
			want: "AA",
		},
		{
			args: args{
				columnNumber: 701,
			},
			want: "ZY",
		},
		{
			args: args{
				columnNumber: 1,
			},
			want: "A",
		}, {
			args: args{
				columnNumber: 2147483647,
			},
			want: "FXSHRXW",
		},
		{
			args: args{
				columnNumber: 28,
			},
			want: "AB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
