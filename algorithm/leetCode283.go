package main

import "fmt"

// 移动非零元素
func moveZeroes(nums []int) {
	cur, tail := 0, 0
	for cur < len(nums) {
		if nums[cur] == 0 {
			cur++
		} else {
			nums[cur], nums[tail] = nums[tail], nums[cur]
			cur++
			tail++
		}
	}
}

// 第二种玩法，把非0的直接前移，后面剩余部分用0补充
func moveZeroesV1(nums []int) {
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[cur] = nums[i]
			cur++
		}
	}
	for i := cur; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
