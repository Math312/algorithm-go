package main

import (
	"algorithm/leetcode/q131"
)

func main() {
	for _,data1 := range q131.Partition("") {
		for _,data2 := range data1{
			print(data2)
			print(";")
		}
		println()
	}
}
