package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	fmt.Scan(&s)

	mi := n
	for i := 1; i < n; i++ {
		mi = min(mi, n-lcs(s[0:i], s[i:])*2)
	}
	fmt.Println(mi)
}

func lcs(s, t string) int {
	m := len(s)
	n := len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[m][n]
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
