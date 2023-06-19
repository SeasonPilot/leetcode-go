package practice2

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		//if b.Next == nil {//fixme: 应该是	if b == nil {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b) //fixme:返回值是新链表的头

	a.Next = reverseKGroup(b, k) //反转链表后a变为 tail

	return newHead
}

func reverse(head *ListNode, tail *ListNode) *ListNode {
	var (
		prev = &ListNode{}
		curr = head
	)

	for curr != tail {
		//next := head.Next //fixme: curr.next
		next := curr.Next

		//head.Next = prev //fixme: curr.next
		curr.Next = prev
		prev = curr //先移动 prev,如果先移动 curr,prev 就无法移动了，prev 就变成和 curr 一样的值了。 curr 的值就变了，prev 依赖 curr 的值。
		curr = next
	}

	return prev
}
