package simple_and_midium

import "testing"

func Test_equalFrequency(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				word: "abbcc",
			},
			want: true,
		},
		{
			args: args{
				word: "abc",
			},
			want: true,
		},{
			args: args{
				word: "abbcc",
			},
			want: true,
		},{
			args: args{
				word: "zz",
			},
			want: true,
		},{
			args: args{
				word: "cccd",
			},
			want: true,
		},{
			args: args{
				word: "ddaccb",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalFrequency(tt.args.word); got != tt.want {
				t.Errorf("equalFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
