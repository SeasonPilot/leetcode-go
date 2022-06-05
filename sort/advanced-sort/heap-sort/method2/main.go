package main

import "fmt"

// 大根堆
// 原文链接：https://blog.csdn.net/m0_37873625/article/details/116140191
//arr 存储堆的数组
//n 数组长度
//i 待维护元素的下标
func heapify(arr []int, n int, i int) {
	max := i
	//左孩子节点
	lChild := 2*i + 1
	//右孩子节点
	rChild := 2*i + 2
	if lChild < n && arr[lChild] > arr[max] {
		max = lChild
	}
	if rChild < n && arr[rChild] > arr[max] {
		max = rChild
	}
	//把父节点换下去并向下调整
	if max != i {
		arr[max], arr[i] = arr[i], arr[max]
		heapify(arr, n, max)
	}
}

func HeapSort(arr []int, n int) {
	//建堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 升序排序
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 将最大值放到最后
		heapify(arr, i, 0)              // 重新计算最大值
	}
}

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	HeapSort(arr, len(arr))
	fmt.Println(arr)
}
