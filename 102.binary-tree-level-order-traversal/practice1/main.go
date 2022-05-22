package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func levelOrder(root *TreeNode) [][]int {
	var ret [][]int // ret result
	if root == nil {
		return ret
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var child []int
		var nextQueue []*TreeNode

		for _, node := range queue {
			child = append(child, node.Val) // process(node)

			// generate_related_nodes(node)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}

		ret = append(ret, child)
		queue = nextQueue //queue.push(nodes)
	}
	return ret
}
