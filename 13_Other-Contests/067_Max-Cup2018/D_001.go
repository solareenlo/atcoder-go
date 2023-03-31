package main

import "fmt"

func main() {
	var n, m, l, x int
	fmt.Scan(&n, &m, &l, &x)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = 1e9
		}
	}

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i+1][j] = min(dp[i+1][j], dp[i][j])
			dp[i+1][(j+a[i])%m] = min(dp[i+1][(j+a[i])%m], dp[i][j]+(j+a[i])/m)
		}
	}
	if dp[n-1][l] <= x {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
