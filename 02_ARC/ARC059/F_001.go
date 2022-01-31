package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	mod := int(1e9 + 7)
	dp := [5010][5010]int{}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			dp[i][j] = (dp[i-1][max(j-1, 0)] + dp[i-1][j+1]*2%mod) % mod
		}
	}
	fmt.Println(dp[n][len(s)])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
