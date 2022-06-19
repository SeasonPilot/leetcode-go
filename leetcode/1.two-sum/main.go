package main

import "fmt"

func main() {
	nums := []int{2, 5, 5, 11}
	target := 10
	//fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
}

// 方法一：暴力枚举  O(n²)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 方法二：哈希表  O(n)
func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i // 存储
	}
	return nil
}
