package main

import "fmt"

func main() {
	var n, K int
	fmt.Scan(&n, &K)

	dp := [55][55][55 * 55]int{}
	dp[0][0][0] = 1

	mod := int(1e9 + 7)
	for i := 1; i < n+1; i++ {
		for j := 0; j < i+1; j++ {
			for k := j + j; k < K+1; k++ {
				dp[i][j][k] += dp[i-1][j][k-2*j] * (2*j + 1) % mod
				dp[i][j][k] %= mod
				dp[i][j][k] += dp[i-1][j+1][k-2*j] * (j + 1) % mod * (j + 1) % mod
				dp[i][j][k] %= mod
				if j >= 1 {
					dp[i][j][k] += dp[i-1][j-1][k-2*j]
					dp[i][j][k] %= mod
				}
			}
		}
	}

	fmt.Println(dp[n][0][K])
}
