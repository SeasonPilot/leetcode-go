package practice2

import (
	"fmt"
	"reflect"
	"testing"
)

type question25 struct {
	para25
	ans25
}

// para 是参数
// one 代表第一个参数
type para25 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans25 struct {
	one []int
}

func Test_Problem25(t *testing.T) {

	qs := []question25{

		{
			para25{
				[]int{1, 2, 3, 4, 5},
				3,
			},
			ans25{[]int{3, 2, 1, 4, 5}},
		},

		{
			para25{
				[]int{1, 2, 3, 4, 5},
				1,
			},
			ans25{[]int{1, 2, 3, 4, 5}},
		},
		{
			para25{
				[]int{1, 2, 3, 4, 5},
				2,
			},
			ans25{[]int{2, 1, 4, 3, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 25------------------------\n")

	for _, q := range qs {
		_, p := q.ans25, q.para25
		fmt.Printf("【input】:%v       【output】:%v\n", p, List2Ints(reverseKGroup(Ints2List(p.one), p.two)))
		fmt.Printf("【结果是否正确】:%v \n", reflect.DeepEqual(q.ans25.one, List2Ints(reverseKGroup(Ints2List(p.one), p.two))))
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
