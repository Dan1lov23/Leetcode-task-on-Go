package main

import "fmt"

func singleNumber(nums []int) []int {

	uniqueNums := []int{}

	for a := 0; a < len(nums); a++ {
	
		flag := true
		num := nums[a]
	
		for b := 0; b < len(nums); b++ {
			if nums[a] == nums[b] && a != b {
				flag = false
				break
			}
		}

		if flag {
			uniqueNums = append(uniqueNums, num)
		}

	}

	return uniqueNums

}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}
