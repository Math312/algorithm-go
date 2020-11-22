package q242

func isAnagram(s string, t string) bool {
	m := map[rune]int{}
	for _, c := range s {
		_, ok := m[c]
		if !ok {
			m[c] = 1
		} else {
			m[c] ++
		}
	}

	for _, c := range t {
		_, ok := m[c]
		if !ok {
			return false
		} else {
			m[c] --
		}
	}

	for _, value := range m {
		if value != 0 {
			return false
		}
	}
	return true

}
