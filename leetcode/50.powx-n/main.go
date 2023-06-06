package main

import "fmt"

//方法一： 快速幂 + 递归   https://leetcode.cn/problems/powx-n/solution/powx-n-by-leetcode-solution/
func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickMul(x, n)
	}
	return 1 / quickMul(x, n)
}

func quickMul(x float64, n int) float64 {
	//terminator
	if n == 0 {
		return 1
	}

	//current

	//drill down
	r := quickMul(x, n/2)

	// final result
	if n%2 == 0 {
		return r * r
	} else {
		return r * r * x
	}

	//reverse
}

func main() {
	fmt.Println(myPow(3, 3))
}
