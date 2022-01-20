package main

import "fmt"

func main() {
	var n, k int
	var s string
	fmt.Scan(&n, &k, &s)

	mod := 1_000_000_007
	dp := [2][303][303]int{}
	dp[0][0][0] = 1
	for i := 0; i < n; i++ {
		for x := 0; x <= k; x++ {
			for y := 0; y <= k; y++ {
				if s[i]^49 != 0 {
					dp[^i&1][x+1][max(y-1, 0)] += dp[i&1][x][y]
					dp[^i&1][x+1][max(y-1, 0)] %= mod
				}
				if s[i]^48 != 0 {
					dp[^i&1][max(x-1, 0)][y+1] += dp[i&1][x][y]
					dp[^i&1][max(x-1, 0)][y+1] %= mod
				}
				dp[i&1][x][y] = 0
			}
		}
	}

	ans := 0
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			ans = ans + dp[n&1][x][y]
			ans %= mod
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
