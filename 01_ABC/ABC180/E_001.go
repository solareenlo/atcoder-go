package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n)
	y := make([]int, n)
	z := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i], &z[i])
	}

	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 1 << 60
		}
	}
	dp[0][0] = 0

	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				dist := abs(x[k]-x[j]) + abs(y[k]-y[j]) + max(0, z[j]-z[k])
				dp[i|(1<<k)][k] = min(dp[i|(1<<k)][k], dp[i][j]+dist)
			}
		}
	}

	fmt.Println(dp[1<<n-1][0])
}

func max(a, b int) int {
	if a > b {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
