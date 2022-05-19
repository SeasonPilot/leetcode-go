package main

import "fmt"

// test rune byte
func main() {
	a := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'}}
	num := a[0][0]
	fmt.Printf("%T\n", num) // byte
	fmt.Printf("%q\n", num)
	fmt.Printf("%c\n", num)
	fmt.Printf("%v\n", num)

	fmt.Printf("%T\n", '1') // rune
	fmt.Printf("%q\n", '1')
	fmt.Printf("%c\n", '1')
	fmt.Printf("%v\n", '1')

	fmt.Println(rune(num))

	fmt.Printf("%T\n", 1)
	fmt.Println(rune(1))

	digit := num - '1'
	fmt.Println(digit)
}
