package main

import (
	"fmt"
	"slices"
)

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {

	result := []int{}
	mainSlice := [][]int{nums1, nums2, nums3}

	for _, slice := range mainSlice {
		fmt.Println(slice)
		for a := 0; a < len(slice); a++ {
			counter := 0
			for _, sliceElement := range mainSlice {
				if slices.Contains(sliceElement, slice[a]) {
					counter++
					if counter == 2 && !slices.Contains(result, slice[a]) {
						result = append(result, slice[a])
					}
				}
			}
		}
	}

	return result

}

func main() {
	fmt.Println(twoOutOfThree([]int{1,1,3,2}, []int{2,3}, []int{3}))
}
