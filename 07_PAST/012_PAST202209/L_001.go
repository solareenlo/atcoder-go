package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var dp, G [200200][3][3]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	A := make([]int, n)
	P := make([]int, m)
	Q := make([]int, m)
	L := make([]int, m)
	R := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &P[i], &Q[i], &L[i], &R[i])
		Q[i]--
		G[Q[i]][L[i]][R[i]] += P[i]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			if i == 0 && j != 0 {
				continue
			}
			for k := 0; k < 3; k++ {
				dp[i][j][k] += G[i][j][k]
				dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j][k])
				dp[i+1][(j+1)%3][(k+2)%3] = max(dp[i+1][(j+1)%3][(k+2)%3], dp[i][j][k]+A[i])
			}
		}
	}
	ans := 0
	for j := 0; j < 3; j++ {
		dp[n][j][0] += G[n][j][0]
		ans = max(ans, dp[n][j][0])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
