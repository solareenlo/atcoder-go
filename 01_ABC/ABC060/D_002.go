package main

import "fmt"

func main() {
	var n, W int
	fmt.Scan(&n, &W)
	w, v := [101]int{}, [101]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&w[i], &v[i])
	}
	cp := w[0]
	for i := range w {
		w[i] -= cp
	}

	res := 0
	dp := [110][350][110]int{}
	for i := 0; i < n; i++ {
		for k := 0; k <= n; k++ {
			for j := 0; j <= 310; j++ {
				dp[i+1][j+w[i]][k+1] = max(dp[i+1][j+w[i]][k+1], dp[i][j][k]+v[i])
				dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j][k])
				if j+k*cp <= W {
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
