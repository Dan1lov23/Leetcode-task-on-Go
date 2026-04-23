package main

import "fmt"

func findDegrees(matrix [][]int) []int {

	var unitedMatrix [] int

	for i := 0; i < len(matrix[0]); i++ {
		unitedMatrix = append(unitedMatrix, 0)
	}

	for _, arr := range matrix {
		for i := 0; i <= len(arr) - 1; i++ {
			unitedMatrix[i] += arr[i]
		}
	}


	return unitedMatrix

}

func main() {
	fmt.Println(findDegrees([][]int{{0,1,0},{1,0,0},{0,0,0}}))
}
