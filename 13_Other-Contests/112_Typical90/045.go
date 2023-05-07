package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)

	X := make([]int, N)
	Y := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &X[i], &Y[i])
	}

	cost := make([]int, 1<<N)
	for k := 0; k < 1<<N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if ((k>>i)&1) != 0 && ((k>>j)&1) != 0 {
					cost[k] = max(cost[k], (X[i]-X[j])*(X[i]-X[j])+(Y[i]-Y[j])*(Y[i]-Y[j]))
				}
			}
		}
	}

	dp := make([][]int, K+1)
	for i := range dp {
		dp[i] = make([]int, 1<<N)
		for j := range dp[i] {
			dp[i][j] = 1 << 60
		}
	}
	dp[0][0] = 0
	for i := 1; i <= K; i++ {
		for S := 0; S < (1 << N); S++ {
			for T := S; T != 0; T = (T - 1) & S {
				dp[i][S] = min(dp[i][S], max(dp[i-1][S-T], cost[T]))
			}
		}
	}
	fmt.Println(dp[K][(1<<N)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
