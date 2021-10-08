package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = 1 << 60
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		dp[i+1] = min(dp[i+1], dp[i]+1)
		for j := 6; i+j < n+1; j *= 6 {
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
		for j := 9; i+j < n+1; j *= 9 {
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
	}
	fmt.Println(dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
