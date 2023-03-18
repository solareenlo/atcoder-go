package main

import (
	"fmt"
)

func main() {
	var n, r, c int
	fmt.Scan(&n, &r, &c)

	dp := make([][]int, r)
	for i := range dp {
		dp[i] = make([]int, c)
		for j := range dp[i] {
			dp[i][j] = int(1e9)
		}
	}
	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j < r; j++ {
			var s string
			fmt.Scan(&s)
			for k := 0; k < c; k++ {
				if s[k] != 'H' {
					dp[j][k] += int(s[k] - '0')
					if j+1 < r {
						dp[j+1][k] = min(dp[j+1][k], dp[j][k])
					}
					if k+1 < c {
						dp[j][k+1] = min(dp[j][k+1], dp[j][k])
					}
					if i+1 < n {
						dp[j][k] = int(1e9)
					}
				}
			}
		}
	}

	fmt.Println(dp[r-1][c-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
