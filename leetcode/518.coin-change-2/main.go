package main

func main() {
	a := []int{7, 13, 29}
	change(20, a)
}

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func fa(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i < amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
