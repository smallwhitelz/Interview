package main

import (
	"fmt"
)

// 暴力枚举，时间复杂度On²，会超时
func maxAreaV2(height []int) int {
	area := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			//	高度由两条线之间最短决定
			var h int
			if height[i] < height[j] {
				h = height[i]
			} else {
				h = height[j]
			}
			area = max(area, h*(j-i))
		}
	}
	return area
}

// 双指针玩法：
// 如果h[left]<h[right]，那么left,left+1...left,right-1都小于left,right的面积
// 如果h[left]>h[right]，那么left+1,right...right-1,right都小于left,right的面积
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		curArea := (right - left) * min(height[left], height[right])
		area = max(area, curArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
