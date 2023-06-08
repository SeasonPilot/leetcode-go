package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var curr = head

	for curr != nil {
		next := curr.Next

		curr.Next = prev //移动指针

		prev = curr
		curr = next // fixme: 要先移动 prev 再移动 curr；因为 curr 最后的值是 nil
	}

	return prev
}
