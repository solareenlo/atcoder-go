package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	G := make([][]int, 1<<17)
	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	var K int
	A := make([]int, 17)
	fmt.Fscan(in, &A[0], &K)
	for i := 1; i <= K; i++ {
		fmt.Fscan(in, &A[i])
	}
	K++

	dist := make([][]int, 17)
	for i := 0; i < K; i++ {
		A[i]--
		dist[i] = make([]int, N)
		for j := range dist[i] {
			dist[i][j] = 1 << 60
		}
		dist[i][A[i]] = 0
		P := make([]int, 0)
		P = append(P, A[i])
		for len(P) > 0 {
			u := P[0]
			P = P[1:]
			for _, v := range G[u] {
				if dist[i][v] > dist[i][u]+1 {
					dist[i][v] = dist[i][u] + 1
					P = append(P, v)
				}
			}
		}
	}

	dp := [1 << 17][17]int{}
	for i := 0; i < 1<<K; i++ {
		for j := 0; j < K; j++ {
			dp[i][j] = 1_000_000_000
		}
	}
	dp[1][0] = 0
	for i := 1; i < 1<<K; i++ {
		for j := 0; j < K; j++ {
			if dp[i][j] < 1_000_000_000 {
				for k := 0; k < K; k++ {
					if (i >> k & 1) == 0 {
						dp[i|1<<k][k] = min(dp[i|1<<k][k], dp[i][j]+dist[j][A[k]])
					}
				}
			}
		}
	}

	ans := 1 << 60
	for i := 0; i < K; i++ {
		ans = min(ans, dp[(1<<K)-1][i])
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
