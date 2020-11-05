package q941

func validMountainArray(A []int) bool {
	label := true
	if len(A) < 3 {
		return false
	}
	left, right := 0, 0
	for i, data := range A {
		if i < len(A)-1 {
			if data == A[i+1] {
				return false
			}
			if label {
				if data < A[i+1] {
					left ++
					continue
				} else {
					right ++
					label = false
				}
			} else {
				if data > A[i+1] {
					right ++
					continue
				} else {
					return false
				}
			}

		}
	}
	return left > 0 && right > 0
}
