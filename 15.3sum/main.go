package main

import "fmt"

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var result [][]int
	for i := 0; i < len(nums); i++ {
		target := nums[i]
		for j := 0; j < len(nums); j++ {
			for k := i + 1; k < len(nums); k++ {
				if nums[j]+nums[k] == -target {
					result = append(result, []int{nums[j], nums[k], target})
				}
			}
		}
	}
	return result
}
