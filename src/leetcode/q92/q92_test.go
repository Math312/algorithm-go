package q92

import "testing"

func TestReverseBetween(t *testing.T)  {
	root := &ListNode{
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
	}
	root = reverseBetween(root,1,3)
	result := []int{3,2,1,4,5}
	for _,data := range result {
		if data != root.Val {
			t.FailNow()
		}
		root = root.Next
	}
}

func TestReverseBetween2(t *testing.T)  {
	root := &ListNode{
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
	}
	root = reverseBetween(root,1,1)
	result := []int{1,2,3,4,5}
	for _,data := range result {
		if data != root.Val {
			t.FailNow()
		}
		root = root.Next
	}
}

func TestReverseBetween3(t *testing.T)  {
	root := &ListNode{
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
	}
	root = reverseBetween(root,3,3)
	result := []int{1,2,3,4,5}
	for _,data := range result {
		if data != root.Val {
			t.FailNow()
		}
		root = root.Next
	}
}

func TestReverseBetween4(t *testing.T)  {
	root := &ListNode{
		Val:  1,
		Next: nil,
	}
	root = reverseBetween(root,1,1)
	result := []int{1}
	for _,data := range result {
		if data != root.Val {
			t.FailNow()
		}
		root = root.Next
	}
}