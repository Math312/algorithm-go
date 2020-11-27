package q454

import "testing"

func TestFourSumCount(t *testing.T) {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}

	if fourSumCount(A, B, C, D) != 2 {
		t.FailNow()
	}
}
