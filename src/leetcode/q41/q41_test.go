package q41

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{1, 2, 0}
	if firstMissingPositive(nums) != 3 {
		t.FailNow()
	}
}

func TestFirstMissingPositive1(t *testing.T) {
	nums := []int{1, 2, 0, -1}
	if firstMissingPositive(nums) != 3 {
		t.FailNow()
	}
}

func TestFirstMissingPositive2(t *testing.T) {
	nums := []int{0, -1, 1, 2}
	if firstMissingPositive(nums) != 3 {
		t.FailNow()
	}
}

func TestFirstMissingPositive3(t *testing.T) {
	nums := []int{0}
	if firstMissingPositive(nums) != 1 {
		t.FailNow()
	}
}

func TestFirstMissingPositive4(t *testing.T) {
	nums := []int{1}
	if firstMissingPositive(nums) != 2 {
		t.FailNow()
	}
}

func TestFirstMissingPositive5(t *testing.T) {
	nums := []int{0,3,1,2,4,5}
	if firstMissingPositive(nums) != 6 {
		t.FailNow()
	}
}

func TestFirstMissingPositive6(t *testing.T) {
	nums := []int{1,1}
	if firstMissingPositive(nums) != 2 {
		t.FailNow()
	}
}


