package main

import "fmt"

// 1. 两数之和

// 哈希表解法
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 || nums == nil {
		return nil
	}
	m := make(map[int]int)
	for index, num := range nums {
		sub := target - num
		if _, ok := m[sub]; !ok {
			m[num] = index // 相当于差值作为key，索引作为值，如果key存在，说明我们需要的差值找到了，那么直接返回索引即可
		} else {
			return []int{index, m[sub]}
		}
	}
	return nil
}

// 暴力枚举
func twoSumV1(nums []int, target int) []int {
	if len(nums) == 0 || nums == nil {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{3, 3}
	sum := twoSum(nums, 6)
	fmt.Println(sum)
}
