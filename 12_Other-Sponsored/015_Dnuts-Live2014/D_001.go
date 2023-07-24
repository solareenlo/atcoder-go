package main

import "fmt"

func main() {
	const MOD = 1000000007

	var dp [1001][1001]int

	for i := 1; i <= 1000; i++ {
		dp[i][1] = 1
		dp[i][i] = 1
		for j := 2; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + j*dp[i-1][j]
			dp[i][j] %= MOD
		}
	}

	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(dp[n][m])
}
