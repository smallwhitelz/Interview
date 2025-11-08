package main

import (
	"fmt"
	"slices"
)

// 合并区间

func mergeIntervals(intervals [][]int) [][]int {
	// 按照左端点从小到大排序
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	ans := make([][]int, 0)
	for _, arr := range intervals {
		if len(ans) > 0 && arr[0] <= ans[len(ans)-1][1] { // 可以合并
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], arr[1]) // 更新右端点最大值
		} else { // 不相交，无法合并
			ans = append(ans, arr) // 新的合并区间
		}
	}
	return ans
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(mergeIntervals(intervals))
}
