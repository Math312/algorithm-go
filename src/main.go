package main

import "algorithm/leetcode/q199"

func main() {
	root := q199.TreeNode{
		Val: 1,
		Left: &q199.TreeNode{
			Val: 2,
			Right: &q199.TreeNode{
				Val: 5,
			},
		},
		Right: &q199.TreeNode{
			Val: 3,
			Right: &q199.TreeNode{
				Val: 4,
			},
		},
	}
	q199.RightSideView(&root)
}
