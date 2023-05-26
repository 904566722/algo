package simple_and_midium

import (
	"reflect"
	"testing"
)

func Test_rightSideView2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView2() = %v, want %v", got, tt.want)
			}
		})
	}
}
