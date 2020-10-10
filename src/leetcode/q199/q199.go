package q199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var space []*TreeNode
	space = append(space, root)
	label := root
	var result []int
	for len(space) != 0 {
		temp := space[0]
		if temp == label {
			result = append(result, temp.Val)
			if temp.Left == nil && temp.Right == nil {
				label = space[len(space)-1]
			} else {
				if temp.Left != nil {
					label = temp.Left
				}
				if temp.Right != nil {
					label = temp.Right
				}
			}
		}
		if temp.Left != nil {
			space = append(space, temp.Left)
		}
		if temp.Right != nil {
			space = append(space, temp.Right)
		}
		space = space[1:]
	}
	return result
}
