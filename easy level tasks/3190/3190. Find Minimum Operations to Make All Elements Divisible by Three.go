package main

import "fmt"

func minimumOperations(nums []int) int {

	operationsCounter := 0

	for _, num := range nums {
		if num%3 != 0 {
			operationsCounter++
		}
	}

	return operationsCounter

}

func main() {
	fmt.Println(minimumOperations([]int{1, 2, 3, 4}))
}
