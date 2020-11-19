package q39_bad

import (
	"container/list"
	"sort"
)

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	aList := list.New()
	result := &[][]int{}
	for i, data := range candidates {
		aList.PushBack(data)
		combinationSumInner(candidates, target-data, i, aList, result)
		aList.Remove(aList.Back())
	}
	return *result
}

func combinationSumInner(candidates []int, target int, index int, aList *list.List, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		aResult := make([]int, aList.Len())
		j := 0
		for e := aList.Front(); e != nil; e = e.Next() {
			aResult[j] = e.Value.(int)
			j++
		}
		*result = append(*result, aResult)
	} else {
		for i := index; i < len(candidates); i++ {
			for j, data := range candidates[i:] {
				aList.PushBack(data)
				combinationSumInner(candidates, target-data, i+j, aList, result)
				aList.Remove(aList.Back())
			}
		}
	}
}
