package main

import (
	"fmt"
	"strconv"
)

func reverseNum(num int) (reverseNum int) {

	strNum := strconv.Itoa(num)
	reverseStr := ""

	for i := len(strNum) - 1; i >= 0; i-- {
		reverseStr += string(strNum[i])
	}

	result, err := strconv.Atoi(reverseStr)

	if err == nil {
		reverseNum = result
	}

	return

}

func mirrorDistance(n int) int {
    
	return reverseNum(n) - n

}

func main() {
	fmt.Println(mirrorDistance(25))
}
