package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	num1, num2, sum, nums2Idx int
}

type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := [][]int{}
	h := &MinHeap{}

	// Push the first pair of each element in nums1 with the first element of nums2
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, Pair{num1: nums1[i], num2: nums2[0], sum: nums1[i] + nums2[0], nums2Idx: 0})
	}

	// Extract k smallest pairs
	for len(*h) > 0 && len(result) < k {
		currentPair := heap.Pop(h).(Pair)
		result = append(result, []int{currentPair.num1, currentPair.num2})

		// If there is a next element in nums2 for the current nums1 element,
		// add the new pair to the heap.
		if currentPair.nums2Idx+1 < len(nums2) {
			nextNums2Idx := currentPair.nums2Idx + 1
			heap.Push(h, Pair{
				num1:    currentPair.num1,
				num2:    nums2[nextNums2Idx],
				sum:     currentPair.num1 + nums2[nextNums2Idx],
				nums2Idx: nextNums2Idx,
			})
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	result := kSmallestPairs(nums1, nums2, k)
	fmt.Println(result) // Output: [[1,2],[1,4],[1,6]]

	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2
	result = kSmallestPairs(nums1, nums2, k)
	fmt.Println(result) // Output: [[1,1],[1,1]]
}
