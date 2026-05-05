package main

import "fmt"

func unitedStringArr(arr []string) (unStr string) {

	for _, stringElement := range arr {
		unStr += stringElement
	}

	return

}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    
	return unitedStringArr(word1) == unitedStringArr(word2)

}

func main() {
	fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
}
