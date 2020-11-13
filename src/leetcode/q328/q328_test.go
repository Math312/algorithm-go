package q328

import "testing"

/* *
 * [0]  ->  [0]
 * */
func TestOddEvenList(t *testing.T)  {
	root := &ListNode{
		Val:  0,
		Next: nil,
	}
	root = oddEvenList(root)
	if root.Val != 0 {
		t.FailNow()
	}
}

/* *
 * [0,1] -> [0,1]
 * */
func TestOddEvenList2(t *testing.T)  {
	root := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	root = oddEvenList(root)
	result := []int{0,1}
	for _,data := range result {
		if data != root.Val {
			t.FailNow()
		}
		root = root.Next
	}
}

/* *
 * [0,1,2] -> [0,2,1]
 * */
func TestOddEvenList3(t *testing.T)  {
	root := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	root = oddEvenList(root)
	result := []int{0,2,1}
	for _,data := range result {
		if data != root.Val {
			t.FailNow()
		}
		root = root.Next
	}
}

/* *
 * [0,1,2,3] -> [0,2,1,3]
 * */
func TestOddEvenList4(t *testing.T)  {
	root := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	root = oddEvenList(root)
	result := []int{0,2,1,3}
	for _,data := range result {
		if data != root.Val {
			t.FailNow()
		}
		root = root.Next
	}
}


/* *
 * [0,1,2,3,4] -> [0,2,4,1,3]
 * */
func TestOddEvenList5(t *testing.T)  {
	root := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	}
	root = oddEvenList(root)
	result := []int{0,2,4,1,3}
	for _,data := range result {
		if data != root.Val {
			t.FailNow()
		}
		root = root.Next
	}
}

/* *
 * [0,1,2,3,4,5] -> [0,2,4,1,3,5]
 * */
func TestOddEvenList6(t *testing.T)  {
	root := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}
	root = oddEvenList(root)
	result := []int{0,2,4,1,3,5}
	for _,data := range result {
		if data != root.Val {
			t.FailNow()
		}
		root = root.Next
	}
}


