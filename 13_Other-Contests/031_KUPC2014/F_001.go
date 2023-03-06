package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MAXN = 1000

	var N int
	fmt.Fscan(in, &N)
	x := make([]int, MAXN)
	y := make([]int, MAXN)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	d := make([]int, MAXN)
	c := make([]int, MAXN)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &d[i], &c[i])
	}
	G := make([][]int, MAXN)
	for i := 1; i < N; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	dp := make([][]int, MAXN)
	for i := range dp {
		dp[i] = make([]int, 3001)
	}
	var dfs func(int, int)
	dfs = func(u, p int) {
		ch := make([]int, 0)
		for _, v := range G[u] {
			if v != p {
				ch = append(ch, v)
				dfs(v, u)
			}
		}
		id := make([]int, len(ch))
		X := make([]int, len(ch))
		sum := 0
		for i := 0; i < len(ch); i++ {
			v := ch[i]
			for (d[u]+d[v]+id[i])*(d[u]+d[v]+id[i]) < (x[u]-x[v])*(x[u]-x[v])+(y[u]-y[v])*(y[u]-y[v]) {
				id[i]++
			}
			X[i] = 1e18
			for j := 3000; j >= id[i]; j-- {
				X[i] = min(X[i], dp[v][j])
			}
			sum += X[i]
		}
		for i := 0; i <= 3000; i++ {
			dp[u][i] = sum
			sum += c[u]
			for j := 0; j < len(ch); j++ {
				if id[j] > 0 {
					id[j]--
					v := ch[j]
					if X[j] > dp[v][id[j]] {
						sum -= X[j] - dp[v][id[j]]
						X[j] = dp[v][id[j]]
					}
				}
			}
		}
	}
	dfs(0, -1)
	ans := dp[0][0]
	for i := 1; i <= 3000; i++ {
		ans = min(ans, dp[0][i])
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
