package q66

func plusOne(digits []int) []int {
	var result []int
	expand := true
	for i:= 0;i < len(digits);i ++ {
		if digits[i] != 9 {
			expand = false
		}
	}
	if expand {
		result = make([]int,len(digits)+ 1)
		for i := 1;i < len(digits) + 1;i ++ {
			result[i] = 0
		}
		result[0] = 1
	} else {
		add := 0
		result = digits
		result[len(digits) - 1] = result[len(digits) - 1] + 1
		for i := len(digits) - 1;i >= 0;i -- {
			sum := add + result[i]
			result[i] = sum % 10
			add = sum / 10
		}
	}
	return result
}
