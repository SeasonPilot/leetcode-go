package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func largestValues(root *TreeNode) []int {
	var ret []int
	// fixme: 异常判断
	if root == nil {
		return ret
	}

	q := []*TreeNode{root}
	for q != nil {
		var nextQ []*TreeNode

		// pop
		var r int
		for i, node := range q {
			// process node
			if i == 0 {
				r = node.Val
			}
			r = max(r, node.Val)

			// add nextQ  generate_related_nodes(node)
			if node.Left != nil {
				nextQ = append(nextQ, node.Left)
			}
			if node.Right != nil {
				nextQ = append(nextQ, node.Right) // fixme: typo
			}
		}
		ret = append(ret, r)
		q = nextQ // fixme: 没有把 nextQ 赋值给 q
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
