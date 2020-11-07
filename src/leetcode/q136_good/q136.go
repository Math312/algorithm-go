package q136_good

/*
* 用异或做
* 0 ^ a = a
* a ^ a = 0
* 异或满足交换律和结合律
*/

func singleNumber(nums []int) int {
	result := 0
	for _,num := range nums{
		result ^= num
	}
	return result
}