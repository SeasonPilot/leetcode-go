package main

import (
	"fmt"
	"sync/atomic"
)

// 取模，有 2 元钱，去买 128 块钱的吹风机，
func main() {
	fmt.Println(2 % 128)
	atomic.CompareAndSwapInt32()
}
