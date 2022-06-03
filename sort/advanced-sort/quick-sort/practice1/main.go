package main

import "fmt"

func QuickSort(arr []int, begin, end int) {
	// terminator
	if end < begin {
		return
	}
	// process logic in current level
	pivot := partition(arr, begin, end)
	// drill down
	QuickSort(arr, begin, pivot-1) // fixme: 这里应该是递归调用
	QuickSort(arr, pivot+1, end)
	// reverse
}

func partition(arr []int, begin, end int) int {
	pivot := end
	counter := begin

	for i := begin; i < end; i++ {
		if arr[i] < arr[pivot] {
			arr[counter], arr[i] = arr[i], arr[counter]
			counter++
		}
	}

	arr[counter], arr[pivot] = arr[pivot], arr[counter]
	return counter
}

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
