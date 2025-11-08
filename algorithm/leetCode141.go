package main

// 环形链表

// 哈希表
func hasCycle(head *ListNode) bool {
	set := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := set[head]; ok {
			return true
		}
		set[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 快慢指针
func hasCycleV1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
