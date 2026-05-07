package main

import (
	"fmt"
)

func uniqueElementCheck(nums []int, element int, elementIndex int) (flag bool) {

	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] == element && elementIndex != i {
			return false
		}
	} 

	return true

}

func sumOfUnique(nums []int) int {

	sum := 0

	for i := 0; i <= len(nums)-1; i++ {
		if uniqueElementCheck(nums, nums[i], i) {
			sum += nums[i]
		}
	}

	return sum

}

func main() {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))
}
