package main

import "fmt"

func main() {
	s := "A"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%T\n", s[i]) // fori 循环字符串 uint8
	}

	for _, x := range s {
		fmt.Printf("%T\n", x) // forr 循环  int32
	}

	fmt.Printf("%T\n", s[0]) // 通过index 访问字符串得到的类型是 uint8
}
