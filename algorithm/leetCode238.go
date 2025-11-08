package main

// 除自身以外数组的乘积

// 假如nums为[1,2,3,4]，那么answer的值分别为[(2,3,4),(1,3,4),(1,2,4),(1,2,3)]
// 如果吧i当前值相乘的时候看做是1那么就有如下样式
//
//	1, 2, 3, 4
//	1, 1, 3, 4
//	1, 2, 1, 4
//	1, 2, 3, 1
//
// 他的对角线1将他们分割成了两个三角形，对于answer的元素，
// 我们可以先计算一个三角形每行的乘积，然后再去计算另外一个三角形每行的乘积，
// 然后各行相乘，就是answer每个对应的元素
func productExceptSelf(nums []int) []int {
	// 前缀积
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = dp[i-1] * nums[i-1]
	}
	// 后缀积
	pd := make([]int, len(nums))
	pd[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		pd[i] = pd[i+1] * nums[i+1]
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = dp[i] * pd[i]
	}
	return ans
}
