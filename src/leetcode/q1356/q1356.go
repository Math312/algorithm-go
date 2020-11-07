package q1356

// 可优化

import (
	"container/list"
	"sort"
)

func sortByBits(arr []int) []int {
	result := make([]int, len(arr))
	countOneMap := map[int]*list.List{}
	maxCount := 0
	for _, num := range arr {
		count := count1(num)
		if count > maxCount {
			maxCount = count
		}
		numList, ok := countOneMap[count]
		if !ok {
			numList = list.New()
			numList.PushBack(num)
			countOneMap[count] = numList
		} else {
			numList.PushBack(num)
		}
	}
	index := 0
	for i := 0; i <= maxCount; i++ {
		numList, ok := countOneMap[i]
		if !ok {
			continue
		}
		startIndex := index
		for j := numList.Front(); j != nil; j = j.Next() {
			result[index] = j.Value.(int)
			index ++
		}
		sort.Ints(result[startIndex:index])
	}
	return result
}

func count1(data int) int {
	if data == 0 {
		return 0
	}
	result := 1
	for data&(data-1) != 0 {
		result++
		data = data&(data-1)
	}
	return result
}
