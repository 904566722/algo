package simple_and_midium

import (
	"testing"
)

func Test_findTarget(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 10,
							Right: &TreeNode{
								Val: 5,
								Left: nil,
								Right: &TreeNode{
									Val: 7,
									Left: nil,
									Right: nil,
								},
							},
						},
						Right: &TreeNode{
							Val: 9,
							Left: nil,
							Right: &TreeNode{
								Val: 11,
								Left: nil,
								Right: nil,
							},
						},
					},
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
