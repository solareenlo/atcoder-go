package main

import (
	"fmt"
	"math"
)

func main() {
	var n, P int
	fmt.Scan(&n, &P)

	dp := [5000][5000]int{}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for x, y := 1, 2; x < i && y < j; x, y = x*10, y+1 {
				dp[i][j] += dp[i-x][j-y] - dp[i-min(x*10, i)][j-y]
			}
			tmp := 0
			if int(math.Log10(float64(i)))+2 == j {
				tmp = 26
			}
			dp[i][j] = (dp[i][j]*25 + dp[i-1][j] + tmp) % P
		}
	}

	ans := 0
	for i := 1; i < n; i++ {
		ans = (ans + dp[n][i] - dp[n-1][i]) % P
	}
	fmt.Println((ans + P) % P)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
