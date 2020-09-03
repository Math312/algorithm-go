package insert_sort

type InsertSort struct {
	
}

func (InsertSort) Sort(source *[]int)  {
	if len(*source) == 0 {
		return
	}
	for i,_ := range (*source)[1:] {
		index := i + 1
		for index > 0 {
			if (*source)[index] < (*source)[index - 1] {
				temp := (*source)[index - 1]
				(*source)[index - 1] = (*source)[index]
				(*source)[index] = temp
			} else {
				break
			}
			index --
		}
	}
}
