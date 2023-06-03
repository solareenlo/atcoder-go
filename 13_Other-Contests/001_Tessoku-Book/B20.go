package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var S, T string
	fmt.Fscan(in, &S, &T)
	s := len(S)
	t := len(T)
	dp := make([][]int, s+1)
	for i := range dp {
		dp[i] = make([]int, t+1)
	}
	for i := 0; i < s+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < t+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= s; i++ {
		for j := 1; j <= t; j++ {
			if S[i-1] == T[j-1] {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j]+1, dp[i][j-1]+1))
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}
	fmt.Println(dp[s][t])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
