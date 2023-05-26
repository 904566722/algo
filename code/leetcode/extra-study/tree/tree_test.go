package tree

import "testing"

func Test_pOrder(t *testing.T) {
	type args struct {
		root *treeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				root: &treeNode{
					val: 1,
					left: &treeNode{
						val: 2,
						left: &treeNode{
							val: 4,
						},
						right: &treeNode{
							val: 5,
						},
					},
					right: &treeNode{
						val: 3,
						left: &treeNode{
							val: 6,
						},
						right: &treeNode{
							val: 7,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			levelOrder(tt.args.root)
		})
	}
}

func Test_pOrderStack(t *testing.T) {
	type args struct {
		root *treeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				root: &treeNode{
					val: 1,
					left: &treeNode{
						val: 2,
						left: &treeNode{
							val: 4,
						},
						right: &treeNode{
							val: 5,
						},
					},
					right: &treeNode{
						val: 3,
						left: &treeNode{
							val: 6,
						},
						right: &treeNode{
							val: 7,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pOrderStack(tt.args.root)
		})
	}
}