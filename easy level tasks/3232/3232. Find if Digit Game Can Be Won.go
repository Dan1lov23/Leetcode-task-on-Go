package main

import "fmt"

func canAliceWin(nums []int) bool {

	singleDigitSum := 0
	doubleDigitSum := 0

	for _, num := range nums {
		if num < 10 {
			singleDigitSum += num
		} else {
			doubleDigitSum += num
		}
	}

	if singleDigitSum == doubleDigitSum {
		return false
	}

	return true

}

func main() {
	fmt.Println(canAliceWin([]int{5, 5, 5, 25}))
}
