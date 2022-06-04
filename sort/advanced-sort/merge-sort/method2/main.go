package main

import "fmt"

// MergeSort 未调通,错误的
func MergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i := left
	j := mid + 1
	k := 0
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			k++
			i++
		} else {
			temp[k] = arr[j]
			k++
			i++
		}

	}
	for i < mid {
		temp[k] = arr[i]
		k++
		i++
	}
	for j <= right {
		temp[k] = arr[j]
		k++
		j++
	}

	for p := 0; p < len(temp); p++ {
		arr[left+p] = temp[p]
	}
}

func main() {
	//s := []int{6, 7, 8, 10, 4, 6, 99}
	s := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}

	MergeSort(s, 0, len(s)-1)
	fmt.Println(s)
}
