package s1

func rangeBitwiseAnd(m int, n int) int {
	i := 0
	f := 0
	if m < n {
		f = f | m
	} else {
		f = f | n
	}
	s := m | n
	for f < s {
		s &= s - 1
		i ++
	}
	return i
}