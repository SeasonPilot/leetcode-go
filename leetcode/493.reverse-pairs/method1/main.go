package main

import "fmt"

// 最终使用此方法
func reversePairs(nums []int) int {
	if len(nums) == 0 { // 没有元素，没有翻转对
		return 0
	}
	return mergeSort(nums, 0, len(nums)-1) // 归并的范围：0到末尾
}

func mergeSort(arr []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := (start + end) >> 1

	cnt := mergeSort(arr, start, mid) + mergeSort(arr, mid+1, end) // 闭区间

	////count elements  第一种写法  双指针，通过 i 统计 cnt
	////此时左右序列已升序，现在做：合并前的统计、以及合并
	//i := start                 // 左序列的开头
	//j := mid + 1               // 右序列的开头
	//for i <= mid && j <= end { // i j 都不越界
	//	if arr[i] > 2*arr[j] {
	//		cnt += mid - i + 1 // i 到 mid，都ok
	//		j++                // 考察下一个j，继续找 i
	//	} else { // 当前i不满足，考察下一个i
	//		i++
	//	}
	//}

	// count elements  第二种写法  双指针，通过 j 统计 cnt
	// 此时左右序列已升序，现在做：合并前的统计
	j := mid + 1
	for i := start; i <= mid; i++ { // i j 都不越界
		for j <= end && arr[i] > 2*arr[j] { // 满足条件就移动 j 指针
			j++ // 考察下一个j，继续找 i
		}
		cnt += j - (mid + 1) // 因为此时左右序列已升序，所以当前 j 往前一直到 (mid + 1) 的元素都满足 arr[i] > 2*arr[j] 条件
	}

	merge(arr, start, mid, end)

	return cnt
}

// 合并两个有序数组
func merge(arr []int, start, mid, end int) {
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
}

func main() {
	s := []int{9, 8, 7, 4, 7, 2, 3, 8, 7, 0}

	fmt.Println(reversePairs(s))
}
