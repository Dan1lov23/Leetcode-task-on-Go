package main

import (
	"fmt"
	"strconv"
)

func squareOfNum(num int) int {
	sum := 0
	stringNum := strconv.Itoa(num)
	for i := 0; i < len(stringNum); i++ {
		digit, err := strconv.Atoi(string(stringNum[i]))
		if err == nil {
			sum += digit * digit
		}
	}
	return sum
}

func isHappy(n int) bool {
	seen := make(map[int]bool)
	for {
		if n == 1 {
			return true
		}
		if seen[n] {
			return false
		}
		seen[n] = true
		n = squareOfNum(n)
	}
}

func main() {
	fmt.Println(isHappy(19))
}