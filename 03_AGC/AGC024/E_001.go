package main

import "fmt"

var m int

func upd(x, y int) int {
	if x+y >= m {
		return x + y - m
	}
	return x + y
}

func main() {
	var n, k int
	fmt.Scan(&n, &k, &m)

	dp := [305][305][305]int{}
	dp[0][1][0] = 1
	for i := 0; i <= n; i++ {
		for j := 1; j <= k; j++ {
			for k := i; k >= 0; k-- {
				if k != 0 {
					dp[i][j][k-1] = upd(dp[i][j][k-1], dp[i][j][k])
				} else {
					dp[i][j+1][i] = upd(dp[i][j+1][i], dp[i][j][k])
				}
				dp[i+1][j][k] = upd(dp[i+1][j][k], dp[i][j][k]*(k+1)%m)
			}
		}
	}

	fmt.Println(dp[n][k][0])
}
