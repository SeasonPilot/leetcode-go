package main

import "fmt"

// 最终使用此方法

// QuickSort 快排：先调配出左右子数组，然后对于左右子数组进行排序
func QuickSort(arr []int, begin, end int) { // begin, end 是 index
	// terminator
	if end < begin {
		return
	}
	// process logic in current level
	pivot := partition(arr, begin, end)
	// drill down
	QuickSort(arr, begin, pivot-1)
	QuickSort(arr, pivot+1, end)
}

func partition(arr []int, begin, end int) int {
	// pivot: 标杆位置，counter: ⼩于pivot的元素的个数
	pivot := end // 初始 pivot 在最后
	counter := begin

	for i := begin; i < end; i++ {
		if arr[i] < arr[pivot] { // 小于 arr[pivot] 的都放在 arr[pivot] 右边
			arr[counter], arr[i] = arr[i], arr[counter] // 交换
			counter++
		}
	}

	// 移动 pivot 到正确位置 (即所有左边的数都小于pivot，右边的数字都大于pivot)
	arr[counter], arr[pivot] = arr[pivot], arr[counter]
	return counter
}

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
