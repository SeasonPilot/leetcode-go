package main

type ListNode struct {
	data int
	Next *ListNode
}

//https://leetcode-cn.com/problems/reverse-linked-list/solution/shi-pin-jiang-jie-die-dai-he-di-gui-hen-hswxy/
// 迭代
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var curr = head

	for curr != nil {
		next := curr.Next // 先存储下一个节点
		curr.Next = prev  // 将指针反转，指向 prev
		prev = curr       // 移动 prev 指针到 curr 位置
		curr = next       // 移动 curr 指针 next 位置
	}
	return prev
}
