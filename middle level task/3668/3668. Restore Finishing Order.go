package main

import "fmt"

func recoverOrder(order []int, friends []int) []int {

	var sortedQueue []int

	for a := 0; a <= len(order)-1; a++ {
		for b := 0; b <= len(friends)-1; b++ {
			if order[a] == friends[b] {
				sortedQueue = append(sortedQueue, order[a])
			}
		}
	}

	return sortedQueue

}

func main() {
	fmt.Println(recoverOrder([]int{1, 4, 5, 3, 2}, []int{2, 5}))
}
