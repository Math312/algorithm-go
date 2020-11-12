package q1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i,data := range nums {
		m[target - data] = i
	}
	for i,data := range nums {
		d,ok := m[data]
		if ok && d != i {
			return []int{i,d}
		}
	}
	return nil
}
