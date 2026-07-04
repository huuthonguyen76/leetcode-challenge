package main

func kthSmallest(matrix [][]int, k int) int {
	push := func(heap *[]int, num int) {
		*heap = append(*heap, num)
		i := len(*heap) - 1
		for i > 0 {
			parent := (i - 1) / 2
			if (*heap)[parent] > (*heap)[i] {
				(*heap)[parent], (*heap)[i] = (*heap)[i], (*heap)[parent]
			}

			i = parent
		}
	}

	pop := func(heap *[]int) int {
		heapLen := len(*heap)

		min := (*heap)[0]

		(*heap)[0], (*heap)[heapLen-1] = (*heap)[heapLen-1], (*heap)[0]
		*heap = (*heap)[:heapLen-1]

		heapLen = len(*heap)

		i := 0
		for {
			left := i*2 + 1
			right := i*2 + 2

			smallest := i

			if left < heapLen && (*heap)[smallest] > (*heap)[left] {
				smallest = left
			}

			if right < heapLen && (*heap)[smallest] > (*heap)[right] {
				smallest = right
			}

			if smallest == i {
				break
			}

			(*heap)[i], (*heap)[smallest] = (*heap)[smallest], (*heap)[i]
			i = smallest
		}

		return min
	}

	// Build heap with all elements
	heap := []int{}
	for _, row := range matrix {
		for _, val := range row {
			push(&heap, val)
		}
	}

	// Pop k times, the kth pop is the answer
	var result int
	for i := 0; i < k; i++ {
		result = pop(&heap)
	}

	return result
}
