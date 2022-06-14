package main

// 方法一：迭代
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//prev := new(ListNode)// fixme: 使用new 初始化的节点的 val 为 0.
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next

		curr.Next = prev // 改变node指向  每次只改变一个指向
		prev = curr      // 移动指针
		curr = next
	}
	return prev
}
