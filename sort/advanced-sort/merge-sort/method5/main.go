package main

import "fmt"

// 对齐小魔王 PPT 上的代码  最终使用此方法 和 method1、 method4 方法
func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) >> 1 // fixme: (end - start) >> 1

	mergeSort(arr, start, mid) // 闭区间
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

// 合并两个有序数组
func merge(arr []int, start, mid, end int) []int {
	var (
		i = start
		j = mid + 1
		k int // temp 中已经填入的元素个数
	)
	temp := make([]int, end-start+1) // 闭区间  辅助数组，存放合并排序的数

	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
			k++
		} else {
			temp[k] = arr[j]
			j++
			k++
		}
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= end {
		temp[k] = arr[j]
		j++
		k++
	}

	for idx, ele := range temp { // 将排序好的 temp copy 到 arr 中
		arr[start+idx] = ele
	}
	return arr
}

func main() {
	s := []int{6, 7, 8, 10, 4, 6, 99}
	mergeSort(s, 0, len(s)-1)
	fmt.Println(s)
}
