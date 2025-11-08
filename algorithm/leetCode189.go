package main

import (
	"fmt"
	"slices"
)

// 轮转数组

func rotate(nums []int, k int) {
	k %= len(nums) // 轮转 k 次等于轮转 k % n 次
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
