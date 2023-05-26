package simple_and_midium

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "abbaca",
			},
			want: "ca",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "abbaca",
			},
			want: "ca",
		},{
			args: args{
				s: "aa",
			},
			want: "",
		},{
			args: args{
				s: "azxxzy",
			},
			want: "ay",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates2(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "abbaca",
			},
			want: "ca",
		},{
			args: args{
				s: "aa",
			},
			want: "",
		},{
			args: args{
				s: "azxxzy",
			},
			want: "ay",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates3(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates3() = %v, want %v", got, tt.want)
			}
		})
	}
}