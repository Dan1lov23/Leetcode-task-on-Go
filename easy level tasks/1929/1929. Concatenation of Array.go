package main

import "fmt"

func getConcatenation(nums []int) []int {

	newArr := nums

	for i := 0; i <= len(nums) - 1; i++ {
		newArr = append(newArr, nums[i])
	}

	return newArr

}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1}))
}
