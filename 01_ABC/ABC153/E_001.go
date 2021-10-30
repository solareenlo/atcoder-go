package main

import "fmt"

func main() {
	var h, n int
	fmt.Scan(&h, &n)

	dp := make([]int, h+1)
	for i := range dp {
		dp[i] = 1 << 60
	}
	dp[0] = 0

	for j := 0; j < n; j++ {
		var a, b int
		fmt.Scan(&a, &b)
		for i := 1; i < h+1; i++ {
			if i <= a {
				dp[i] = min(dp[i], b)
			} else {
				dp[i] = min(dp[i], max(dp[i-a], 0)+b)
			}
		}
	}

	fmt.Println(dp[h])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
