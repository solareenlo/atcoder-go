package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	const mod = 1_000_000_007
	dp := [5005][5005]int{}
	dp[0][m] = 1
	for i := 1; i <= n; i++ {
		sum := 0
		for j := m; j >= 1; j-- {
			dp[i][j] = (sum*j + dp[i-1][j]*(j+1)) % mod
			sum = (sum*2 + dp[i-1][j]) % mod
		}
	}

	sum := 0
	for i := 1; i <= m; i++ {
		sum = (sum + dp[n][i]) % mod
	}
	fmt.Println(sum)
}
