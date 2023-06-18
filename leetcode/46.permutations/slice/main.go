package main

import "fmt"

func main() {
	var used []int // nil 的 slice 不能访问索引，不能赋值索引
	temp := make([]int, 10)

	fmt.Printf("%v,%v\n", used, len(used))
	//fmt.Println(used[0]) // nil 的 slice 不能访问索引，不能赋值索引
	//used[0] = 1          // nil 的 slice 不能访问索引，不能赋值索引
	used = temp // nil 的 slice 可以赋值
	fmt.Println(used[0])
	fmt.Println(used)
}
