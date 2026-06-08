package main

import "fmt"

func majorityElement(nums []int) []int {
    n := len(nums)
    threshold := n / 3
    freq := make(map[int]int)

    for _, num := range nums {
        freq[num]++
    }

    result := []int{}
    for num, count := range freq {
        if count > threshold {
            result = append(result, num)
        }
    }

    return result
}

func main() {
    fmt.Println(majorityElement([]int{1}))
}