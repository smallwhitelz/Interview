package main

import "fmt"

// 滑动窗口
// 窗口状态->存在重复字符
// 1、如果(left，right)满足状态(存在重复字符)则(left，right+1...end)也满足状态。(遍历时右移left指针)
// 2、如果(left，right)不满足状态(不存在重复字符)则(left+1...right，right)也不满足状态。(遍历时右移right指针)
func lengthOfLongestSubstring(s string) int {
	ans := 0
	charCnt := map[byte]int{}
	for left, right := 0, 0; right < len(s); right++ {
		// 第一次进入for循环是空子串，不满足状态
		// right加入窗口
		charCnt[s[right]]++
		// 满足状态的时候
		for charCnt[s[right]] > 1 {
			// 根据情况收集答案
			// left移除窗口
			charCnt[s[left]]--
			left++
		}
		// 不满足状态（根据情况收集答案）
		ans = max(ans, right-left+1)
	}
	return ans
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
