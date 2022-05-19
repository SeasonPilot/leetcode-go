package main

import "fmt"

func main() {
	var root = new(TreeNode)
	var RightNode = new(TreeNode)
	var LeftNode = new(TreeNode)

	root.Val = 1
	root.Right = RightNode

	RightNode.Val = 2
	RightNode.Left = LeftNode
	LeftNode.Val = 3

	fmt.Println(inorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}

	inorder(root)
	return
}
