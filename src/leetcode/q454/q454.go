package q454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	group1 := make(map[int]int, len(A)*len(B))
	for _, a := range A {
		for _, b := range B {
			group1[a+b] ++
		}
	}
	rs := 0
	for _, c := range C {
		for _, d := range D {
			rs += group1[-c-d]
		}
	}
	return rs
}
