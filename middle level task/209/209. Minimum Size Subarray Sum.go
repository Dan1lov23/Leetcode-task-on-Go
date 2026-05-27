package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {

	length := len(nums)

	sum := 0

	for _, num := range nums {
		sum += num
	}

	if sum < target {
		return 0
	}

	for a := 0; a < len(nums); a++ {

		sum, counter := nums[a], 1

		for b := a + 1; b < len(nums); b++ {
			sum += nums[b]
			counter++
			if sum >= target {
				break
			}
		}

		if length > counter && sum >= target {
			length = counter
		}

		fmt.Println(counter)

	}

	return length

}

func main() {
	fmt.Println(minSubArrayLen(6, []int{10, 2, 3}))
}
