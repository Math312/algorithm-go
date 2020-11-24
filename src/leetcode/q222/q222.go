package q222

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodeList := list.List{}
	nodeList.PushBack(root)
	count := 0
	for nodeList.Len() > 0 {
		node := nodeList.Front();
		nodeList.Remove(node)
		count ++
		temp := node.Value.(*TreeNode)
		if temp.Left != nil {
			nodeList.PushBack(temp.Left)
		}
		if temp.Right != nil {
			nodeList.PushBack(temp.Right)
		}
	}
	return count
}
