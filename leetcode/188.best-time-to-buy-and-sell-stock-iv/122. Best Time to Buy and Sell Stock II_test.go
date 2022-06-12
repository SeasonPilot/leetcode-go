package main

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
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans122 struct {
	one int
}

func Test_Problem122(t *testing.T) {

	qs := []question122{

		{
			para122{[]int{2, 4, 1}, 2},
			ans122{2},
		},
		{
			para122{[]int{3, 2, 6, 5, 0, 3}, 2},
			ans122{7},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 122------------------------\n")

	for _, q := range qs {
		a, p := q.ans122, q.para122
		fmt.Printf("【结果是否正确】:%v \n", a.one == maxProfit(p.k, p.one))
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxProfit(p.k, p.one))
	}
	fmt.Printf("\n\n\n")
}
