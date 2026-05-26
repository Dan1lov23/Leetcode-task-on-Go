package main

import "fmt"

func includeCounter(nums []int, element int) (counter int) {

	for _, numElement := range nums {
		if numElement == element {
			counter++
		}
	}

	return

} 

func singleNumber(nums []int) int {

	for _, numElement := range nums {
		if includeCounter(nums, numElement) == 1 {
			return numElement
		}
	}

	return -1

}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}
