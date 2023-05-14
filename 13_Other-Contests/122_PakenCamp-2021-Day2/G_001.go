package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var G, A [15]int
	var dp [15][1 << 15]int

	var N, M int
	fmt.Fscan(in, &N, &M)
	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		G[u] |= 1 << v
		G[v] |= 1 << u
	}
	for i := 0; i < (1 << N); i++ {
		dp[0][i] = -int(1e9)
	}
	dp[0][1] = 0
	for i := 0; i < N-1; i++ {
		for j := 0; j < N; j++ {
			fmt.Fscan(in, &A[j])
		}
		for j := 0; j < (1 << N); j++ {
			dp[i+1][j] = -int(1e9)
		}
		for j := 0; j < (1 << N); j++ {
			if dp[i][j] >= 0 {
				to := 0
				for k := 0; k < N; k++ {
					if ((j >> k) & 1) != 0 {
						to |= G[k]
					}
				}
				for k := 0; k < N; k++ {
					if ((j>>k)&1) == 0 && ((to>>k)&1) != 0 {
						dp[i+1][j|(1<<k)] = max(dp[i+1][j|(1<<k)], dp[i][j]+A[k])
					}
				}
			}
		}
	}
	fmt.Println(dp[N-1][(1<<N)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
