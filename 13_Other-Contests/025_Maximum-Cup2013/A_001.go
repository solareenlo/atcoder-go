package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e9)

	var n int
	fmt.Fscan(in, &n)
	if n == 2 {
		fmt.Println(0, 0)
		return
	}

	var g [15][15]int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			fmt.Fscan(in, &g[i][j])
			g[j][i] = g[i][j]
		}
	}

	var dp [1 << 15][15]int
	for i := 0; i < 1<<15; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
	}
	dp[1][0] = 0
	for i := 1; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if (((^i) >> k) & 1) != 0 {
					dp[i|1<<k][k] = min(dp[i|1<<k][k], dp[i][j]+g[j][k])
				}
			}
		}
	}

	res := INF
	for i := 0; i < n; i++ {
		res = min(res, dp[(1<<n)-1][i]+g[i][0])
	}
	fmt.Println(n, res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
