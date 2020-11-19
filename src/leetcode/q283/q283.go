package q283

/* *
* 方法1
* */
func moveZeroes(nums []int)  {
	index := 0
	for _,data := range nums {
		if data != 0 {
			nums[index] = data
			index ++
		}
	}
	for ;index< len(nums); {
		nums[index] = 0;
	}
}
/* *
*  方法2
* */
func moveZeroes2(nums []int)  {
	index := 0
	for i,data := range nums {
		if data != 0 {
			nums[i],nums[index] = nums[index],nums[i]
			index ++
		}
	}
}
