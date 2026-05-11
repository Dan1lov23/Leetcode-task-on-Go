package main

import "fmt"

func twoSum(numbers []int, target int) []int {
    
	for a := 0; a <= len(numbers) - 1; a++ {
		for b := a + 1; b <= len(numbers) - 1; b++ {
			if numbers[a] +  numbers[b] == target {
				return []int{a + 1, b + 1}
			}
		}
	}

	return []int{}

}

func main() {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
