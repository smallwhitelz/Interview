package main

import "fmt"

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	charMap := map[byte]int{}
	for i := 0; i < len(p); i++ {
		charMap[p[i]]++
	}
	for left, right := 0, 0; right < len(s); right++ {
		charMap[s[right]]--
		for charMap[s[right]] < 0 {
			charMap[s[left]]++
			left++
		}
		if right-left+1 == len(p) {
			ans = append(ans, left)
		}
	}
	return ans
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}
