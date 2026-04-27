package main

import "fmt"

func findMax(arr []int) (max int) {

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return

}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	result := []bool{}

	for _, num := range candies {
		fmt.Println(findMax(candies))
		if num+extraCandies >= findMax(candies) {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result

}

func main() {
	fmt.Println(kidsWithCandies([]int{1, 3, 9}, 4))
}
