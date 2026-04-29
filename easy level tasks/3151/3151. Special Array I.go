package main

import "fmt"

func numCheck(num int) bool {

	if num % 2 == 0 {
		return true
	}

	return false

}

func isArraySpecial(nums []int) bool {

	if len(nums) == 0 && len(nums) == 1 {
		return true
	}

	for i := 0; i < len(nums) - 1; i++ {

		if numCheck(nums[i]) && !numCheck(nums[i + 1]) || !numCheck(nums[i]) && numCheck(nums[i + 1]) {
			//
		} else {
			return false
		}

	}

	return true

}

func main() {

	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}))	

}
