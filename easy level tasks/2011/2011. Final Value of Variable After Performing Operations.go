package main

import "fmt"

func finalValueAfterOperations(operations []string) int {

	counter := 0

	for _, word := range operations {
		for _, strElement := range word {
			if string(strElement) == "+" {
				counter++
				break
			} else if string(strElement) == "-" {
				counter--
				break
			}
		}
	}

	return counter

}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"++X","X++","X++"}))
}
