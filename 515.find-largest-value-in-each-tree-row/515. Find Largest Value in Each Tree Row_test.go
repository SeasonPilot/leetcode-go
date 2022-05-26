package main

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question515 struct {
	para515
	ans515
}

// para 是参数
// one 代表第一个参数
type para515 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans515 struct {
	one []int
}

func Test_Problem515(t *testing.T) {

	qs := []question515{

		{
			para515{[]int{}},
			ans515{[]int{}},
		},

		{
			para515{[]int{1}},
			ans515{[]int{1}},
		},

		{
			para515{[]int{1, 3, 2, 5, 3, structures.NULL, 9}},
			ans515{[]int{1, 3, 9}},
		},

		{
			para515{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans515{[]int{3, 20, 15}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 515------------------------\n")

	for _, q := range qs {
		_, p := q.ans515, q.para515
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", largestValues(root))
	}
	fmt.Printf("\n\n\n")
}

// NULL 方便添加测试数据
var NULL = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
