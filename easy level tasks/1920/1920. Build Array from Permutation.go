package main

import "fmt"

func buildArray(nums []int) []int {

	var newArr []int

	for a := 0; a < len(nums); a++ {
        newArr = append(newArr, nums[nums[a]])
	}

	return newArr

}

func main() {
	fmt.Println(buildArray([]int{0, 2, 3, 1}))
}
