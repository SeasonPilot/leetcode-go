package main

import "fmt"

func main() {
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	fmt.Println(pairs[')'])
	// map的key 不存在时返回的值是 0
	fmt.Println(pairs['['])

	// map的key 不存在时返回的值是 0，结果是 false
	v, ok := pairs['[']
	fmt.Println(v, ok)
}
