package main

import "fmt"

func includeCheck(nums []int, element int) bool {

	for _, num := range nums {
		if num == element {
			return true
		}
	}

	return false

}

func intersection(nums1 []int, nums2 []int) []int {

	result := []int{}

	for _, num := range nums1 {
		if includeCheck(nums2, num) && !includeCheck(result, num) {
			result = append(result, num)
		}
	}

	return result

}

func main() {
	fmt.Println(intersection([]int{4, 9, 5}, []int{9,4,9,8,4}))
}
