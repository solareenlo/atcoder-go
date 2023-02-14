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
	W := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		fmt.Fscan(in, &W[i])
		W[i] += W[i-1]
	}

	inf := 1 << 60
	dp := make([][]int, 4000)
	for i := range dp {
		dp[i] = make([]int, 4001)
	}
	idx := make([][]int, 4000)
	for i := range idx {
		idx[i] = make([]int, 4001)
	}
	for l := 0; l < N-1; l++ {
		dp[l][l+2] = W[l+2] - W[l]
		idx[l][l+2] = l + 1
	}

	for d := 3; d < N+1; d++ {
		for l := 0; l < N-d+1; l++ {
			r := l + d
			dp[l][r] = inf
			for i := idx[l][r-1]; i < idx[l+1][r]+1; i++ {
				if dp[l][r] > dp[l][i]+dp[i][r]+W[r]-W[l] {
					dp[l][r] = dp[l][i] + dp[i][r] + W[r] - W[l]
					idx[l][r] = i
				}
			}
		}
	}
	fmt.Println(dp[0][N])
}
