package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	h := make([]int, n+2)
	dp := make([]int, n+2)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
		dp[i] = 1 << 60
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		dp[i+1] = min(dp[i+1], dp[i]+abs(h[i]-h[i+1]))
		dp[i+2] = min(dp[i+2], dp[i]+abs(h[i]-h[i+2]))
	}
	fmt.Println(dp[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
