package main

import "fmt"

//func findKthLargest(nums []int, k int) int {
//
//}

func QuickSort(nums []int, begin, end int) {
	if begin >= end {
		return
	}

	pivot := partition(nums, begin, end)
	QuickSort(nums, begin, pivot)
	QuickSort(nums, pivot+1, end)
}

func partition(nums []int, begin, end int) int {
	pivot := end
	counter := begin

	for i := begin; i <= end; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[counter] = nums[counter], nums[i]
			counter++
		}
	}

	nums[pivot], nums[counter] = nums[counter], nums[pivot]
	return counter - 1 // fixme: pivot 如果分给前半部分，那前半部分排序结果永远不变，就会死递归
}

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
