package main

import "sort"

// 169. 多数元素

func majorityElement(nums []int) int {
	n := len(nums)
	mp := make(map[int]int, n)
	for _, num := range nums {
		mp[num]++
		if mp[num] > n/2 {
			return num
		}
	}
	return -1
}

// 排序后中间的数一定是多数元素
func majorityElementV1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 摩尔投票法
// 时间复杂度 O(n)，空间复杂度 O(1)
// 该算法的核心思想是通过抵消不同元素的数量，最终剩下的元素即为多数元素。
// 因为多数元素的数量超过了数组长度的一半，所以它不可能被完全抵消掉。
func majorityElementV2(nums []int) int {
	// 记录候选人
	candidate := 0
	// 票数
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
