package main

import "fmt"

// 和为K的子数组

func subarraySum(nums []int, k int) int {
	// key：前缀和的值，value：前缀和取值为key的子数组个数
	preSumCnt := map[int]int{}
	// 表示子数组为空，和为0，个数为1
	preSumCnt[0] = 1
	longerPreSum := 0
	ans := 0
	for _, val := range nums { // 统计以每个位置结尾的满足要求的子数组
		longerPreSum += val
		shortPreSum := longerPreSum - k
		_, ok := preSumCnt[shortPreSum]
		if ok {
			ans += preSumCnt[shortPreSum]
		} else {
			ans += 0
		}
		preSumCnt[longerPreSum]++
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	k := 3
	fmt.Println(subarraySum(nums, k))
}
