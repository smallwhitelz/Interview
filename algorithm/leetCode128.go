package main

// 128. 最长连续序列

// 用哈希表查找这个数前面一个数是否存在，即num-1在序列中是否存在。存在那这个数肯定不是开头，直接跳过。
//因此只需要对每个开头的数进行循环，直到这个序列不再连续，因此复杂度是O(n)。
//以题解中的序列举例:
//[100，4，200，1，3，4，2]
//去重后的哈希序列为：
//[100，4，200，1，3，2]
//按照上面逻辑进行判断：
//元素100是开头,因为没有99，且以100开头的序列长度为1
//元素4不是开头，因为有3存在，过，
//元素200是开头，因为没有199，且以200开头的序列长度为1
//元素1是开头，因为没有0，且以1开头的序列长度为4，因为依次累加，2，3，4都存在。
//元素3不是开头，因为2存在，过，
//元素2不是开头，因为1存在，过。
//完

func longestConsecutive(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	numSet := make(map[int]bool)
	for _, num := range nums {
		// 去重
		numSet[num] = true
	}
	longStreak := 0
	// key
	for num := range numSet {
		// 如果当前数字的前一位没在序列中，那么就是一个序列的开头
		if !numSet[num-1] {
			currentNum := num
			// 当前序列长度
			currentStreak := 1
			// 判断这个序列最长能有多长
			// 比如1是开头，那么看2在不在，3在不在，4在不在，在的话就继续加长度
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longStreak < currentStreak {
				longStreak = currentStreak
			}
		}
	}
	return longStreak
}

func main() {
	nums := []int{1, 0, 1, 2}
	longestConsecutive(nums)
}
