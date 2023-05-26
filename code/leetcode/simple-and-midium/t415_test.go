package simple_and_midium

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				num1: "11",
				num2: "123",
			},
			want: "134",
		}, {
			args: args{
				num1: "123",
				num2: "11",
			},
			want: "134",
		}, {
			args: args{
				num1: "",
				num2: "11",
			},
			want: "11",
		}, {
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addStrings2(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				num1: "11",
				num2: "123",
			},
			want: "134",
		}, {
			args: args{
				num1: "123",
				num2: "11",
			},
			want: "134",
		}, {
			args: args{
				num1: "",
				num2: "11",
			},
			want: "11",
		}, {
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings2(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings2() = %v, want %v", got, tt.want)
			}
		})
	}
}
