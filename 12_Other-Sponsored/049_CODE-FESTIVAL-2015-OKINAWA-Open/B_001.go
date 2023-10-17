package main

import (
	"fmt"
)

func main() {
	var N, X, Y int
	fmt.Scan(&N, &X, &Y)
	a := make([]int, N)
	b := make([]int, N)

	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, 10001)
	}

	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	dp[0][0] = 0

	for i := 0; i < N; i++ {
		fmt.Scan(&a[i], &b[i])
		for j := i; j >= 0; j-- {
			for k := 0; k <= 10000; k++ {
				if dp[j][k] == -1 {
					continue
				}
				newK := min(10000, k+b[i])
				dp[j+1][newK] = max(dp[j+1][newK], dp[j][k]+a[i]+b[i])
			}
		}
	}

	for i := 0; i <= N; i++ {
		for j := Y; j <= 10000; j++ {
			if dp[i][j] >= X+Y {
				fmt.Println(i)
				return
			}
		}
	}
	fmt.Println(-1)
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
