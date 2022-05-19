package main

import "fmt"

func main() {
	generateParenthesis(3)
}

func generateParenthesis(n int) []string {
	generate(0, 2*n, "")
	return nil
}

func generate(level int, max int, currentStings string) {
	// 递归终结者
	if level >= max {
		// process result
		fmt.Println(currentStings)
		return
	}

	// 处理当前层理逻辑; 往格子里放括号
	s1 := currentStings + "("
	s2 := currentStings + ")"

	// 下探到下一层
	generate(level+1, max, s1)
	generate(level+1, max, s2)
	// 清理当前层
}
