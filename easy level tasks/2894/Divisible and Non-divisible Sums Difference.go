package main

import "fmt"

func differenceOfSums(n int, m int) int {

	divSum := 0
	notDivSum := 0

	for i := 0; i <= n; i++ {
		if i % m == 0 {
			divSum += i
		} else {
			notDivSum += i
		}
	}

	return notDivSum - divSum

}

func main() {
	fmt.Println(differenceOfSums(5, 6))
}
