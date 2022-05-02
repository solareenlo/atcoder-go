package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var H, W int
	fmt.Fscan(in, &H, &W)

	dp := [101][101][202]int{}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			var A int
			fmt.Fscan(in, &A)
			for k := H + W; k > 0; k-- {
				dp[i][j][k] = max(dp[i][j][k], dp[i][j][k-1]+A)
			}
			for k := H + W; k > 0; k-- {
				if i+1 < H {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j][k])
				}
				if j+1 < W {
					dp[i][j+1][k] = max(dp[i][j+1][k], dp[i][j][k])
				}
			}
		}
	}

	for k := 1; k < H+W; k++ {
		fmt.Fprintln(out, dp[H-1][W-1][k])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
