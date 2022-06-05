package main

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question206 struct {
	para206
	ans206
}

// para 是参数
// one 代表第一个参数
type para206 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans206 struct {
	one []int
}

func Test_Problem206(t *testing.T) {

	qs := []question206{

		{
			para206{[]int{9, 3, 7}},
			ans206{[]int{6, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 206------------------------\n")

	for _, q := range qs {
		p1, p := q.ans206, q.para206
		a := structures.Ints2List(p1.one)
		b := structures.Ints2List(p.one)
		Add(a, b)
	}
}
