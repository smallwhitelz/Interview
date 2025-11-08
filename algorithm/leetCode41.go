package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		m[num] = struct{}{}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return len(nums) + 1
}

func main() {
	nums := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(nums))
}
