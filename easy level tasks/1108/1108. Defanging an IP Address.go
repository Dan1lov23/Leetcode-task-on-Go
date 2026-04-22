package main

import "fmt"

func defangIPaddr(address string) string {

	ipAddress := ""

	for _, stringElement := range address {
		if string(stringElement) == "." {
			ipAddress += "[.]"
		} else {
			ipAddress += string(stringElement)
		}
	}

	return ipAddress

}

func main() {

	fmt.Println(defangIPaddr("1.1.1.1"))

}
