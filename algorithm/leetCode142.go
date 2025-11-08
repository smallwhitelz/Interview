package main

// 环形链表2

// 哈希表
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	set := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := set[head]; ok {
			return head
		}
		set[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 快慢指针

func detectCycleV1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

func main() {

}
