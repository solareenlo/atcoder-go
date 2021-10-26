package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	h := [310]int{}
	for i := 1; i < n+1; i++ {
		fmt.Scan(&h[i])
	}

	dp := [310][310]int{}
	for i := 1; i < 310; i++ {
		dp[0][i] = 1 << 60
	}

	for i := 1; i < n+2; i++ {
		for j := 0; j < k+1; j++ {
			dp[i][j] = 1 << 60
			for r := i - 1; r >= 0; r-- {
				if i-r-1 <= j {
					dp[i][j] = min(dp[i][j], dp[r][j-i+r+1]+max(0, h[i]-h[r]))
				}
			}
		}
	}

	fmt.Println(dp[n+1][k])
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
