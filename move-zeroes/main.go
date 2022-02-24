package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	count := 0
	for i, num := range nums {
		if num == 0 {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}

	for i := len(nums) - count; i < len(nums); i++ {
		nums[i] = 0
	}

	fmt.Println(nums)
}
