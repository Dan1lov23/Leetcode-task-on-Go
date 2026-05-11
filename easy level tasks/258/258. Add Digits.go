package main

import (
	"fmt"
	"strconv"
)

func getNumbersDigitsSum(num int) (digitSum int) {

	stringNum := strconv.Itoa(num)

	for i := 0; i <= len(stringNum) - 1; i++ {
		number, err := strconv.Atoi(string(stringNum[i]))
		if err == nil {
			digitSum += number
		}
	}

	return

}

func addDigits(num int) int {

	result := getNumbersDigitsSum(num)

	for {
		if result <= 9 {
			return result
		}
		result = getNumbersDigitsSum(result)
	}

}

func main() {
	fmt.Println(addDigits(38))
}
