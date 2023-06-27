package main

//方法一：栈
//遇到左括号就入栈，遇到右括号就去栈中寻找最近的左括号，看是否匹配。

//我们遍历给定的字符串 s。当我们遇到一个左括号时，我们会期望在后续的遍历中，有一个相同类型的右括号将其闭合。
//由于 后遇到的左括号要先闭合，因此我们可以将这个左括号放入栈顶。
//
//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/valid-parentheses/solution/you-xiao-de-gua-hao-by-leetcode-solution/
//来源：力扣（LeetCode）

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
		if pairs[r] == 0 { //当 r 是左括号时
			stack = append(stack, r) //入栈
		} else { //当 r 是右括号时
			//如果不是相同类型的括号，或者栈中并没有左括号，那么字符串 s 无效，返回 False。
			//if stack[len(stack)-1] != pairs[r] || len(stack) == 0 { //fixme: panic: runtime error: index out of range [-1] ；栈为空时 stack[len(stack)-1] 切片下标越界，panic；所有要将 stack[len(stack)-1] 放在后面
			if len(stack) == 0 || stack[len(stack)-1] != pairs[r] { //右括号，栈为空 或者 从栈顶弹出左括号和右括号不匹配
				return false
			} else {
				stack = stack[:len(stack)-1] //括号匹配,弹栈
			}
		}
	}

	//在遍历结束后，如果栈中没有左括号，说明我们将字符串 s 中的所有左括号闭合，返回 True，否则返回 False。
	return len(stack) == 0
}
