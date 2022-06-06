package main

// leetcode/0493.Reverse-Pairs/493. Reverse Pairs.go
// 解法一 归并排序 mergesort，时间复杂度 O(n log n)
func reversePairs(nums []int) int {
	buf := make([]int, len(nums))
	return mergesortCount(nums, buf)
}

func mergesortCount(nums, buf []int) int {
	if len(nums) <= 1 {
		return 0
	}
	mid := (len(nums) - 1) / 2
	cnt := mergesortCount(nums[:mid+1], buf)
	cnt += mergesortCount(nums[mid+1:], buf)

	// count elements
	for i, j := 0, mid+1; i < mid+1; i++ { // Note!!! j is increasing.
		for ; j < len(nums) && nums[i] <= 2*nums[j]; j++ {
		}
		cnt += len(nums) - j
	}

	// merge sort
	copy(buf, nums)
	for i, j, k := 0, mid+1, 0; k < len(nums); {
		if j >= len(nums) || i < mid+1 && buf[i] > buf[j] {
			nums[k] = buf[i]
			i++
		} else {
			nums[k] = buf[j]
			j++
		}
		k++
	}
	return cnt
}
