package main

import "fmt"

func majorityElement(nums []int) int {

	for a := 0; a < len(nums); a++ {

		element, sum := nums[a], 0

		for b := 0; b < len(nums); b++ {
			if nums[a] == nums[b] && a != b {
				sum += nums[b]
			}
		}

		if len(nums) % 2 != 0 && sum > (len(nums) - 1) / 2 {
			return element
		} else if len(nums) % 2 == 0 && sum > len(nums) / 2 {
			return element
		}
	}

	return 0

}

func main() {
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}
