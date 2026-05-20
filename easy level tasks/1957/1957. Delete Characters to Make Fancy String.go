package main

import "fmt"
import "strings"

func makeFancyString(s string) string {
    var sb strings.Builder
    sb.Grow(len(s))

    count := 0
    for i := 0; i < len(s); i++ {
        if i > 0 && s[i] == s[i-1] {
            count++
        } else {
            count = 0
        }

        if count < 2 {
            sb.WriteByte(s[i])
        }
    }

    return sb.String()
}


func main() {
	fmt.Println(makeFancyString("aaabaaaa"))
}
