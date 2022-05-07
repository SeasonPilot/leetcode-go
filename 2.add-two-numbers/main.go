package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/add-two-numbers/
//https://leetcode-cn.com/problems/add-two-numbers/solution/jian-dan-yi-dong-javacpythonjs-pei-yang-y2w6g/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1} // 哑节点
	curr, carry := dummy, 0     // 当前遍历的节点  进位
	for l1 != nil || l2 != nil {
		x, y := 0, 0 // 链表循环完时用 0 占位
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		total := x + y + carry
		curr.Next = &ListNode{Val: total % 10}
		curr = curr.Next // 移动 curr 指针到下一个节点
		carry = total / 10

		if l1 != nil {
			l1 = l1.Next // 移动 l1 到下一个节点
		}
		if l2 != nil {
			l2 = l2.Next // 移动 l2 到下一个节点
		}
	}

	// 循环结束后，如果进位不为 0 则加入结果中
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
