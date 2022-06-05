package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

// 回溯算法（深度优先遍历）
// 递归函数在 generateParenthesis 里面的版本
func generateParenthesis(n int) []string {
	var result []string

	var generate func(left, right, max int, currentStings string)
	generate = func(left, right, max int, currentStings string) {
		// 递归终结者
		if left == max && right == max {
			// process result
			result = append(result, currentStings)
			return
		}

		// 处理当前层理逻辑; 往格子里放括号
		if left < max {
			s1 := currentStings + "("
			// 下探到下一层
			generate(left+1, right, max, s1)
		}
		if right < left {
			s2 := currentStings + ")"
			// 下探到下一层
			generate(left, right+1, max, s2)
		}

		// 清理当前层
	}

	generate(0, 0, n, "")
	return result
}
