package main

import (
	"fmt"
	"reflect"
	"testing"
)

type question37 struct {
	para37
	ans37
}

// para 是参数
// one 代表第一个参数
type para37 struct {
	s [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans37 struct {
	s [][]byte
}

func Test_Problem37(t *testing.T) {

	qs := []question37{

		{
			para37{[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}},
			ans37{[][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 37------------------------\n")

	for _, q := range qs {
		a, p := q.ans37, q.para37
		fmt.Printf("【input】:%c \n\n", p.s)
		solveSudoku(p.s)
		fmt.Printf("【output】:%c \n\n", p.s)
		fmt.Printf("【是否正确】:%v \n\n", reflect.DeepEqual(p.s, a.s))
	}
	fmt.Printf("\n\n\n")
}
