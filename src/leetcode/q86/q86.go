package q86

type ListNode struct {
	Val  int
	Next *ListNode
}
// 待优化
func partition(head *ListNode, x int) *ListNode {
	result := &ListNode{}
	tempNode := result
	for temp := head; temp != nil; temp = temp.Next {
		if temp.Val >= x {
			continue
		} else {
			tempNode.Next = &ListNode{
				Val: temp.Val,
			}
			tempNode = tempNode.Next
		}
	}
	for temp := head; temp != nil; temp = temp.Next {
		if temp.Val >= x {
			tempNode.Next = &ListNode{
				Val: temp.Val,
			}
			tempNode = tempNode.Next
		} else {
			continue
		}
	}
	return result.Next
}
