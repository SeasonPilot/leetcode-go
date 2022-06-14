package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := new(ListNode)
	next := head

	for next != nil {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
	}
	//return next // fixme: return cur   tmp next 最后是nil的
	return cur
}
