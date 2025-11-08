package main

import "fmt"

// 回文链表

func isPalindrome(head *ListNode) bool {
	stack := make([]int, 0)
	cur := head
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}
	for head != nil {
		if head.Val != stack[len(stack)-1] {
			return false
		}
		head = head.Next
		stack = stack[:len(stack)-1]
	}
	return true
}
func main() {

}
