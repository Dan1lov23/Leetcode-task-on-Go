package main

import (
	"fmt"
)

func arrayNesting(nums []int) int {
	n := len(nums)
	maxLen := 0
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		if !visited[i] {
			count := 0
			j := i
			for !visited[j] {
				visited[j] = true
				j = nums[j]
				count++
			}
			if count > maxLen {
				maxLen = count
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(arrayNesting([]int{1, 2, 3, 4}))
}
