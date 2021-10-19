package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	s := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&s[i])
	}

	t := make([]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Scan(&t[i])
	}

	dp := make([][]int, 2020)
	for i := range dp {
		dp[i] = make([]int, 2020)
	}

	mod := int(1e9 + 7)
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if s[i] == t[j] {
				dp[i][j] = dp[i-1][j] + dp[i][j-1] + 1
			} else {
				dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % mod
				dp[i][j] += -dp[i-1][j-1] + mod
			}
			dp[i][j] %= mod
		}
	}
	fmt.Println((dp[n][m] + 1) % mod)
}
