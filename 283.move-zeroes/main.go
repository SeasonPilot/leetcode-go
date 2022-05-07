package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

// 所谓要把0移动到数组后面，其实就是把非0数给移动到数组前面，而每个非0数需要移动的步数其实就是这个非0数前面0的个数。
func moveZeroes(nums []int) {
	count := 0 // 记录当前 0 的个数
	for i, num := range nums {
		if num == 0 {
			count++
		} else {
			nums[i-count] = nums[i] // 非 0 数需要移动的步数其实就是这个非 0 数前面 0 的个数。
		}
	}

	for i := len(nums) - count; i < len(nums); i++ {
		nums[i] = 0
	}

	fmt.Println(nums)
}
