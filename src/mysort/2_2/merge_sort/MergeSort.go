package merge_sort

type MergeSort struct {

}

func (MergeSort) Sort(source *[]int)  {
	merge(source)
}

func merge(source *[]int)  {
	if len(*source) == 0 || len(*source) == 1 {
		return
	}
	if len(*source) == 2 {
		if (*source)[0] > (*source)[1] {
			temp := (*source)[0]
			(*source)[0] = (*source)[1]
			(*source)[1] = temp
		}
		return
	}
	mid := len(*source)/2
	left := (*source)[:mid-1]
	merge(&left)
	right := (*source)[mid:]
	merge(&right)
	for i:= 1;i < len(*source);i ++ {
		for j := i;j >= 1;j -- {
			if (*source)[j] < (*source)[j - 1] {
				temp := (*source)[j]
				(*source)[j] = (*source)[j-1]
				(*source)[j-1] = temp
			} else {
				break
			}
		}
	}

}
