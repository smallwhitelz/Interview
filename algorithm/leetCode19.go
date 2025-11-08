package main

import "fmt"

// 删除链表的倒数第N个节点

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	cur = dummyHead
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummyHead.Next
}

// 栈玩法
func removeNthFromEndV1(head *ListNode, n int) *ListNode {
	stack := []*ListNode{}
	dummyHead := &ListNode{Val: -1, Next: head}
	cur := dummyHead
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}
	for i := 0; i < n; i++ {
		stack = stack[:len(stack)-1]
	}
	pre := stack[len(stack)-1]
	pre.Next = pre.Next.Next
	return dummyHead.Next
}

func main() {
	list := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}
	fmt.Println(removeNthFromEnd(list, 2))
}
