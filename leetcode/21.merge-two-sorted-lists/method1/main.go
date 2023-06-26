package method1

type ListNode struct {
	Val  int
	Next *ListNode
}

//方法一：迭代
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	newHead := dummy

	for list1 != nil && list2 != nil {

		if list1.Val >= list2.Val {
			//newHead.Val = list2.Val //fixme:不是把 list2 的值赋值给新链表，而是要把 list2 节点 连接到新链表
			newHead.Next = list2
			list2 = list2.Next //fixme:取值的链表才需要移动
		} else {
			newHead.Next = list1
			list1 = list1.Next //fixme:先比较值再进行节点移动
		}
		newHead = newHead.Next //fixme:如何移动链表？？？ 没移动的原因是没有把 list2 节点 连接到新链表，而只是把 list2 的值赋值给新链表
	}

	//for list1 != nil {
	//	newHead.Next = list1
	//	list1 = nil //fixme:把剩余的list1全部添加到newHead后，需要将list1置空；或者判断条件改成 if list1 != nil {}
	//}

	if list1 != nil {
		newHead.Next = list1
	}
	if list2 != nil {
		newHead.Next = list2
	}
	return dummy.Next //fixme: return dummy.Next
}

//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	dummyHead := &ListNode{}
//	p := dummyHead
//	for list1 != nil && list2 != nil {
//		if list1.Val >= list2.Val {
//			p.Next = list2
//			list2 = list2.Next
//		} else {
//			p.Next = list1
//			list1 = list1.Next
//		}
//		p = p.Next
//	}
//	if list1 != nil {
//		p.Next = list1
//	}
//	if list2 != nil {
//		p.Next = list2
//	}
//	return dummyHead.Next
//}

//作者：qian2
//链接：https://leetcode.cn/problems/merge-two-sorted-lists/solution/go-die-dai-by-qian2-9jp8/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
