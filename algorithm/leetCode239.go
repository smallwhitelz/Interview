package main

import "fmt"

// 滑动窗口最大值

// 单调队列
// 队首元素为当前窗口最大值
// 元素加入窗口时：
// 1.如果队列为空，直接入队
// 2.如果元素<队尾元素，直接入队
// 3.如果元素>=队尾元素，队尾元素弹出
// 元素移出窗口时：
// 如果队首元素是移出的元素，队首元素弹出
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	deque := make([]int, 0, len(nums))
	// 将窗口扩大到K
	for i := 0; i < k; i++ {
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1] // pop_back
		}
		deque = append(deque, i)
	}
	// 记录第一个窗口的最大值
	res = append(res, nums[deque[0]])
	// 将窗口向右移动，增加一个元素，移出一个元素
	for i := k; i < len(nums); i++ {
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1] // pop_back
		}
		deque = append(deque, i)
		removeIndex := i - k
		if deque[0] == removeIndex {
			// pop_front
			deque = deque[1:]
		}
		// 当前滑动窗口最大值为队首
		res = append(res, nums[deque[0]])
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	nums = nums[:len(nums)-1]
	fmt.Println(nums)
	maxSlidingWindow(nums, k)
}
