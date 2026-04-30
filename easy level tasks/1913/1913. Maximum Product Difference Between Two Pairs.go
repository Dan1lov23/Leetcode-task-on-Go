package main

import (
	"fmt"
)

func getMax(nums []int) (int, int) {

	max, maxIndex := nums[0], 0

	for i := 0; i <= len(nums) - 1; i++ {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}

	return max, maxIndex

}

func getMin(nums []int) (int, int) {

	min, minIndex := nums[0], 0

	for i := 0; i <= len(nums) - 1; i++ {
		fmt.Println(nums[i])
		if nums[i] < min {
			min = nums[i]
			minIndex = i
		}
	}

	return min, minIndex

}

func maxProductDifference(nums []int) int {

	max1, maxIndex1 := getMax(nums)
	nums = append(nums[:maxIndex1], nums[maxIndex1 + 1:]...)
	
	max2, maxIndex2 := getMax(nums)
	nums = append(nums[:maxIndex2], nums[maxIndex2 + 1:]...)

	min1, minIndex1 := getMin(nums)
	nums = append(nums[:minIndex1], nums[minIndex1 + 1:]...)

	min2, minIndex2 := getMin(nums)
	nums = append(nums[:minIndex2], nums[minIndex2 + 1:]...)

	fmt.Println(max1, max2, min1, min2)

	return max1 * max2 - min1 * min2

}

func main() {
	fmt.Println(maxProductDifference([]int{5, 6, 2, 7, 4}))
}
