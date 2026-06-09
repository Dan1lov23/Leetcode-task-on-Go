package main

import "fmt"

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }

    maxRob := make([]int, len(nums))
    maxRob[0] = nums[0]
    maxRob[1] = max(nums[0], nums[1])

    for i := 2; i < len(nums); i++ {
        maxRob[i] = max(maxRob[i-1], maxRob[i-2]+nums[i])
    }

    return maxRob[len(nums)-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(rob([]int{2, 3, 2}))
    fmt.Println(rob([]int{1, 2, 3, 1}))
    fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}