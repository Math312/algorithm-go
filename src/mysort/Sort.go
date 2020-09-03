package mysort

import (
	"fmt"
	"math"
)

type Sort interface {
	Sort(source *[]int)
}

type SortUtil struct {

}

func (SortUtil) IsSorted(source []int) bool  {
	min := math.MinInt32
	for _,num := range source {
		if num < min {
			return false
		} else {
			min = num
		}
	}
	return true
}



func (SortUtil) PrintNumArray(source []int)  {
	for _,num := range source {
		fmt.Printf("%-10d",num)
		fmt.Print(",")
	}
	fmt.Println()
}