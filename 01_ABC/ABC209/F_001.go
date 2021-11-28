package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	h := make([]int, n+2)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&h[i])
	}

	mod := int(1e9 + 7)
	dp := [4004][4004]int{}
	dp[1][1] = 1
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if h[i] == h[i+1] {
				dp[i+1][j] = dp[i][i]
			} else if h[i] < h[i+1] {
				dp[i+1][j] = (dp[i][i] - dp[i][j-1] + mod) % mod
			} else {
				dp[i+1][j] = dp[i][j-1]
			}
			dp[i+1][j] = (dp[i+1][j-1] + dp[i+1][j]) % mod
		}
	}

	fmt.Println(dp[n][n])
}
