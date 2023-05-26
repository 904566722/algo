package simple_and_midium

import "testing"

func Test_havePathToLeaf(t *testing.T) {
	type args struct {
		root   *TreeNode
		tgtSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := havePathToLeaf(tt.args.root, tt.args.tgtSum); got != tt.want {
				t.Errorf("havePathToLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
