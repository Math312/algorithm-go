package q327

// 前缀和
func countRangeSum(nums []int, lower int, upper int) int {
	numsProcessed := make([]int, len(nums))
	temp:= 0
	for i,num := range nums {
		temp += num
		numsProcessed[i] = temp
	}
	result := 0
	for i:= 0;i < len(numsProcessed); i ++ {
		for j:= i;j < len(numsProcessed);j ++ {
			if j == i {
				temp = numsProcessed[i]
			} else {
				temp = numsProcessed[j] - numsProcessed[i]
			}
			if temp <= upper && temp >= lower {
				result ++
			}
		}
	}
	return result
}
