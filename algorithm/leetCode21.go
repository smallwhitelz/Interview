package main

// 合并两个有序链表

// 递归
func mergeTwoListsV1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoListsV1(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsV1(list1, list2.Next)
		return list2
	}
}

// 迭代
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{Val: -1}
	pre := preHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			preHead.Next = list1
			list1 = list1.Next
		} else {
			preHead.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if pre.Next == list1 && list1 == nil {
		pre.Next = list2
	} else {
		pre.Next = list1
	}
	return preHead.Next
}
