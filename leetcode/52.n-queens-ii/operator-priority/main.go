package main

import "fmt"

func main() {
	a := 1<<3 - 1         // 左移 优先级大于 减号
	fmt.Printf("%v\n", a) // 7

	fmt.Printf("%v\n", 1<<3-1) //

	fmt.Printf("%v\n", 1&(1<<4-1)) // 1  //& 优先级大于 -
	fmt.Printf("%v\n", 1&1<<4-1)   // 15  从左至右依次计算  先 & << -
}
