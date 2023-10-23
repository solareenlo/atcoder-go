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

	var P, F [10][50000]int
	var dp [11][50000]int
	var DP, L, R [50000]int

	var H, W int
	fmt.Fscan(in, &H, &W)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &P[i][j])
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &F[i][j])
		}
	}
	for j := 1; j < W; j++ {
		dp[0][j] = -int(1e18)
	}
	for i := 0; i < H; i++ {
		now := 0
		for j := 0; j < W; j++ {
			if j != 0 {
				now += P[i][j-1] - F[i][j-1] - F[i][j]
			}
			if now < 0 {
				now = 0
			}
			L[j] = now
		}
		now = 0
		for j := W - 1; j >= 0; j-- {
			if j+1 < W {
				now += P[i][j+1] - F[i][j+1] - F[i][j]
			}
			if now < 0 {
				now = 0
			}
			R[j] = now
		}
		for j := 0; j < W; j++ {
			DP[j] = dp[i][j] + L[j]
		}
		now = -int(1e18)
		for j := 0; j < W; j++ {
			now = max(now, DP[j]) + P[i][j] - F[i][j]
			dp[i+1][j] = now + R[j]
		}
		for j := 0; j < W; j++ {
			DP[j] = dp[i][j] + R[j]
		}
		now = -int(1e18)
		for j := W - 1; j >= 0; j-- {
			now = max(now, DP[j]) + P[i][j] - F[i][j]
			dp[i+1][j] = max(dp[i+1][j], now+L[j])
		}
	}
	for i := 0; i < W; i++ {
		fmt.Fprintln(out, dp[H][i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
