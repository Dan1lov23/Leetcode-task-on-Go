package main

import "fmt"

func makeFancyString(s string) string {
    
	newString := ""

	for i := 0; i < len(s); i++ {
		if i + 1 < len(s) && i + 2 < len(s) && s[i] == s[i + 1] && s[i] == s[i + 2] {
			//
		} else {
			newString += string(s[i])
		}
	}

	return newString

}

func main() {
	fmt.Println(makeFancyString("aaabaaaa"))
}
