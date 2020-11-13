package q328

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	temp := head
	if temp == nil {
		return nil
	}
	for !(temp.Next == nil || temp.Next.Next == nil) {
		temp = temp.Next.Next
	}
	temp3 := temp
	temp2last := head
	temp2 := head.Next
	for temp2last != temp {
		temp2last.Next = temp2.Next
		temp2.Next = temp3.Next
		temp3.Next = temp2
		temp2 = temp2last.Next.Next
		temp2last = temp2last.Next
		temp3 = temp3.Next
	}
	return head
}