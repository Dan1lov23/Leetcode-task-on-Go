package main

import "fmt"

func getElementsSlice(matrix [][]int) (elementsSlice []int) {

	for _, slice := range matrix {
		for _, sliceElement := range slice {
			elementsSlice = append(elementsSlice, sliceElement)
		}
	}

	return

}

func reverseSlice(slice []int) (reverseSlice []int) {

	for i := len(slice) - 1; i >= 0; i-- {
		reverseSlice = append(reverseSlice, slice[i])
	}

	return

}

func rotate(matrix [][]int) {

	rotateMatrix := [][]int{}

	actualSliceLen := len(matrix[0])

	allElementsSlice := getElementsSlice(matrix)

	counter := 0
	startPosition := 0

	for counter < len(matrix) {

		subSlice, indexForDeletes, iterCounter := []int{}, []int{}, 0

		for i := startPosition; i < len(allElementsSlice); i++ {

			if iterCounter%actualSliceLen == 0 {
				subSlice = append(subSlice, allElementsSlice[i])
				indexForDeletes = append(indexForDeletes, i)
			}

			iterCounter++

		}

		rotateMatrix = append(rotateMatrix, reverseSlice(subSlice))
		counter++
		startPosition++
	}

	fmt.Println(rotateMatrix)

}

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
