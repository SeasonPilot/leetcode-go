package main

import "fmt"

func main() {
	s := "()[]{}"
	s1 := "(("
	fmt.Println(isValid(s))
	fmt.Println(isValid(s1))
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []rune
	for _, r := range s {
		// map的key 不存在时返回的值是 0；即为左括号
		if pairs[r] == 0 {
			stack = append(stack, r)
			//如果不是相同的类型，或者栈中并没有左括号，那么字符串 s 无效，返回 False。
		} else if stack[len(stack)-1] != pairs[r] || len(stack) == 0 {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	//在遍历结束后，如果栈中没有左括号，说明我们将字符串 s 中的所有左括号闭合，返回 True，否则返回 False。
	return len(stack) == 0
}
