package q461

func HammingDistance(x int, y int) int {
	i := 0
	z := x ^ y
	for z != 0 {
		z &= z -1
	}
	return i
}
