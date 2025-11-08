package main

import (
	"fmt"
	"sort"
)

// 力扣11题：三数之和

// 1. 将第一个数固定，然后后面两个数相加就是0-第一个数的值
// 2. 数组排序，使用双指针求解两数之和
// 3. 移动指针过程中进行去重
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		first := nums[i]
		if first > 0 {
			// 第一个数大于0，直接结束，因为排序后三数之和等于0，第一个都大于0，后面相加肯定会更大
			break
		}
		target := 0 - first
		left, right := i+1, len(nums)-1
		for left < right {
			twoSum := nums[left] + nums[right]
			if twoSum < target {
				left++
			} else if twoSum > target {
				right--
			} else {
				second := nums[left]
				third := nums[right]
				res = append(res, []int{first, second, third})
				// 去重，左边的数还相等的话没必要在计算，因为结果是不重复的三元组
				for left < right && nums[left] == second {
					left++
				}
				for left < right && nums[right] == third {
					right--
				}
			}
			for i+1 < len(nums) && nums[i+1] == first {
				i++
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Ints(nums)
	fmt.Println(nums)
}
