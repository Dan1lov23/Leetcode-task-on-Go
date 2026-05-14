package main

import (
	"fmt"
	"sort"
	"strconv"
)

func splitNum(num int) int {
	digits := []int{}
	for _, ch := range strconv.Itoa(num) {
		digits = append(digits, int(ch-'0'))
	}
	sort.Ints(digits)

	num1, num2 := 0, 0
	for i, d := range digits {
		if i%2 == 0 {
			num1 = num1*10 + d
		} else {
			num2 = num2*10 + d
		}
	}
	return num1 + num2
}

func main() {
	fmt.Println(splitNum(11))
}
