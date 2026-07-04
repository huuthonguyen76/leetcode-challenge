package leetcode

import "fmt"

func twoSum(nums []int, target int) []int {
	var surplusMap = make(map[int]int)

	for i, ele_1 := range nums {
		curSurplus := target - ele_1

		curPos, isExist := surplusMap[curSurplus]
		if curPos == i && isExist {
			continue
		}
		if isExist {
			return []int{i, curPos}
		}

		surplusMap[ele_1] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
