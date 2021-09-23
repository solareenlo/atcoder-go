package main

import "fmt"

func main() {
	var n, W int
	fmt.Scan(&n, &W)
	w, v := [101]int{}, [101]int{}
	for i := 1; i <= n; i++ {
		fmt.Scan(&w[i], &v[i])
	}

	dp := [101][310][101]int{}
	res := 0
	for i := 1; i <= n; i++ {
		for k := 1; k <= i; k++ {
			for j := 0; j <= i*3; j++ {
				if j < w[i]-w[1] {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-w[i]+w[1]][k-1]+v[i])
				}
				if j+k*w[1] <= W {
					res = max(res, dp[n][j][k])
				}
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
