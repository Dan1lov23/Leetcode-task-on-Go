package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {

	result := []int{}

	for a := 0; a < len(nums); a++ {
		counter := 0
		for b := 0; b < len(nums); b++ {
			if nums[a] > nums[b] {
				counter++
			}
		}
		result = append(result, counter)
	}

	return result

}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}
