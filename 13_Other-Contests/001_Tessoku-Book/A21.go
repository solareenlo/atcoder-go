package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	P := make([]int, N+1)
	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &P[i], &A[i])
	}
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}
	dp[1][N] = 0
	for i := 1; i <= N; i++ {
		for j := N; j >= i; j-- {
			a, b := 0, 0
			if i > 1 {
				a = dp[i-1][j]
				if P[i-1] <= j && P[i-1] >= i {
					a += A[i-1]
				}
			}
			if j < N {
				b = dp[i][j+1]
				if P[j+1] <= j && P[j+1] >= i {
					b += A[j+1]
				}
			}
			dp[i][j] = max(a, b)
		}
	}
	ans := 0
	for i := 1; i <= N; i++ {
		ans = max(ans, dp[i][i])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
