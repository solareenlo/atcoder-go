package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	var score func(int, int) int
	score = func(l, r int) int {
		cnt := 0
		for i := 0; i < m; i++ {
			if l <= a[i] && b[i] <= r {
				cnt++
			}
		}
		return cnt
	}

	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -int(1e9)
		}
	}
	dp[0][0] = 0
	for i := 0; i < k; i++ {
		for j := 0; j < n+1; j++ {
			for l := 0; l < j; l++ {
				dp[i+1][j] = max(dp[i+1][j], dp[i][l]+score(l+1, j))
			}
		}
	}
	fmt.Println(dp[k][n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
