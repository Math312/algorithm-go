package q222

import "testing"

func TestCountNodes(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,

		},
	}
	if countNodes2(root) != 2 {
		t.FailNow()
	}
}
