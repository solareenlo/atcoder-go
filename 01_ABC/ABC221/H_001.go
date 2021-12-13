package main

import "fmt"

func main() {
	const mod = 998244353

	var n, m int
	fmt.Scan(&n, &m)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	sum := make([]int, n+1)
	sum[0] = 1
	for i := 1; i < n+1; i++ {
		for j := i; j < n+1; j++ {
			dp[i][j] = sum[j-i] + dp[i][j-i]
			dp[i][j] %= mod
		}
		for j := 0; j < n+1; j++ {
			sum[j] += dp[i][j]
			sum[j] %= mod
			if m <= i {
				sum[j] -= dp[i-m][j]
				sum[j] %= mod
			}
		}
	}

	for i := 1; i < n+1; i++ {
		res := dp[i][n]
		if res < 0 {
			res += mod
		}
		fmt.Println(res)
	}
}
