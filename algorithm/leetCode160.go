package main

// 相交链表

// 哈希表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	set := map[*ListNode]struct{}{}
	for headA != nil {
		set[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := set[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 双指针
// A走完了去走B的路，B走完了自己的再去走A的路，如果有相交的节点，两个一定会相遇；
// 没有相交的节点就会为nil
func getIntersectionNodeV1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}
