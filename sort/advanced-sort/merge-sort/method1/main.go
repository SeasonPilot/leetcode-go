package main

import "fmt"

// 分治  最终使用此方法和 method4、 method5 方法
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
func merge(l []int, r []int) []int {
	lLen := len(l)
	rLen := len(r)
	res := make([]int, 0)

	lIndex, rIndex := 0, 0 // 两个切片的下标，插入一个数据，下标加一
	for lIndex < lLen && rIndex < rLen {
		if l[lIndex] > r[rIndex] {
			res = append(res, r[rIndex])
			rIndex++
		} else {
			res = append(res, l[lIndex])
			lIndex++
		}
	}
	if lIndex < lLen { // 左边的还有剩余元素
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}

	return res
}

func main() {
	s := []int{6, 7, 8, 10, 4, 6, 99}
	res := mergeSort(s)
	fmt.Println(res)
}
