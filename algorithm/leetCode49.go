package main

import (
	"slices"
)

// 49. 字母异位词分组
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
// 这里输出可以按照任何循序
// 所以有个思路：能互相异位的词他的字母和数量一定是相同的

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 || len(strs) == 1 || strs == nil {
		return [][]string{strs}
	}
	m := make(map[string][]string)
	for _, str := range strs {
		// 先将字符串转换为字节切片
		sorted := []byte(str)
		// 对每一个字节切片进行排序
		slices.Sort(sorted)
		// 排序后再转为字符串放入到map中
		sortedS := string(sorted)
		m[sortedS] = append(m[sortedS], str)
	}
	ans := make([][]string, 0, len(m))
	for _, val := range m {
		ans = append(ans, val)
	}
	return ans
}

// 计数法
// 思路：字母只有26个，并且这里都是小写，那么相同的字母一定值是相同的，所以可以在同一个位置计数
func groupAnagramsV1(strs []string) [][]string {
	if len(strs) == 0 || len(strs) == 1 || strs == nil {
		return [][]string{strs}
	}
	mp := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, val := range mp {
		ans = append(ans, val)
	}
	return ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagramsV1(strs)
}
