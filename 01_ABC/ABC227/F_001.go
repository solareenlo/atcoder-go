package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	a := [35][35]int{}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	res := 1 << 60
	dp := [35]int{}
	for x := 1; x < n+1; x++ {
		for y := 1; y < m+1; y++ {
			for i := range dp {
				dp[i] = 1 << 60
			}
			dp[1] = a[x][y] * k
			for i := 1; i < n+1; i++ {
				for j := 1; j < m+1; j++ {
					dp[j] = min(dp[j], dp[j-1]) + max(a[i][j]-a[x][y], 0)
				}
			}
			res = min(res, dp[m])
		}
	}

	fmt.Println(res)
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
