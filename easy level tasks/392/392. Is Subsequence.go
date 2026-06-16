package main

import (
	"fmt"
	"strings"
)

func includeCheck(slice []string, element string) bool {

	for _, stringElement := range slice {
		if element == stringElement {
			return true
		}
	}

	return false

}

func getSliceFromString(s string) (slice []string) {

	for i := 0; i < len(s); i++ {
		slice = append(slice, string(s[i]))
	}

	return

}

func isSubsequence(s string, t string) bool {
	
	if strings.Contains(t, s) {
		return true
	}

	sSlice, tSlice, checkString := getSliceFromString(s), getSliceFromString(t), ""

	for _, stringElement := range tSlice {
		if includeCheck(sSlice, stringElement) {
			checkString += stringElement
		}
	}

	return s == checkString

}

func main() {
	fmt.Println(isSubsequence("acb", "ahbgdc"))
}
