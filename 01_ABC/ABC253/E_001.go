package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	const N = 8000
	var dp [N][N]int
	for i := 1; i <= m; i++ {
		dp[1][i] = i
	}

	const mod = 998244353
	for i := 2; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if k != 0 {
				dp[i][j] = (dp[i][j-1] + dp[i-1][m] - dp[i-1][min(j+k-1, m)] + dp[i-1][max(j-k, 0)] + mod) % mod
			} else {
				dp[i][j] = dp[i-1][j] * m % mod
			}
		}
	}
	fmt.Println((dp[n][m] + mod) % mod)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
