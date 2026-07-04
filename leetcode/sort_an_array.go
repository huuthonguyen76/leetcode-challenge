package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

func baselineSort(nums []int) []int {
	for i := 0; i < 1000; i++ {
		nums = append(nums, rand.Intn(1000))
	}
	return nums
}

func test(nums []int) {
	fmt.Println(nums)
	sort.Ints(nums)
}

// Merge Sort
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	leftArr := sortArray(append([]int(nil), nums[:len(nums)/2]...))
	rightArr := sortArray(append([]int(nil), nums[len(nums)/2:]...))

	leftCursor := 0
	rightCursor := 0
	result := []int{}
	for {
		if leftCursor >= len(leftArr) {
			result = append(result, rightArr[rightCursor:]...)
			break
		}
		if rightCursor >= len(rightArr) {
			result = append(result, leftArr[leftCursor:]...)
			break
		}
		if leftArr[leftCursor] < rightArr[rightCursor] {
			result = append(result, leftArr[leftCursor])
			leftCursor++
		} else {
			result = append(result, rightArr[rightCursor])
			rightCursor++
		}
	}

	return result
}

func main() {
	arr := []int{5, 1, 1, 2, 0, 0}

	baselineArr := baselineSort(arr)
	sortArr := sortArray(arr)

	if !reflect.DeepEqual(baselineArr, sortArr) {
		fmt.Println("Sorting algorithm is not correct")
	}
}
