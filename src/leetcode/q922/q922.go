package q922

/*
* 24ms 87.36%
* 6.3M 37.57%
 */
func SortArrayByParityII(A []int) []int {
	a, b := 0, 1
	for a < len(A) || b < len(A) {
		if a < len(A) && isOddNumber(A[a]) {
			a += 2
			continue
		}
		if b < len(A) && !isOddNumber(A[b]) {
			b += 2
			continue
		}

		if a < len(A) && !isOddNumber(A[a]) {
			temp := A[b]
			A[b] = A[a]
			A[a] = temp
			continue
		}
		if b < len(A) && isOddNumber(A[b]) {
			temp := A[b]
			A[b] = A[a]
			A[a] = temp
			continue
		}
	}
	return A
}

func isOddNumber(num int) bool {
	return num&1 == 0
}
