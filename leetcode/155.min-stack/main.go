package main

import (
	"math"
)

func main() {

}

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64}, // Push 的时候要比较 top 和 x，所以得初始化一个值
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	// 获取 minStack 顶部的元素
	top := this.minStack[len(this.minStack)-1]
	// 比较 top 和 x，将值小的压栈
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

// Top 获取堆栈顶部的元素
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
