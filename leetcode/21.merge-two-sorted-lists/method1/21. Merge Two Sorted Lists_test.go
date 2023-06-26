package method1

import (
	"fmt"
	"reflect"
	"testing"
)

type question21 struct {
	para21
	ans21
}

// para 是参数
// one 代表第一个参数
type para21 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans21 struct {
	one []int
}

func Test_Problem21(t *testing.T) {

	qs := []question21{

		{
			para21{[]int{}, []int{}},
			ans21{[]int{}},
		},

		{
			para21{[]int{1}, []int{1}},
			ans21{[]int{1, 1}},
		},

		{
			para21{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans21{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		{
			para21{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans21{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		{
			para21{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans21{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			para21{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans21{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			para21{[]int{2, 3, 4}, []int{4, 5, 6}},
			ans21{[]int{2, 3, 4, 4, 5, 6}},
		},

		{
			para21{[]int{1, 3, 8}, []int{1, 7}},
			ans21{[]int{1, 1, 3, 7, 8}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 21------------------------\n")

	for _, q := range qs {
		_, p := q.ans21, q.para21
		fmt.Printf("【input】:%v       【output】:%v\n", p, List2Ints(mergeTwoLists(Ints2List(p.one), Ints2List(p.another))))
		fmt.Printf("【结果是否正确】:%v \n", reflect.DeepEqual(q.ans21.one, List2Ints(mergeTwoLists(Ints2List(p.one), Ints2List(p.another)))))
	}
	fmt.Printf("\n\n\n")
}

// List2Ints convert List to []int
func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}
