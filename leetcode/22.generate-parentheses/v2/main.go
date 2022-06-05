package main

import "fmt"

var result []string

func main() {
	fmt.Println(generateParenthesis(3))
}

// 回溯算法（深度优先遍历）
// generateParenthesis 和递归函数分开写的版本
func generateParenthesis(n int) []string {
	generate(0, 0, n, "")
	return result
}

func generate(left, right, max int, currentStings string) {
	// 递归终结者
	if left == max && right == max { // left, right 都用完了
		// process result
		result = append(result, currentStings)
		return
	}

	// 处理当前层理逻辑; 往格子里放括号
	if left < max { // left 没用完就可以继续放 left
		s1 := currentStings + "("
		// 下探到下一层
		generate(left+1, right, max, s1)
	}
	if right < left { // 当使用的 right 数量 小于 left 数量 可以使用 right
		s2 := currentStings + ")"
		// 下探到下一层
		generate(left, right+1, max, s2)
	}

	// 清理当前层
}
