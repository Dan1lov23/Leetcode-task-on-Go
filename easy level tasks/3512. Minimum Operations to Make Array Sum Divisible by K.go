package main

import "fmt"

func getArrElementsSum(arr []int) (sum int) {

	for _, number := range arr {
		sum += number
	}

	return

}

func minOperations(nums []int, k int) int {

	operationCounter := 0

	sum := getArrElementsSum(nums)

	for {
		if sum%k == 0 {
			return operationCounter
		}
		sum--
		operationCounter++
	}

}

func main() {
	fmt.Println(minOperations([]int{3, 9, 7}, 5))
}
