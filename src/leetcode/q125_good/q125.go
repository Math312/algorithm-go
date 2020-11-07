package q125_good

func isPalindrome(s string) bool {
	head, tail := 0, len(s)-1
	for head < tail {
		for head <= tail && !checkValid(s[head]) {
			head++
		}
		for tail >= head && !checkValid(s[tail]) {
			tail--
		}
		if head > tail {
			return true
		} else {
			if !checkEquals(s[head],s[tail]) {
				return false
			}
			head++
			tail--
		}
	}
	return true
}

func checkValid(c uint8) bool {
	return (c >= 'A' && c <= 'Z')||(c >= 97 && c <= 122) || (c >= '0' && c <= '9')
}

func checkEquals(c1, c2 uint8) bool {
	if (c1 >= '0' && c1 <= '9') && (c2 >= '0' && c2 <= '9') {
		return c1 == c2
	}
	if c1 >= 'a' {
		c1 = c1 - ('a' - 'A')
	}
	if c2 >= 'a' {
		c2 = c2 - ('a' - 'A')
	}
	return c1 == c2
}
