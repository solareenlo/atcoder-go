package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	mod := 998244353
	dp := [3003][30003][3]int{}
	dp[0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < k+1; j++ {
			dp[i+1][j][0] += dp[i][j][0]
			dp[i+1][j][0] %= mod
			dp[i+1][j][1] += dp[i][j][0] + dp[i][j][1]
			dp[i+1][j][1] %= mod
			dp[i+1][j][2] += dp[i][j][0] + dp[i][j][1] + dp[i][j][2]
			dp[i+1][j][2] %= mod
			if j+a[i] <= k {
				dp[i+1][j+a[i]][1] += dp[i][j][0] + dp[i][j][1]
				dp[i+1][j+a[i]][1] %= mod
				dp[i+1][j+a[i]][2] += dp[i][j][0] + dp[i][j][1]
				dp[i+1][j+a[i]][2] %= mod
			}
		}
	}

	fmt.Println(dp[n][k][2])
}
