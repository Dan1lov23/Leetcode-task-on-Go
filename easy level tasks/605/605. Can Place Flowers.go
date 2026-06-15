package main

import "fmt"

func check(nums []int) bool {

	for a := 0; a < len(nums); a++ {
		if nums[a] == 0 && nums[a+1] == 0 && nums[a-1] == 0 {
			return true
		}
	}

	return false

}

func canPlaceFlowers(flowerbed []int, n int) bool {

	counter := 0

	for {
		for a := 0; a < len(flowerbed); a++ {
			if a == 0 && flowerbed[a-1] == 0 && flowerbed[a + 1] == 0 && a != 0 && a + 1 < len(flowerbed) {
				counter++
				flowerbed[a] = flowerbed[a] + 1
				break
			}

			if counter == n {
				return true
			}

			if !check(flowerbed) {
				return false
			}

			fmt.Println(counter, flowerbed)
		}
	}

}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}
