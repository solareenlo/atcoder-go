package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sum := 0
	G := [17][17]int{}
	g := make([]int, 1<<17)
	for i := 1; i <= m; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		a--
		b--
		sum += c
		G[a][b] = c
		G[b][a] = c
		for j := 0; j < (1 << n); j++ {
			if (j>>a&1) != 0 && (j>>b&1) != 0 {
				g[j] += c
			}
		}
	}

	dp := [1 << 17][17]int{}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[1][0] = 0
	for i := 1; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] >= 0 {
				for k := 0; k < n; k++ {
					if (i>>k&1) == 0 && G[j][k] != 0 {
						dp[i|(1<<k)][k] = max(dp[i|(1<<k)][k], dp[i][j]+G[j][k])
					}
				}
				all := ((1 << n) - 1) ^ i
				for k := all; k > 0; k = (k - 1) & all {
					dp[i|k][j] = max(dp[i|k][j], dp[i][j]+g[k|(1<<j)])
				}
			}
		}
	}
	fmt.Println(sum - dp[(1<<n)-1][n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
