package main

// 2. 两数相加

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{Val: -1}
	cur := pre
	carry := 0
	for l1 != nil || l2 != nil {
		x := 0
		if l1 != nil {
			x = l1.Val
		}
		y := 0
		if l2 != nil {
			y = l2.Val
		}
		sum := x + y + carry
		// 计算进位
		carry = sum / 10
		// 计算当前值，也就是如果10以上取余数就是当前位置应该放的值
		sum = sum % 10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry == 1 {
		cur.Next = &ListNode{Val: carry}
	}
	return pre.Next
}
