package hill_sort

type HillSort struct {
}

func (HillSort) Sort(source *[]int) {
	length := len(*source)
	h := 1
	for h < length/3 {
		h = 3 * h + 1
	}
	for h >= 1 {
		for i := h; i < length; i +=h  {
			for j := i;j >= h;j -=h {
				if (*source)[j] < (*source)[j - h]  {
					temp := (*source)[j]
					(*source)[j] = (*source)[j-h]
					(*source)[j-h] = temp
				} else {
					break
				}
			}
		}
		h = h / 3
	}
}
