package q141

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针法
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	h1 := head
	h2 := head
	for {
		if h2 == nil || h1 == nil {
			return false
		}
		h1 = h1.Next
		h2 = h2.Next
		if h2 == nil {
			return false
		}
		h2 = h2.Next
		if h1 == h2 {
			return true
		}
	}
}
