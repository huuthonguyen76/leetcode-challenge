package main

import (
	"cmp"
	"fmt"
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	numsList := []int{}

	for _, num := range nums {
		if _, isExisted := countMap[num]; !isExisted {
			numsList = append(numsList, num)
		}

		countMap[num] += 1
	}

	slices.SortFunc(numsList, func(num1, num2 int) int {
		return cmp.Compare(
			countMap[num1], countMap[num2],
		)
	})

	return numsList[len(numsList)-k:]
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}
