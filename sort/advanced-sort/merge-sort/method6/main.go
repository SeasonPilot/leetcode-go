package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	length := len(arr)
	// terminator
	if length == 1 {
		return arr //最后切割只剩下一个元素
	}

	// process logic in current level
	mid := length >> 1

	// drill down
	leftArr := mergeSort(arr[:mid])
	rightArr := mergeSort(arr[mid:])

	// merge result
	return merge(leftArr, rightArr)

	// reverse
}

// 把两个有序切片合并成一个有序切片
func merge(n1 []int, n2 []int) []int {
	// n1 和 n2 归并填入 nums
	p1, p2 := 0, 0
	nums := make([]int, len(n1)+len(n2))
	for i := range nums {
		if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
			nums[i] = n1[p1]
			p1++
		} else {
			nums[i] = n2[p2]
			p2++
		}
	}
	return nums
}

func main() {
	s := []int{6, 7, 8, 10, 4, 6, 99}
	res := mergeSort(s)
	fmt.Println(res)
}
