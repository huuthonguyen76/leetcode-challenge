package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	Val  int
	Num1 int
	Num2 int
}

// MinHeap implements heap.Interface for a min-heap of Nodes
type MinHeap []Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}

	h := &MinHeap{}
	heap.Init(h)

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			heap.Push(h, Node{
				Val:  num1 + num2,
				Num1: num1,
				Num2: num2,
			})
		}
	}

	result := [][]int{}
	for h.Len() > 0 && len(result) < k {
		node := heap.Pop(h).(Node)
		result = append(result, []int{node.Num1, node.Num2})
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 4, 5, 6}
	nums2 := []int{3, 5, 7, 9}
	k := 3

	output := kSmallestPairs(nums1, nums2, k)
	fmt.Println("Result", output)

}
