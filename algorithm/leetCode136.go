package main

// 136. 只出现一次的数字

// 位运算玩法
// a ^ a = 0
// a ^ 0 = a
// a ^ b ^ a = (a ^ a) ^ b = 0 ^ b = b
// 因此对所有数字进行异或运算，成对出现的数字会被抵消，最终剩下的就是只出现一次的数字
func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

// 哈希玩法
func singleNumberV1(nums []int) int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	minRes := nums[0]
	for val, _ := range mp {
		if mp[val] < mp[minRes] {
			minRes = val
		}
	}
	return minRes
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	println(singleNumber(nums))
}
