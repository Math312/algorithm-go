package q131
/**
* 需要优化
*/
func partition(s string) [][]string {
	if len(s) == 1 {
		return [][]string{{s}}
	}
	var result [][]string
	for i := 1; i <= len(s); i++ {
		if checkPalindrome(s[0:i]) {
			rightResults := partition(s[i:])
			if len(rightResults) > 0 {
				for _,rightResult := range rightResults{
					basicArray := []string{s[0:i]}
					basicArray = append(basicArray, rightResult...)
					result = append(result, basicArray)
				}
			} else {
				result = append(result, []string{s[0:i]})
			}

		}
	}

	return result
}

func checkPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
