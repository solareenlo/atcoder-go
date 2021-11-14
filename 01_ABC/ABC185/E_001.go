package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	dp := [1005][1005]int{}
	a := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&a[i])
		dp[i][0] = i
	}

	b := make([]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Scan(&b[i])
		dp[0][i] = i
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if a[i] == b[j] {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])+1)
			} else {
				dp[i][j] = min(dp[i-1][j-1]+1, min(dp[i-1][j], dp[i][j-1])+1)
			}
		}
	}

	fmt.Println(dp[n][m])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
