package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	result := []int{}

	for a := 0; a < len(nums1); a++ {
		flag := false
		indexForStart := 0
		for i := 0; i < len(nums2); i++ {
			if nums2[i] == nums1[a] {
				indexForStart = i
			}
		}
		for b := indexForStart; b < len(nums2); b++ {
			flag = false
			if nums1[a] < nums2[b] {
				result = append(result, nums2[b])
				flag = true
				break
			}
		}
		if !flag {
			result = append(result, -1)
		}
	}

	return result

}

func main() {
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
