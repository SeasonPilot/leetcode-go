package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var a, b = head, head
	// process logic in current level
	for i := 0; i < k; i++ {
		// terminator
		if b == nil {
			return head
		}
		b = b.Next
	}
	newList := reverse(a, b)
	// drill down
	a.Next = reverseKGroup(b, k)
	return newList
}

func reverse(head, tail *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
