package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Log2(1))
	fmt.Println(math.Log2(2))
	fmt.Println(math.Log2(5))
	fmt.Println(math.Log2(7))
	fmt.Println(math.Log2(12))
	fmt.Println(math.Floor(math.Log2(7)))
	fmt.Println(math.Floor(math.Log2(7)) + 1)

	fmt.Println()

	fmt.Println(2/2 - 1)
	fmt.Println(5/2 - 1)
	fmt.Println(7/2 - 1)

}
