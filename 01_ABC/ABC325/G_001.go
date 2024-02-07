package main

import "fmt"

func main() {
	var s string
	var k int
	fmt.Scan(&s, &k)
	n := len(s)
	var dp [303][303]int
	for l := n - 1; l >= 0; l-- {
		for r := l + 1; r <= n; r++ {
			dp[l][r] = 1 + dp[l+1][r]
			if s[l] == 'o' {
				for i := l; i < r; i++ {
					if s[i] == 'f' && dp[l+1][i] == 0 {
						dp[l][r] = min(dp[l][r], max(0, dp[i+1][r]-k))
					}
				}
			}
		}
	}
	fmt.Println(dp[0][n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
