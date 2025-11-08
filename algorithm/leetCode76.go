package main

import "fmt"

// 最小覆盖子串

// 窗口状态：覆盖所有子串
// 如果left,right覆盖t中所有子串，就有(left, right+1...end)覆盖所有子串
// 如果left,right没有覆盖t所有子串，就有(left+1...right,right)没有覆盖
func minWindow(s string, t string) string {
	charCnt := map[byte]int{}
	for i := 0; i < len(t); i++ {
		charCnt[t[i]]++
	}
	totalCnt := 0 // 窗口中有效字符的数量
	ans_left, ans_right := 0, len(s)
	for left, right := 0, 0; right < len(s); right++ {
		// right加入窗口
		charCnt[s[right]]--
		// 表明该字符是有效字符
		if charCnt[s[right]] >= 0 {
			totalCnt++
		}
		for totalCnt == len(t) {
			// 根据情况收集答案
			if ans_right-ans_left > right-left {
				ans_left = left
				ans_right = right
			}
			// left移出窗口
			charCnt[s[left]]++
			if charCnt[s[left]] > 0 {
				totalCnt--
			}
			left++
		}
	}
	if ans_right == len(s) {
		return ""
	}
	return s[ans_left : ans_right+1]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
