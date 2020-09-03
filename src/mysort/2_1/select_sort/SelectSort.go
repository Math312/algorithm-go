package select_sort

import "algorithm/mysort"

type SelectSort struct {

}

func (SelectSort) Sort(source *[]int)  {
	index := 0
	for index < len(*source) {
		min_index := 0
		split := (*source)[index: ]
		for i,num := range split{
			if split[min_index] > num {
				min_index = i
			}
		}
		if min_index != 0 {
			temp := split[0]
			split[0] = split[min_index]
			split[min_index] = temp
		}
		index += 1
		mysort.SortUtil{}.PrintNumArray(*source)
	}

}


