package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 3e18 + 7

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)
	for tc := 0; tc < T; tc++ {
		var N int64
		fmt.Fscan(in, &N)

		dp := make([][][2]int64, 70)
		for i := range dp {
			dp[i] = make([][2]int64, 200)
			for j := range dp[i] {
				dp[i][j][0] = INF
				dp[i][j][1] = INF
			}
		}

		dp[0][100][0] = 0
		t := int64(1)

		for i := 0; i < 60; i, t = i+1, t*2 {
			for j := 0; j < 200; j++ {
				for k := 0; k < 2; k++ {
					if dp[i][j][k] == INF {
						continue
					}

					if k == 1 && ((N>>i)&1) != 0 {
						dp[i+1][j][1] = min(dp[i+1][j][1], dp[i][j][k])
						dp[i+1][j][1] = min(dp[i+1][j][1], dp[i][j][k]+t)
					} else if k == 1 || ((N>>i)&1) != 0 {
						dp[i+1][j+1][0] = min(dp[i+1][j+1][0], dp[i][j][k])
						dp[i+1][j-1][1] = min(dp[i+1][j-1][1], dp[i][j][k]+t)
					} else {
						dp[i+1][j][0] = min(dp[i+1][j][0], dp[i][j][k])
						dp[i+1][j][0] = min(dp[i+1][j][0], dp[i][j][k]+t)
					}
				}
			}
		}
		fmt.Fprintln(out, dp[60][100][0])
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
