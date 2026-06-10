package main

import "fmt"

func getTwoMaxNums(nums []int) (max1 int, max2 int) {

	max1Index, max2Index := 0, 0

	for a := 0; a < len(nums); a++ {
		if nums[a] > max1 {
			max1, max1Index = nums[a], a
		}
	}

	nums = append(nums[:max1Index], nums[max1Index+1:]...)

	for b := 0; b < len(nums); b++ {
		if nums[b] > max2 {
			max2, max2Index = nums[b], b
		}
	}

	nums = append(nums[:max2Index], nums[max2Index+1:]...)

	return

}

func maxProduct(nums []int) int {
	max1, max2 := getTwoMaxNums(nums)
	return (max1 - 1) * (max2 - 1)
}

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
}
