package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var N, M int
	fmt.Fscan(in, &N, &M)

	g := make([][]int, N)
	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	dfn := make([]int, N)
	low := make([]int, N)
	for i := 0; i < N; i++ {
		dfn[i] = INF
		low[i] = INF
	}
	sum := make([]int, N)
	id := 0
	ans := 0
	var dfs func(int, int)
	dfs = func(cur, par int) {
		dfn[cur] = id
		low[cur] = id
		id++
		sum[cur] = 1
		tot := N - 1
		for _, neigh := range g[cur] {
			if neigh == par {
				continue
			}
			if dfn[neigh] == INF {
				dfs(neigh, cur)
				sum[cur] += sum[neigh]
				if dfn[cur] <= low[neigh] {
					tot -= sum[neigh]
					ans += sum[neigh] * tot
				}
				low[cur] = min(low[cur], low[neigh])
			} else {
				low[cur] = min(low[cur], dfn[neigh])
			}
		}
	}
	dfs(0, -1)
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
