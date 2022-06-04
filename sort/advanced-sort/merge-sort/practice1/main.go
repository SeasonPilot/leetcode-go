package main

import "fmt"

func mergeSort(arr []int) []int {
	length := len(arr)
	// terminator
	if length == 1 {
		return arr
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

func merge(leftArr, rightArr []int) []int {
	lLen := len(leftArr)
	rLen := len(rightArr)
	var ret []int

	var i, j int
	for i < lLen && j < rLen { // fixme:最大index 应该是 len()-1
		if leftArr[i] <= rightArr[j] {
			ret = append(ret, leftArr[i])
			i++
		} else {
			ret = append(ret, rightArr[j]) // fixme: typo  rightArr[i]
			j++
		}
	}

	if i < lLen {
		//ret = append(ret, leftArr...)
		ret = append(ret, leftArr[i:]...) //fixme: 应该是 剩余 数组,不是整个数组
	}
	if j < rLen {
		ret = append(ret, rightArr[j:]...)
	}
	return ret
}

func main() {
	arr := []int{6, 7, 8, 10, 4, 6, 99}

	fmt.Println(mergeSort(arr))
}
