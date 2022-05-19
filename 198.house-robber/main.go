package main

import "fmt"

func main() {
	//testCase1 := []int{1, 2, 3, 1}
	testCase2 := []int{2, 7, 9, 3, 1}
	//var testCase3 []int    // nil
	//testCase4 := []int{}   // len=0
	//fmt.Println(rob(testCase1))
	fmt.Println(practice1(testCase2))
}

func rob(nums []int) int {
	n := len(nums)

	if nums == nil {
		fmt.Println("数组 nil")
		return 0
	}
	if n == 0 {
		fmt.Println("数组 0")
		return 0
	}

	dp := initTwoDimensionalSlice(n, 2)
	fmt.Printf("%v\n", dp) // [[0 0] [0 0] [0 0] [0 0]]

	// 这两行的作用是？  初始化二维数组
	dp[0][0] = 0
	dp[0][1] = nums[0]
	fmt.Printf("%v\n", dp) // [[0 1] [0 0] [0 0] [0 0]]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][1], dp[i-1][0])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	fmt.Printf("%v\n", dp)
	return max(dp[n-1][0], dp[n-1][1])
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 二维切片初始化
func initTwoDimensionalSlice(m, n int) [][]int {
	a := make([][]int, m) // 二维切片，3行
	for i := range a {
		a[i] = make([]int, n) // 每一行4列
	}
	return a
}

func practice1(nums []int) int {
	n := len(nums)

	if n == 0 || nums == nil {
		return 0
	}

	dp := initTwoDimensionalSlice(n, 2)

	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return max(dp[n-1][0], dp[n-1][1])
}
