package main

// 27. 移除元素
// 拷贝覆盖
// 主要思路是遍历数组 nums，每次取出的数字变量为 num，同时设置一个下标 ans
// 在遍历过程中如果出现数字与需要移除的值不相同时，则进行拷贝覆盖 nums[ans] = num，ans 自增 1
// 如果相同的时候，则跳过该数字不进行拷贝覆盖，最后 ans 即为新的数组长度
// 这种思路在移除元素较多时更适合使用，最极端的情况是全部元素都需要移除，遍历一遍结束即可
// 时间复杂度：O(n)，空间复杂度：O(1)
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// 交换移除
// 主要思路是遍历数组 nums，遍历指针为 i，总长度为 ans
// 在遍历过程中如果出现数字与需要移除的值不相同时，则 i 自增 1 ，继续下一次遍历
// 如果相同的时候，则将 nums[i]与nums[ans-1] 交换，即当前数字和数组最后一个数字进行交换，交换后就少了一个元素，故而 ans 自减 1
// 这种思路在移除元素较少时更适合使用，最极端的情况是没有元素需要移除，遍历一遍结束即可
// 时间复杂度：O(n)，空间复杂度：O(1)
func removeElementV1(nums []int, val int) int {
	ans := len(nums)
	for i := 0; i < ans; {
		if nums[i] == val {
			nums[i] = nums[ans-1]
			ans--
		} else {
			i++
		}
	}
	return ans
}

func main() {
	nums := []int{3, 2, 2, 3}
	removeElement(nums, 3)
}
