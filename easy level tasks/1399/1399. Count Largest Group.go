package main

import (
	"fmt"
)

func sumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func countLargestGroup(n int) int {
	counts := make(map[int]int)
	maxSize := 0
	for i := 1; i <= n; i++ {
		s := sumOfDigits(i)
		counts[s]++
		if counts[s] > maxSize {
			maxSize = counts[s]
		}
	}
	result := 0
	for _, cnt := range counts {
		if cnt == maxSize {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(countLargestGroup(13))
}
