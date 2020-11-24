package q222

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	deepth := 0
	temp := root
	for temp.Left != nil {
		deepth++
		temp = temp.Left
	}
	high := 1<<(deepth+1) - 1
	low := 1 << deepth
	return innerCount(root, int(low), int(high), deepth)
}

func innerCount(root *TreeNode, low int, high int, deepth int) int {
	for low < high {
		mid := (low + high + 1) / 2
		if exists(root, mid, deepth) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

func exists(root *TreeNode, num int, index int) bool {
	temp := 1 << (index - 1)
	tempNode := root
	for temp > 0 {
		bit := num & temp
		if bit != 0 {
			if tempNode.Right != nil {
				tempNode = tempNode.Right
			} else {
				return false
			}
		} else {
			if tempNode.Left != nil {
				tempNode = tempNode.Left
			} else {
				return false
			}
		}
		temp >>= 1
	}
	return true
}
