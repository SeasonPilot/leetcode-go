package main

import "fmt"

func main() {
	fmt.Println(climbStairs(15))
}

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
