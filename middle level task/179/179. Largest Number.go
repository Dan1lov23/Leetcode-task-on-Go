package main

import (
	"sort"
	"strconv"
	"strings"
	"fmt"
)

func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, v := range nums {
		ss[i] = strconv.Itoa(v)
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] > ss[j]+ss[i]
	})

	if len(ss) > 0 && ss[0] == "0" {
		return "0"
	}

	var b strings.Builder
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}

func main() {
	fmt.Println(largestNumber([]int{3,30,34,5,9}))
}
