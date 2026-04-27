package main

func findGCD(num1 int, num2 int) (maxDev int) {

	maxNum := num1
	minNum := num2

	if num1 < num2 {
		maxNum = num2
		minNum = num1
	}

	for a := 1; a < maxNum; a++ {
		if maxNum % a == 0 && minNum % a == 0 {
			maxDev = a
		}
	}
	
	return

}

func insertGreatestCommonDivisors(arr []int) []int {

	var newArr []int

	for i := 0; i < len(arr); i++ {
		newArr = append(newArr, arr[i])
		if i + 1 < len(arr) {
			newArr = append(newArr, findGCD(arr[i], arr[i + 1]))
		}
	}

	return newArr

}

func main() {
	insertGreatestCommonDivisors([]int{7})
}
