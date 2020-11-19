package q83

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	first := head.Next
	before := head
	for first != nil {
		if before.Val == first.Val {
			before.Next = first.Next
		} else {
			before = before.Next
		}
		first = first.Next
	}
	return head
}
