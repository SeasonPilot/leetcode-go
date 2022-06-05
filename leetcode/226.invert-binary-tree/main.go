package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetCode写法
// https://leetcode.cn/problems/invert-binary-tree/solution/shou-hua-tu-jie-san-chong-xie-fa-di-gui-liang-chon/
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left) // 翻转子树内部
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}

func invertTree(root *TreeNode) *TreeNode {
	// 1.terminator
	if root == nil {
		return nil
	}

	// 2.process logic in current level
	root.Left, root.Right = root.Right, root.Left

	// 3.drill down
	invertTree(root.Left)
	invertTree(root.Right)

	// 4.清理当前层
	// 无

	return root
}
