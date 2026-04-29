package main

import "fmt"

func elementInclude(arr []int, element int) bool {

	for _, num := range arr {
		if num == element {
			return true
		}
	}

	return false

}

func findIntersectionValues(nums1 []int, nums2 []int) []int {

	var counter1, counter2 int

	for _, num := range nums1 {
		
		if elementInclude(nums2, num) {
			counter1++
		}
	}

	for _, num := range nums2 {
		
		if elementInclude(nums1, num) {
			counter2++
		}
	}

	return []int{counter1, counter2}

}

func main() {

	fmt.Println(findIntersectionValues([]int{4,3,2,3,1}, []int{2,2,5,2,3,6}))

}
