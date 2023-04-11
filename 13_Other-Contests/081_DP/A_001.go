package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	h := make([]int, n+3)
	for i := 1; i <= n; i++ {
		fmt.Scan(&h[i])
	}

	var dp [100001]int
	dp[1] = 0
	dp[2] = abs(h[1] - h[2])
	for i := 3; i <= n; i++ {
		dp[i] = min(abs(h[i]-h[i-1])+dp[i-1], abs(h[i]-h[i-2])+dp[i-2])
	}
	fmt.Println(dp[n])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
