package method1

import (
	"fmt"
	"testing"
)

type question122 struct {
	para122
	ans122
}

// para 是参数
// one 代表第一个参数
type para122 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans122 struct {
	one int
}

func Test_Problem122(t *testing.T) {

	qs := []question122{

		{
			para122{[]int{}},
			ans122{0},
		},

		{
			para122{[]int{3, 3, 5, 0, 0, 3, 1, 4}},
			ans122{6},
		},

		{
			para122{[]int{7, 6, 4, 3, 1}},
			ans122{0},
		},

		{
			para122{[]int{1, 2, 3, 4, 5}},
			ans122{4},
		},

		{
			para122{[]int{1}},
			ans122{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 122------------------------\n")

	for _, q := range qs {
		a, p := q.ans122, q.para122
		fmt.Printf("【结果是否正确】:%v \n", a.one == maxProfit(p.one))
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxProfit(p.one))
	}
	fmt.Printf("\n\n\n")
}
