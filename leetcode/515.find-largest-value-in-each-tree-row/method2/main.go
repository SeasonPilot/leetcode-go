package main

//链接：https://leetcode.cn/problems/find-largest-value-in-each-tree-row/solution/bfs-by-linbingyuan-2/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res = []int{}
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var length = len(queue)
		var max = -1 << 63
		for i := 0; i < length; i++ {
			if max < queue[i].Val {
				max = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, max)
		queue = queue[length:]
	}
	return res
}
