package q1370

func sortString(s string) string {
	m := ['z' + 1]int{}
	for _, c := range s {
		m[c] ++
	}
	length := len(s)
	result := make([]byte, 0, length)
	for len(result) < length {
		for i := byte('a'); i <= 'z'; i++ {
			if m[i] > 0 {
				result = append(result, i)
				m[i] --
			}
		}

		for i := byte('z'); i >= 'a'; i-- {
			if m[i] > 0 {
				result = append(result, i)
				m[i] --
			}
		}
	}
	return string(result)
}
