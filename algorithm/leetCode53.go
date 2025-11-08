package main

import (
	"fmt"
	"math"
)

// 最大子数组和，前缀和解法：长的前缀和减去短的前缀和可以算出中间差值
// [-2, 1, -3, 4, -1, 2, 1, -5, 4]
//
//	0   1   2  3  4   5  6  7   8
//
// [3,6]的和 = [0,6]的和-[0,2]的和
// 子数组和=cur-short
// 如何让子数组和最大，那就固定cur，找出最小的short就可以
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := math.MinInt
	curPrefixSum, minPrefixSum := 0, 0
	for _, num := range nums {
		curPrefixSum += num
		// 算出当前子数组最大和
		ans = max(ans, curPrefixSum-minPrefixSum)
		// 计算出和当前前缀和相比最小的short
		minPrefixSum = min(minPrefixSum, curPrefixSum)
	}
	return ans
}

// 动态规划
// 定义dp[i]是以nums[i]结尾的最大子数组和
// 就存在：当dp[i-1]>=0时 dp[i] = dp[i-1]+nums[i]
//
//	当dp[i-1]<0时，dp[i] = nums[i]
//
// 简单说就是前一个大于0，我们就加上，小于0，就不要，然后在dp数组中找到最大值
func maxSubArrayV1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] >= 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
