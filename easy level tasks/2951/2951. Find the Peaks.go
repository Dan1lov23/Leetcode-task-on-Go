package main

import "fmt"

func findPeaks(mountain []int) []int {

	result := []int{}

	for i := 0; i < len(mountain); i++ {
		if i - 1 >= 0 && i + 1 < len(mountain) && mountain[i] > mountain[i + 1] && mountain[i] > mountain[i - 1] {
			result = append(result, i)
		}
	}

	return result

}

func main() {
	fmt.Println(findPeaks([]int{1, 4, 3, 8, 5}))
}
