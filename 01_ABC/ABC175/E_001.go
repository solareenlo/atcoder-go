package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var R, C, K int
	fmt.Fscan(in, &R, &C, &K)

	g := [3001][3001]int{}
	for i := 0; i < K; i++ {
		var r, c, v int
		fmt.Fscan(in, &r, &c, &v)
		g[r][c] = v
	}
	dp := [3001][3001][4]int{}
	for i := 1; i <= R; i++ {
		for j := 1; j <= C; j++ {
			for k := 1; k <= 3; k++ {
				dp[i][j][k] = max(dp[i-1][j][3]+g[i][j], max(dp[i][j-1][k-1]+g[i][j], dp[i][j-1][k]))
			}
		}
	}

	fmt.Println(dp[R][C][3])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
