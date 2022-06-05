package main

import (
	"fmt"
	"math"
	"reflect"
)

// HeapSortMax 大根堆,index 0 是最大值。维护堆
func HeapSortMax(arr []int, n int) []int {
	if n <= 1 {
		return arr
	}

	//depth := n/2 - 1 // 二叉树深度  // 第 n 个元素 上一层的层级，root depth 为 0
	depth := int(math.Floor(float64((n - 1) / 2))) // 索引为 n 的父结点的索引

	for i := depth; i >= 0; i-- {

		// index
		max := i // 假定最大的位置就在 i 的位置
		lChild := 2*i + 1
		rChild := 2*i + 2

		//if lChild <= n-1 && arr[lChild] > arr[max] { // 防止越过界限
		if lChild < n && arr[lChild] > arr[max] { // 防止越过界限
			max = lChild
		}
		if rChild < n && arr[rChild] > arr[max] { // 防止越过界限
			max = rChild
		}
		//把父节点换下去并向下调整
		if max != i {
			arr[i], arr[max] = arr[max], arr[i]
		}
	}

	return arr
}

func HeapSort(arr []int) []int {
	n := len(arr)

	//for i := 0; i < n; i++ {
	//	lastlen := n - i
	//	HeapSortMax(arr, lastlen) // 重新计算最大值
	//	if i < n {
	//		arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0] // 将最大值放到最后
	//	}
	//}

	// 升序排序
	for i := n - 1; i > 0; i-- {
		HeapSortMax(arr, i+1) // 重新计算最大值
		if i < n {
			arr[0], arr[i] = arr[i], arr[0] // 将最大值放到最后
		}
	}

	return arr
}

//堆排序
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	//arr := []int{100, 92}
	ans := HeapSort(arr)
	fmt.Println(ans)
	expect := []int{1, 2, 5, 8, 9, 10, 12, 30, 45, 63, 234}
	fmt.Println(reflect.DeepEqual(expect, ans))
}
