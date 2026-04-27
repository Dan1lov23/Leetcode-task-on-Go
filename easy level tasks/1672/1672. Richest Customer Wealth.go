package main

import "fmt"

func getArrElementsSum(arr []int) (sum int) {

	for _, arrElement := range arr {
		sum += arrElement
	}

	return

}

func maximumWealth(accounts [][]int) int {

	max := 0

	for _, arr := range accounts {
		if getArrElementsSum(arr) > max {
			max = getArrElementsSum(arr)
		}
	}

	return max

}

func main() {
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
}
