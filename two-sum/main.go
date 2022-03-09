package main

import "fmt"

func main() {
	nums := []int{2, 5, 5, 11}
	target := 10
	fmt.Println(twoSum(nums, target))
}

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
