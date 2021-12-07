package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	dp := make([]int, n+1)
	dp[0] = 1

	const mod = 998244353
	for i := 0; i < n; i++ {
		for j := n; j >= 0; j-- {
			dp[j] *= max(0, j-i/m)
			dp[j] %= mod
			if j != 0 {
				dp[j] += dp[j-1]
				dp[j] %= mod
			}
		}
	}

	for j := 0; j < n; j++ {
		fmt.Println(dp[j+1])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
