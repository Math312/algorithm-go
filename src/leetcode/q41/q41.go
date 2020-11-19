package q41

func firstMissingPositive(nums []int) int {
	for i, _ := range nums {
		for nums[i] != i+1 {
			if nums[i]-1 >= 0 && nums[i]-1 < len(nums) && nums[nums[i]-1] != nums[i] {
				nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			} else {
				nums[i] = -1
				break
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
