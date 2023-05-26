package simple_and_midium

import (
	"reflect"
	"testing"
)

func Test_isSubtree2(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
				},
				subRoot: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree2(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNext(t *testing.T) {
	type args struct {
		pat []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				pat: []int{1, 1, 2, 1, 1, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNext(tt.args.pat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kmpSearch(t *testing.T) {
	type args struct {
		nums []int
		pat  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{1, 2, 2, 1, 2, 2, 1, 3},
				pat:  []int{1, 2, 2, 1, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmpSearch(tt.args.nums, tt.args.pat); got != tt.want {
				t.Errorf("kmpSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
