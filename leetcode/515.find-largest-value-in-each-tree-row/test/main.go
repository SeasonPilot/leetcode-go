package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var nextQ []*TreeNode
	nextQ = append(nextQ, nil)

	fmt.Println(len(nextQ))   // 1
	fmt.Println(nextQ == nil) // false
	fmt.Printf("%#v\n", nextQ)
	fmt.Printf("%+v\n", nextQ)
}
