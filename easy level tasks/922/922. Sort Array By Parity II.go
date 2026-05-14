package main

import "fmt"

func sortArrayByParityII(nums []int) []int {

	result := []int{}

	evenSlice, notEvenSlice := []int{}, []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			evenSlice = append(evenSlice, nums[i])
		} else {
			notEvenSlice = append(notEvenSlice, nums[i])
		}
	}

	evenCounter, notEvenCounter := 0, 0

	for i := 0; i < len(nums); i++ {

		if i%2 == 0 {
			result = append(result, evenSlice[evenCounter])
			evenCounter++
		} else {
			result = append(result, notEvenSlice[notEvenCounter])
			notEvenCounter++
		}
		fmt.Println(evenCounter, notEvenCounter)
	}

	return result

}

func main() {
	fmt.Println(sortArrayByParityII([]int{2, 3}))
}
