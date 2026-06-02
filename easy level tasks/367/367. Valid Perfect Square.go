package main

import "fmt"

func isPerfectSquare(num int) bool {

	for i := 0; i <= num; i++ {
		fmt.Println(i * i, num)
		if i * i == num {
			return true
		}
	}

	return false

}

func main() {
	fmt.Println(isPerfectSquare(1))
}
