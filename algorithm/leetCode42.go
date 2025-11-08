package main

// 接雨水

// 分而治之，动态规划
func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	preMax := make([]int, len(height))
	preMax[0] = height[0]
	for i := 1; i < len(height)-1; i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}
	sufMax := make([]int, len(height))
	sufMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}
	ans := 0
	for i := 1; i < len(height)-1; i++ {
		if height[i] >= preMax[i-1] || height[i] >= sufMax[i+1] {
			continue
		}
		ans += min(preMax[i-1], sufMax[i+1]) - height[i]
	}
	return ans
}

// 双指针
// lmax left左边最大值 当lmax < rmax时，left处的水可以确定，为0或者lmax-height[left]
// rmax right右边最大值 当lmax >= rmax时，right处的水可以确定，为0或者rmax-height[right]
// 和盛水最多的容器有点像，能接多少水和最小的一边相关
func trapV1(height []int) int {
	if len(height) < 2 {
		return 0
	}
	left, right := 1, len(height)-2
	lmax, rmax := height[0], height[len(height)-1]
	ans := 0
	for left <= right {
		if lmax < rmax {
			ans += max(0, lmax-height[left])
			lmax = max(lmax, height[left])
			left++
		} else {
			ans += max(0, rmax-height[right])
			rmax = max(rmax, height[right])
			right--
		}
	}
	return ans
}
