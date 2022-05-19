package main

// 假设链表中每一个节点的值都在 0 - 9 之间，那么链表整体就可以代表一个整数。
// 给定两个这种链表，请生成代表两个整数相加值的结果链表。
// 例如：链表 1 为 9->3->7，链表 2 为 6->3，最后生成新的结果链表为 1->0->0->0。
// 使用纯链表操作，不能使用： 堆栈，数组切片，转为数字（字符串）再加

type Node struct {
	Val  int
	Next *Node
}

func Add(l1, l2 *Node) *Node {
	nl1 := reverse(l1)
	nl2 := reverse(l2)

	var dummy *Node //
	carry := 0
	curr := dummy //

	for nl1 != nil || nl2 != nil {
		x, y := 0, 0
		if nl1 != nil {
			x = nl1.Val
		}
		if nl2 != nil {
			x = nl2.Val
		}

		sum := x + y + carry
		curr.Next = &Node{Val: sum % 10}
		carry = sum / 10

		curr = curr.Next //

		// 错误：忘记移动 nl1 nl2
	}
	if carry != 0 {
		curr.Next = &Node{Val: carry}
	}

	result := reverse(dummy.Next)
	return result
}

func reverse(l *Node) *Node {
	pre := &Node{}
	curr := l

	for curr != nil {
		next := l.Next // 错误：这里应该是 curr
		l.Next = pre   // 错误：这里应该是 curr
		pre = curr
		curr = next
	}
	return curr // 错误：这里应该返回 pre
}
