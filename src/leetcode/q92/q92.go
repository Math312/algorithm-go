package q92

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    vHead := &ListNode{
        Val:  0,
        Next: head,
    }
    startLast := vHead
    for i:=0;i < m - 1;i ++ {
        startLast = startLast.Next
    }
    middleEnd := startLast.Next
    startLast.Next = nil
    tempHead := middleEnd
    var tempTail *ListNode
    for i:= m;i <= n;i ++ {
        temp:= tempHead
        tempHead = tempHead.Next
        temp.Next = tempTail
        tempTail = temp
    }
    startLast.Next = tempTail
    middleEnd.Next = tempHead
    return vHead.Next
}
