package main

import "fmt"

func main() {

	row := "a........."        // golang的字符串本质上是字节slices
	fmt.Printf("%T\n", row[0]) // uint8 即 字节

	row[0] = []byte(".")[0]
	row[0] = byte('.')
	row[0] = '.'
	row[0] = "."

	fmt.Printf("%T\n", []byte(".")[0]) // uint8
	fmt.Printf("%T\n", byte('.'))      // uint8
	fmt.Printf("%T\n", '.')            // int32   码点
	fmt.Printf("%T\n", ".")            // string

}
