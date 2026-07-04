package main

import "fmt"

func findPeakElement(nums []int) int {
	newNum := []int{}
	newNum = append(newNum, -9999999999)
	newNum = append(newNum, nums...)
	newNum = append(newNum, -9999999999)

	peak := -1
	for idx, _ := range newNum {
		if idx == 0 {
			continue
		}

		if idx == len(newNum)-1 {
			break
		}
		fmt.Println(newNum[idx-1], newNum[idx], newNum[idx+1])
		if newNum[idx] > newNum[idx-1] && newNum[idx] > newNum[idx+1] {
			peak = idx - 1
		}
	}
	return peak
}

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	// nums := []int{1}
	fmt.Println(findPeakElement(nums))
}
