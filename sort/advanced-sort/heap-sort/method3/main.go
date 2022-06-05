package main

import (
	"fmt"
	"math"
	"reflect"
)

// 最终使用次方法

// MaxHeapSort 大根堆,index 0 是最大值。维护堆
func MaxHeapSort(arr []int, n int) []int {
	if n < 2 { // fixme: 边界条件
		return arr
	}
	i := n - 1

	parent := int(math.Floor(float64((i - 1) / 2))) // fixme:  floor((i-1)/2)

	//for p := parent; p > 0; p-- {
	for ; parent >= 0; parent-- { // fixme: 要一直循环到索引为 0 的元素，即堆顶
		child1, child2 := 2*parent+1, 2*parent+2

		if child1 <= i && arr[child1] > arr[parent] {
			arr[parent], arr[child1] = arr[child1], arr[parent]
		}

		if child2 <= i && arr[child2] > arr[parent] { // fixme: 越界条件是 <=, = 是不越界的
			arr[parent], arr[child2] = arr[child2], arr[parent]
		}
	}

	return arr
}

func HeapSort(arr []int) {
	n := len(arr)

	for i := n - 1; i > 0; i-- {
		MaxHeapSort(arr, i+1) //fixme: 每次操作的都是同一个数组，只需要操作未进行排序的部分即可，所以需要传递给维护堆函数需要维护的数据长度。
		arr[0], arr[i] = arr[i], arr[0]
	}
}

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}

	HeapSort(arr)
	fmt.Println(arr)
	expect := []int{1, 2, 5, 8, 9, 10, 12, 30, 45, 63, 234}
	fmt.Println(reflect.DeepEqual(expect, arr))
}
