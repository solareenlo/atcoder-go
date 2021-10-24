package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	dp := [5050][5050]int{}
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = 1
				if i > 0 {
					dp[i][j] = max(1, dp[i-1][j-1]+1)
				}
				res = max(res, min(j-i, dp[i][j]))
			}
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
