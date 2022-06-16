package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归解法
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间 [a, b) 包含 k 个待反转元素
	var a, b = head, head

	// process logic in current level
	for i := 0; i < k; i++ {
		// terminator  // 不足 k 个，不需要反转，base case
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 反转前 k 个元素
	newHead := reverse(a, b) // 入参(1，3); 返回 2,1 。 翻转链表后 a 为变成了 tail
	// drill down // 递归反转后续链表并连接起来
	a.Next = reverseKGroup(b, k) // 入参(3，2)。 翻转后的尾部 指向后面 翻转后的链表
	// reverse
	return newHead
}

/* 反转区间 [a, b) 的元素，注意是左闭右开 */
func reverse(head, tail *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	// for 终止的条件改一下就行了
	for curr != tail {
		nex := curr.Next // 1. 存储
		curr.Next = prev // 2. 反转 Next 指针
		prev = curr      // 3. 移动指针
		curr = nex
	}
	return prev // fixme: return prev
}
