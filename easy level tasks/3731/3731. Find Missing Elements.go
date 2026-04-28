package main

import "fmt"

func findMinInArr(nums []int) (int, int) {

	minElement := nums[0]
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < minElement {
			minElement = nums[i]
			index = i
		}
	}

	return minElement, index

}

func sort(nums []int) (sortedArr []int) {

	for counter := len(nums); counter != 0; counter-- {
		minElement, index := findMinInArr(nums)
		sortedArr = append(sortedArr, minElement)
		nums = append(nums[:index], nums[index+1:]...)
	}

	return

}

func findMissingElements(nums []int) []int {

	missingElements := []int{}
	sortArr := sort(nums)

	for i := 0; i <= len(sortArr)-1; i++ {
		if i+1 < len(sortArr) && sortArr[i]+1 != sortArr[i+1] {
			missingElement := sortArr[i] + 1
			missingElements = append(missingElements, missingElement)
			for num := missingElement + 1; num+1 <= sortArr[i+1]; num++ {
				missingElements = append(missingElements, num)
			}
		}
	}

	return missingElements

}

func main() {
	fmt.Println(findMissingElements([]int{8, 1, 6, 3, 4}))
}
