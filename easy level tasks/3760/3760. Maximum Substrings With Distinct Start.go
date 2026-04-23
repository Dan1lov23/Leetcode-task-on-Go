package main

func checkStringInclude(s string, element string) (flag bool) {

	for _, symbol := range s {
		if string(symbol) == element {
			return true
		}
	}

	return false

}

func maxDistinct(s string) int {
    maxCounter := 0

	maxDiStr := ""

	for _, symbol := range s {
		if !checkStringInclude(maxDiStr, string(symbol)) {
			maxDiStr += string(symbol)
			maxCounter = len(maxDiStr)
		}
	}

	return maxCounter
}

func main() {
	maxDistinct("abab")
}
