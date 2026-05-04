package main

import (
	"fmt"
	"strconv"
)

func countDigits(num int) int {

	counter := 0

	stringNum := strconv.Itoa(num)

	for i := 0; i <= len(stringNum) - 1; i++ {

		numberForDiv, err := strconv.Atoi(string(stringNum[i]))
		
		if err != nil {
			// work with error
		}

		if num%numberForDiv == 0 {
			counter++
		}
	}

	return counter

}

func main() {
	fmt.Println(countDigits(1248))
}
