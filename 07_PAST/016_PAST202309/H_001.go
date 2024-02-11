package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int, m+2)
			for k := range dp[i][j] {
				dp[i][j][k] = -1 << 60
			}
		}
	}
	dp[0][0][0] = 0
	for i := 0; i < n; i++ {
		var v int
		fmt.Fscan(in, &v)
		for k := 0; k <= m; k++ {
			dp[i+1][1][k+1] = max(dp[i+1][1][k+1], dp[i][0][k])
			dp[i+1][0][k] = max(dp[i+1][0][k], dp[i][0][k]+v)
			dp[i+1][0][k] = max(dp[i+1][0][k], dp[i][1][k]+v)
		}
	}
	fmt.Println(max(dp[n][0][m], dp[n][1][m]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
