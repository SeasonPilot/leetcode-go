package main

import (
	"reflect"
	"sort"
)

// 暴力求解
// 排序 + DeepEqual 进行去重
func threeSum(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	sort.Ints(nums) // 排序

	for i := 0; i < n; i++ {
		target := nums[i]
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[j]+nums[k] == -target {
					tmp := []int{nums[i], nums[j], nums[k]}
					//sort.Ints(tmp) // 排序
					result = append(result, tmp)
				}
			}
		}
	}
	return killRepetion(result)
}

// 二维 slice 去重
func killRepetion(nums [][]int) [][]int {
	newRes := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := i + 1; j < len(nums); j++ {
			if reflect.DeepEqual(nums[i], nums[j]) {
				flag = true
				break
			}
		}
		if !flag {
			newRes = append(newRes, nums[i])
		}
	}
	return newRes
}
