package main

import (
	"fmt"
	"math"
	"sort"
)

func sum(nums []int) float64 {
	total := 0
	for _, num := range nums {
		total += num
	}
	return float64(total)
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) <= 3 {
		return int(sum(nums))
	}

	sort.Ints(nums)

	targetFloat := float64(target)

	numsFloat := []float64{}
	for _, num := range nums {
		numsFloat = append(numsFloat, float64(num))
	}

	smallestDistance := 99999.0
	smallestSum := 0.0

	for i := 2; i <= len(numsFloat)-1; i++ {
		cursor := i
		left := 0
		right := cursor - 1

		curSum := numsFloat[left] + numsFloat[right] + numsFloat[cursor]
		curDistance := math.Abs(targetFloat - curSum)

		for {
			if left >= right-1 {
				break
			}

			curSumLeft := numsFloat[left+1] + numsFloat[right] + numsFloat[cursor]
			curSumRight := numsFloat[left] + numsFloat[right-1] + numsFloat[cursor]
			curDistanceLeft := math.Abs(targetFloat - curSumLeft)
			curDistanceRight := math.Abs(targetFloat - curSumRight)

			if curDistanceLeft <= curDistanceRight {
				if curDistance >= curDistanceLeft {
					curDistance = curDistanceLeft
					curSum = curSumLeft
				}
				left++
			} else {
				if curDistance >= curDistanceRight {
					curDistance = curDistanceRight
					curSum = curSumRight
				}
				right--
			}

			// fmt.Println(curSum, curDistance, curDistanceLeft, curDistanceRight)
			// fmt.Println(left, right)
			// fmt.Println("==========")
		}

		if curDistance <= smallestDistance {
			smallestDistance = curDistance
			smallestSum = curSum
		}

	}

	return int(smallestSum)
}

func main() {
	nums := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	target := 1
	output := threeSumClosest(nums, target)
	fmt.Println("Result", output)
	if output != 60 {
		fmt.Println("Wrong")
	}
}
